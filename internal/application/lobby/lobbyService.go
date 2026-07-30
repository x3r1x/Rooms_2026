package lobby

import (
	"gamedevRooms/internal/adapters/broadcast"
	_map "gamedevRooms/internal/adapters/map"
	"gamedevRooms/internal/application/game"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/recovery"
	"gamedevRooms/internal/state"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type LobbyService struct {
	mu           sync.RWMutex
	state        string
	players      map[string]*domain.LobbyPlayer
	playersReady int
	gameService  *game.GameService

	addChan    chan lobbyAddEvent
	readyChan  chan *domain.LobbyPlayer
	removeChan chan lobbyRemoveEvent

	gameState        *state.GameState
	mapManager       *_map.MapManager
	broadcastService *broadcast.BroadcastService
}

type lobbyAddEvent struct {
	player *domain.LobbyPlayer
	conn   *websocket.Conn
}

type lobbyRemoveEvent struct {
	id    string
	force bool
}

func NewLobbyService(
	gameStateProvider *state.GameState,
	mapManager *_map.MapManager,
	broadcastService *broadcast.BroadcastService,
) *LobbyService {
	l := &LobbyService{
		state:            domain.WaitingLobbyState,
		players:          make(map[string]*domain.LobbyPlayer),
		gameState:        gameStateProvider,
		mapManager:       mapManager,
		broadcastService: broadcastService,
		addChan:          make(chan lobbyAddEvent),
		readyChan:        make(chan *domain.LobbyPlayer),
		removeChan:       make(chan lobbyRemoveEvent),
	}
	go l.run()
	return l
}

func (l *LobbyService) AddPlayerToLobby(nickname string, conn *websocket.Conn) string {
	l.mu.RLock()
	currentCount := len(l.players)
	l.mu.RUnlock()
	if currentCount >= 10 {
		conn.Close()
		return ""
	}
	player := domain.NewLobbyPlayer(nickname)
	l.addChan <- lobbyAddEvent{
		player: player,
		conn:   conn,
	}
	return player.Id
}

func (l *LobbyService) UpdatePlayerInLobby(msg *domain.LobbyPlayer) {
	l.readyChan <- msg
}

func (l *LobbyService) RemovePlayerFromLobby(id string, force bool) {
	l.removeChan <- lobbyRemoveEvent{id, force}
}

func (l *LobbyService) GetState() string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.state
}

func (l *LobbyService) GetGameService() *game.GameService {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.gameService
}

func (l *LobbyService) run() {
	defer recovery.Recover()
	for {
		select {
		case reg := <-l.addChan:
			l.broadcastService.AddConnection(reg.player.Id, reg.conn)
			l.players[reg.player.Id] = reg.player
			log.Printf("Игрок %s присоединился. Всего: %d", reg.player.Nickname, len(l.players))
			l.broadcastLobbyState()
		case updPlayer := <-l.readyChan:
			player, exists := l.players[updPlayer.Id]
			if !exists {
				continue
			}
			wasReady := player.Ready

			if player.Ready == updPlayer.Ready {
				continue
			}
			player.Ready = updPlayer.Ready

			if updPlayer.Ready && !wasReady {
				l.playersReady++
				l.players[player.Id].PlayerClass = updPlayer.PlayerClass
			} else if !updPlayer.Ready && wasReady {
				l.playersReady--
			}
			log.Printf("Игрок %s готовность: %v (готово: %d/%d)",
				updPlayer.Id, updPlayer.Ready, l.playersReady, len(l.players))

			l.broadcastLobbyState()
			if l.playersReady == len(l.players) &&
				len(l.players) >= domain.MinCountOfPlayers &&
				l.state == domain.WaitingLobbyState {
				log.Println("Все готовы! Запускаем игру...")
				l.StartGame()
			}
		case delPlayer := <-l.removeChan:
			l.removeUser(delPlayer)
		}
	}
}

func (l *LobbyService) broadcastLobbyState() {
	players := make([]domain.LobbyPlayerMessage, 0, len(l.players))
	for _, p := range l.players {
		players = append(players, domain.LobbyPlayerMessage{
			Nickname: p.Nickname,
			Id:       p.Id,
			Ready:    p.Ready,
		})
	}

	for _, p := range l.players {
		msg := domain.ServerLobbyMessage{
			State:   l.state,
			OwnId:   p.Id,
			Players: players,
		}
		l.broadcastService.BroadcastToPlayer(p.Id, msg)
	}
}

