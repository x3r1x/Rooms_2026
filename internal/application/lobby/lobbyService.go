package lobby

import (
	"gamedevRooms/internal/application/game"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/ports"
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
	removeChan chan string

	gameStateProvider ports.GameStateProvider
	mapManager        ports.MapManager
	broadcastService  ports.BroadcastService
}

type lobbyAddEvent struct {
	player *domain.LobbyPlayer
	conn   *websocket.Conn
}

func NewLobbyService(
	gameStateProvider ports.GameStateProvider,
	mapManager ports.MapManager,
	broadcastService ports.BroadcastService,
) *LobbyService {
	l := &LobbyService{
		state:             domain.WaitingLobbyState,
		players:           make(map[string]*domain.LobbyPlayer),
		gameStateProvider: gameStateProvider,
		mapManager:        mapManager,
		broadcastService:  broadcastService,
		addChan:           make(chan lobbyAddEvent),
		readyChan:         make(chan *domain.LobbyPlayer),
		removeChan:        make(chan string),
	}
	go l.run()
	return l
}

func (l *LobbyService) AddPlayerToLobby(nickname string, conn *websocket.Conn) string {
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

func (l *LobbyService) RemovePlayerFromLobby(id string) {
	l.removeChan <- id
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

func (l *LobbyService) removeUser(id string) {
	player, exists := l.players[id]
	if !exists {
		l.broadcastService.RemoveConnection(id)
		if l.state == domain.OngoingGameState && l.gameService != nil {
			l.gameService.DeletePlayer(id)
		}
		return
	}

	if player.Ready {
		l.playersReady--
	}

	l.broadcastService.RemoveConnection(id)
	if l.state == domain.OngoingGameState && l.gameService != nil {
		l.gameService.DeletePlayer(id)
	}

	delete(l.players, id)
	log.Printf("Игрок %s покинул лобби. Осталось: %d", id, len(l.players))
	l.broadcastLobbyState()
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

	l.mapManager.LoadMapObjects(l.gameStateProvider)

	for _, player := range l.players {
		if player.Ready {
			l.gameStateProvider.AddPlayer(&domain.PlayerGameState{
				Id:          player.Id,
				Nickname:    player.Nickname,
				Health:      domain.MaxPlayerHealth,
				X:           domain.PlayerSpawnPointX,
				Y:           domain.PlayerSpawnPointY,
				Angle:       domain.InitDirection,
				MoveX:       domain.InitValue,
				MoveY:       domain.InitValue,
				ShootTimer:  domain.InitValue,
				RebornTimer: domain.InitValue,
				BodyCount:   domain.InitValue,
				DeathCount:  domain.InitValue,
			})
		}
	}

	roomMessages := l.mapManager.GetRoomMessages()
	l.sendReadyState(roomMessages)
	l.doCountdown()

	l.state = domain.OngoingGameState
	if len(roomMessages) > 0 {
		var firstRoomId string
		for roomId := range roomMessages {
			firstRoomId = roomId
			break
		}
		for _, player := range l.gameStateProvider.GetAllPlayers() {
			l.gameStateProvider.SetPlayerRoom(player.Id, firstRoomId)
		}
	}

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

	for _, player := range l.gameStateProvider.GetAllPlayers() {
		l.gameStateProvider.RemovePlayer(player.Id)
	}

	l.broadcastLobbyState()
}
