package lobby

import (
	"encoding/json"
	"gamedevRooms/internal/adapters/map"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/game"
	"log"

	"github.com/gorilla/websocket"
)

func (l *Lobby) StartGame() {
	if l.state == domain.OngoingGameState || len(l.players) == 0 {
		return
	}

	gameState := game.NewGameState()

	for _, player := range l.players {
		if player.Ready {
			gameState.AddPlayer(&model.PlayerGameState{
				Id:          player.Id,
				Nickname:    player.Nickname,
				Health:      domain.MaxPlayerHealth,
				X:           domain.PlayerSpawnPointX,
				Y:           domain.PlayerSpawnPointY,
				Angle:       domain.InitDirection,
				MoveX:       domain.InitValue,
				MoveY:       domain.InitValue,
				Connection:  player.Connection,
				ShootTimer:  domain.InitValue,
				RebornTimer: domain.InitValue,
				BodyCount:   domain.InitValue,
				DeathCount:  domain.InitValue,
			})
		}
	}

	mapManager := _map.NewMapManager(gameState)
	roomMessages := mapManager.GetRoomMessages()
	l.sendReadyState(roomMessages)
	l.doCountdown()

	l.gameLoop = game.NewGameLoop(gameState, l.gameFinishChan)
	l.state = domain.OngoingGameState

	go l.gameLoop.Run()
	l.playersReady = 0
}

func (l *Lobby) sendReadyState(roomMessages map[string]domain.RoomMessage) {
	msg := domain.ServerReadyMessage{
		State:     domain.ReadyLobbyState,
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

func (l *Lobby) returnToLobby() {
	log.Println("Игра завершена! Возврат в лобби...")

	l.state = domain.WaitingLobbyState
	l.playersReady = 0
	l.gameLoop = nil

	for _, player := range l.players {
		player.Ready = false
	}

	l.broadcastLobbyState()

	log.Printf("Лобби открыто! Игроков: %d", len(l.players))
}