func (l *LobbyService) sendReadyState(roomMessages map[string]domain.RoomMessage) {
	msg := domain.ServerReadyMessage{
		State:     domain.ReadyLobbyState,
		Countdown: 5.0,
		Map:       roomMessages,
	}

	for _, p := range l.players {
		l.broadcastService.BroadcastToPlayer(p.Id, msg)
	}
}

func (l *LobbyService) removeUser(event lobbyRemoveEvent) {
	player, exists := l.players[event.id]
	if !exists {
		l.broadcastService.RemoveConnection(event.id)
		if l.state == domain.OngoingGameState && l.gameService != nil {
			l.gameService.DeletePlayer(event.id)
		}
		return
	}

	if player.Ready && !event.force {
		return
	}

	if player.Ready {
		l.playersReady--
	}

	l.broadcastService.RemoveConnection(event.id)
	if l.state == domain.OngoingGameState && l.gameService != nil {
		l.gameService.DeletePlayer(event.id)
	}

	delete(l.players, event.id)
	log.Printf("Игрок %s покинул лобби. Осталось: %d", event.id, len(l.players))
	l.broadcastLobbyState()
	if l.state == domain.WaitingLobbyState {
		currentTotal := len(l.players)
		if currentTotal >= domain.MinCountOfPlayers &&
			l.playersReady == currentTotal {

			log.Println("После удаления игрока все оставшиеся готовы! Запускаем игру...")
			l.StartGame()
		}

	}
}

func (l *LobbyService) doCountdown() {
	for i := 5; i > 0; i-- {
		msg := domain.ServerCountdownMessage{
			State:     domain.CountdownLobbyState,
			Countdown: i,
		}
		l.broadcastService.BroadcastToAll(msg)
		time.Sleep(1 * time.Second)
	}
}

func (l *LobbyService) StartGame() {
	if l.state == domain.OngoingGameState || len(l.players) == 0 {
		return
	}

	if l.gameService == nil {
		log.Println("ERROR: GameService is nil! Cannot start game.")
		return
	}
	l.gameState.SetRoomManager(l.mapManager)

	for _, player := range l.players {
		if player.Ready {
			l.gameState.AddPlayer(domain.NewPlayerGameState(
				player.Id,
				player.Nickname,
				player.PlayerClass,
			))
		}
	}

	objects := l.mapManager.LoadMapObjects(l.gameState.GetCountOfPlayers())
	l.gameState.SetObjects(objects)
	roomMessages := l.mapManager.GetRoomMessages()
	roomIDs := make([]string, 0, len(roomMessages))
	for roomId := range roomMessages {
		roomIDs = append(roomIDs, roomId)
	}
	l.state = domain.OngoingGameState

	if len(roomMessages) > 0 {
		i, k := 1, 0
		var spawnX, spawnY float64
		for _, player := range l.gameState.GetAllPlayers() {
			position := i % 3
			switch position {
			case 0:
				spawnX = domain.RoomPixelWidth - domain.PlayerHalfSize - 50
				spawnY = domain.RoomPixelHeight / 2
			case 1:
				spawnX = domain.RoomPixelWidth / 2
				spawnY = domain.RoomPixelHeight - domain.PlayerHalfSize - 50
			case 2:
				spawnX = domain.PlayerHalfSize + 50
				spawnY = domain.RoomPixelHeight / 2
			}
			player.X = spawnX
			player.Y = spawnY
			l.gameState.SetPlayerRoom(player.Id, roomIDs[k])
			if i < 3 {
				i++
			} else {
				i = 1
				k++
			}
		}
	}
	l.sendReadyState(roomMessages)
	l.doCountdown()

	go l.gameService.Run()
	l.playersReady = 0
	for _, player := range l.players {
		delete(l.players, player.Id)
	}
}

func (l *LobbyService) SetGameService(gs *game.GameService) {
	l.gameService = gs
}

func (l *LobbyService) HandleGameEnd() {
	log.Println("Игра закончилась, возвращаемся в лобби")
	l.state = domain.WaitingLobbyState
	l.playersReady = 0

	for _, player := range l.gameState.GetAllPlayers() {
		l.gameState.RemovePlayer(player.Id)
	}

	l.broadcastLobbyState()
}
