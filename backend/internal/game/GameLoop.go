package game

import (
	"encoding/json"
	"fmt"
	"gamedevRooms/internal/model"
	"log"
	"math"
	"time"

	"github.com/gorilla/websocket"
)

type GameLoop struct {
	game *GameState
}

func NewGameLoop(game *GameState) *GameLoop {
	return &GameLoop{game: game}
}

func (gl *GameLoop) Run() {
	ticker := time.NewTicker(model.TickTime * time.Millisecond)
	defer ticker.Stop()

	for {
		select {

		case reg := <-gl.game.RegisterChan:
			gl.handleRegister(reg)
		case del := <-gl.game.DeleteChan:
			gl.handleDelete(del)
		case upload := <-gl.game.UpdateChan:
			gl.handleUpdate(upload)
		case <-ticker.C:
			gl.updateGameState()
			snapshot := gl.createSnapshot()
			gl.broadcast(model.ServerMessage{
				Type:    "a",
				Players: snapshot,
				Bullets: gl.game.GetAllBullets(),
			})
		}
	}
}

func (gl *GameLoop) handleRegister(reg *model.PlayerState) {
	fmt.Println("Register ", reg)
	gl.game.AddPlayer(reg)
}

func (gl *GameLoop) handleDelete(del string) {
	fmt.Println("Delete ", del)
	gl.game.RemovePlayer(del)
}

func (gl *GameLoop) handleUpdate(upd model.ClientMessage) {
	player, exist := gl.game.GetPlayer(upd.Id)
	if !exist {
		return
	}
	player.Angle = upd.Angle
	player.MoveX = upd.MX
	player.MoveY = upd.MY

	if upd.IsShoot && player.ShootTimer <= 0 {
		gl.spawnBullet(player)
		player.ShootTimer = model.ShootCooldown
	}
}

func (gl *GameLoop) spawnBullet(player *model.PlayerState) {
	localX := model.PlayerVisualSize / 2.0
	localY := (model.PlayerVisualSize / 2.0) - (model.BulletWidth + (model.PlayerVisualSize * 0.1))

	rotatedDX := localX*math.Cos(player.Angle) - localY*math.Sin(player.Angle)
	rotatedDY := localX*math.Sin(player.Angle) + localY*math.Cos(player.Angle)
	bullet := model.Bullet{
		X:         player.X + rotatedDX,
		Y:         player.Y + rotatedDY,
		Direction: player.Angle,
		Life:      model.BulletLife,
		OwnerId:   player.Id,
	}
	gl.game.AddBullet(bullet)
}

func (gl *GameLoop) updateGameState() {
	gl.game.IncrementTick()
	gl.updateShooterTimers()
	gl.updateBullets()
	gl.updatePlayers()
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

		if bullet.Life > 0 && !gl.checkCollision(bullet) {
			activeBullets = append(activeBullets, bullet)
		}
	}
	gl.game.SetBullets(activeBullets)
}

func (gl *GameLoop) checkCollision(bullet model.Bullet) bool {
	for _, player := range gl.game.GetAllPlayers() {
		if player.Id == bullet.OwnerId {
			continue
		}
		if bullet.X == player.X && bullet.Y == player.Y {
			player.Health -= model.BulletDamage
			fmt.Println("HIT! The player: ", player.Id, ", got shoot. Now he have this health", player.Health)
			return true
		}
	}
	return false
}

func (gl *GameLoop) updatePlayers() {
	for _, player := range gl.game.GetAllPlayers() {
		gl.normaliseDirection(&player.MoveX, &player.MoveY)

		player.X += player.MoveX * float64(model.TickTime) * model.PlayerSpeed
		player.Y += player.MoveY * float64(model.TickTime) * model.PlayerSpeed
	}
}

func (gl *GameLoop) normaliseDirection(moveX, moveY *float64) {
	var vectorLength = math.Sqrt(*moveX**moveX + *moveY**moveY)

	if vectorLength != 0 {
		*moveX /= vectorLength
		*moveY /= vectorLength
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
			gl.game.DeleteChan <- id
		}
	}
}
