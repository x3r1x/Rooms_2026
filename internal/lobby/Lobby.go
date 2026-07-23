package lobby

import (
	"encoding/json"
	"gamedevRooms/internal/game"
	"gamedevRooms/internal/model"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Lobby struct {
	state           string
	players         map[string]*LobbyPlayer
	playersReady    int
	gameLoop        *game.GameLoop
	addChan         chan *LobbyPlayer
	readyChan       chan *LobbyPlayer
	removeChan      chan string
	getStateChan    chan chan string
	getGameLoopChan chan chan *game.GameLoop
}

func NewLobby() *Lobby {
	l := &Lobby{
		state:           model.WaitingLobbyState,
		players:         make(map[string]*LobbyPlayer),
		addChan:         make(chan *LobbyPlayer),
		readyChan:       make(chan *LobbyPlayer),
		removeChan:      make(chan string),
		getStateChan:    make(chan chan string),
		getGameLoopChan: make(chan chan *game.GameLoop),
	}
	go l.Run()
	return l
}

func (l *Lobby) AddPlayerToLobby(nickname string, conn *websocket.Conn) string {
	player := NewLobbyPlayer(nickname)
	player.Connection = conn
	l.addChan <- player
	return player.Id
}

func (l *Lobby) UpdatePlayerInLobby(msg *LobbyPlayer) {
	l.readyChan <- msg
}

func (l *Lobby) RemovePlayerFromLobby(id string) {
	l.removeChan <- id
}

func (l *Lobby) GetState() string {
	responseChan := make(chan string)
	l.getStateChan <- responseChan
	return <-responseChan
}

func (l *Lobby) GetGameLoop() *game.GameLoop {
	responseChan := make(chan *game.GameLoop)
	l.getGameLoopChan <- responseChan
	return <-responseChan
}

func (l *Lobby) Run() {
	for {
		select {
		case regPlayer := <-l.addChan:
			l.players[regPlayer.Id] = regPlayer
			log.Printf("Игрок %s присоединился. Всего: %d", regPlayer.Nickname, len(l.players))
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
				len(l.players) >= model.MinCountOfPlayers &&
				l.state == model.WaitingLobbyState {
				log.Println("Все готовы! Запускаем игру...")
				l.StartGame()
			}
		case delPlayer := <-l.removeChan:
			l.removeUser(delPlayer)
		case responseChan := <-l.getStateChan:
			responseChan <- l.state
		case responseChan := <-l.getGameLoopChan:
			responseChan <- l.gameLoop
		}
	}
}

func (l *Lobby) broadcastLobbyState() {
	players := make([]model.LobbyPlayerMessage, 0, len(l.players))
	for _, p := range l.players {
		players = append(players, model.LobbyPlayerMessage{
			Nickname: p.Nickname,
			Id:       p.Id,
			Ready:    p.Ready,
		})
	}

	for _, p := range l.players {
		msg := model.ServerLobbyMessage{
			State:   l.state,
			OwnId:   p.Id,
			Players: players,
		}

		data, _ := json.Marshal(msg)
		if p.Connection != nil {
			if err := p.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println("Невозможно отправить состояние лобби игроку")
			}
		}
	}
}

func (l *Lobby) removeUser(id string) {
	player, exists := l.players[id]
	if !exists {
		return
	}

	if player.Ready {
		l.playersReady--
	}

	delete(l.players, id)
	log.Printf("Игрок %s покинул лобби. Осталось: %d", id, len(l.players))
}

func (l *Lobby) StartGame() {
	if l.state == model.OngoingGameState || len(l.players) == 0 {
		return
	}

	gameState := game.NewGameState()

	for _, player := range l.players {
		if player.Ready {
			gameState.AddPlayer(&model.PlayerGameState{
				Id:          player.Id,
				Nickname:    player.Nickname,
				Health:      model.MaxPlayerHealth,
				X:           model.PlayerSpawnPointX,
				Y:           model.PlayerSpawnPointY,
				Angle:       model.InitDirection,
				MoveX:       model.InitValue,
				MoveY:       model.InitValue,
				Connection:  player.Connection,
				ShootTimer:  0,
				RebornTimer: 0,
				BodyCount:   0,
				DeathCount:  0,
			})
		}
	}

	mapManager := game.NewMapManager(gameState)
	roomMessages := mapManager.GetRoomMessages()
	l.sendReadyState(roomMessages)
	l.doCountdown()

	l.gameLoop = game.NewGameLoop(gameState)
	l.state = model.OngoingGameState

	go l.gameLoop.Run()
	//for id := range l.players {
	//	delete(l.players, id)
	//}
	l.playersReady = 0
}

func (l *Lobby) sendReadyState(roomMessages map[string]model.RoomMessage) {
	msg := model.ServerReadyMessage{
		State:     model.ReadyLobbyState,
		Countdown: 5.0,
		Map:       roomMessages,
	}
	data, _ := json.Marshal(msg)
	for _, p := range l.players {
		if p.Connection != nil {
			if err := p.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
				return
			}
		}
	}
}

func (l *Lobby) doCountdown() {
	for i := 5; i > 0; i-- {
		msg := model.ServerCountdownMessage{
			State:     model.CountdownLobbyState,
			Countdown: i,
		}
		data, _ := json.Marshal(msg)
		for _, p := range l.players {
			if err := p.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
				return
			}
		}
		time.Sleep(1 * time.Second)
	}
}
