package game

import (
	"encoding/json"
	"gamedevRooms/internal/model"
	"log"
	"math"
	"time"

	"github.com/gorilla/websocket"
)

type GameLoop struct {
	game             *GameState
	collisionService *CollisionService
	updateChan       chan model.ClientMessage
	registerChan     chan *model.PlayerState
	deleteChan       chan string
}

func NewGameLoop(game *GameState) *GameLoop {
	return &GameLoop{game: game,
		collisionService: NewCollisionService(game),
		updateChan:       make(chan model.ClientMessage),
		registerChan:     make(chan *model.PlayerState),
		deleteChan:       make(chan string),
	}
}

func (gl *GameLoop) RegisterPlayer(player *model.PlayerState) {
	gl.registerChan <- player
}

func (gl *GameLoop) UpdatePlayer(msg model.ClientMessage) {
	gl.updateChan <- msg
}

func (gl *GameLoop) DeletePlayer(id string) {
	gl.deleteChan <- id
}
func (gl *GameLoop) Run() {
	ticker := time.NewTicker(model.TickTime * time.Millisecond)
	defer ticker.Stop()

	for {
		select {

		case reg := <-gl.registerChan:
			gl.game.AddPlayer(reg)
		case del := <-gl.deleteChan:
			gl.game.RemovePlayer(del)
		case upload := <-gl.updateChan:
			gl.game.UpdatePlayer(upload)
		case <-ticker.C:
			gl.game.IncrementTick()
			gl.updateShooterTimers()
			gl.updateBullets()
			gl.updatePlayers()
			gl.broadcast(model.ServerMessage{
				Type:    "a",
				Players: gl.createSnapshot(),
				Bullets: gl.game.GetAllBullets(),
			})
		}
	}
}

func (gl *GameLoop) updateShooterTimers() {
	for _, player := range gl.game.GetAllPlayers() {
		if player.ShootTimer > 0 {
			player.ShootTimer--
		}
	}
}

func (gl *GameLoop) updateBullets() {
	activeBullets := make([]model.Bullet, 0)
	for _, bullet := range gl.game.GetAllBullets() {
		bullet.Life--
		bullet.X += math.Cos(bullet.Direction) * model.MaxBulletSpeed
		bullet.Y += math.Sin(bullet.Direction) * model.MaxBulletSpeed

		if bullet.Life > 0 {
			hit, player := gl.collisionService.CheckBulletCollision(bullet)
			if hit {
				gl.collisionService.HandleHit(player, bullet)
			} else {
				activeBullets = append(activeBullets, bullet)
			}
		}
	}
	gl.game.SetBullets(activeBullets)
}

func (gl *GameLoop) updatePlayers() {
	for _, player := range gl.game.GetAllPlayers() {
		vectorLength := math.Sqrt(player.MoveX*player.MoveX + player.MoveY*player.MoveY)

		if vectorLength != 0 {
			player.MoveX /= vectorLength
			player.MoveY /= vectorLength
		}

		player.X += player.MoveX * model.TickTime * model.PlayerSpeed
		player.Y += player.MoveY * model.TickTime * model.PlayerSpeed
	}
}

func (gl *GameLoop) createSnapshot() []model.PlayerState {
	snapshot := make([]model.PlayerState, 0, len(gl.game.GetAllPlayers()))
	for _, player := range gl.game.GetAllPlayers() {
		snapshot = append(snapshot, *player)
	}
	return snapshot
}

func (gl *GameLoop) broadcast(message model.ServerMessage) {
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
