package game

import (
	"encoding/json"
	"gamedevRooms/internal/adapters/collision"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/model"
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

func (gl *GameLoop) shouldStop() bool {
	return !gl.game.IsGameActive() || len(gl.game.GetAllPlayers()) == 0
}

func (gl *GameLoop) cleanup() {
	log.Println("GameLoop: очистка ресурсов...")
	gl.game.SetBullets([]model.Bullet{})
	gl.game.SetObjects(make(map[string]*model.Object))
	gl.game.SetGameActive(false)
	if gl.finishChan != nil {
		select {
		case gl.finishChan <- true:
			log.Println("GameLoop: отправлен сигнал о завершении в лобби")
		default:
			log.Println("GameLoop: канал завершения уже содержит сигнал")
		}
	}
}

func (gl *GameLoop) getStatistics() []model.PlayerFinalState {
	stats := make([]model.PlayerFinalState, 0, len(gl.game.GetAllPlayers()))
	for _, player := range gl.game.GetAllPlayers() {
		stats = append(stats, model.PlayerFinalState{
			Nickname: player.Nickname,
			Id:       player.Id,
			Deaths:   player.DeathCount,
			Kills:    player.BodyCount,
		})
	}
	return stats
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

func (gl *GameLoop) broadcastFinal(message domain.ServerEndMessage) {
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
