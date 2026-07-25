package game

import (
	"encoding/json"
	"gamedevRooms/internal/adapters/collision"
	"gamedevRooms/internal/domain"
	"log"

	"github.com/gorilla/websocket"
)

type GameLoop struct {
	game             *GameState
	collisionService *collision.CollisionService
	updateChan       chan domain.ClientGameMessage
	deleteChan       chan string
	finishChan       chan bool
	stopChan         chan bool
}

func NewGameLoop(game *GameState, finishChan chan bool) *GameLoop {
	return &GameLoop{game: game,
		collisionService: collision.NewCollisionService(game),
		updateChan:       make(chan domain.ClientGameMessage),
		deleteChan:       make(chan string),
		finishChan:       finishChan,
		stopChan:         make(chan bool)}
}

func (gl *GameLoop) broadcast(message domain.ServerGameMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}
	for id, p := range gl.game.GetAllPlayers() {
		if p.Connection == nil {
			continue
		}
		err := p.Connection.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Ошибка отправки: ", err)
			gl.deleteChan <- id
		}
	}
}
