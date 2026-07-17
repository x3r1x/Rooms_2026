package game

import (
	"encoding/json"
	"fmt"
	"gamedevRooms/internal/models"
	"log"
	"math"
	"time"

	"github.com/gorilla/websocket"
)

type GameLoop struct {
	game *models.GameState
}

func NewGameLoop(game *models.GameState) *GameLoop {
	return &GameLoop{game: game}
}

func (gl *GameLoop) Run() {
	ticker := time.NewTicker(models.TickTime * time.Millisecond)
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
			gl.broadcast(models.ServerMessage{
				Type:    "a",
				Players: snapshot,
				Bullets: gl.game.GetAllBullets(),
			})
		}
	}
}

func (gl *GameLoop) handleRegister(reg models.PlayerState) {
	fmt.Println("Register ", reg)
	gl.game.AddPlayer(reg)
}

func (gl *GameLoop) handleDelete(del string) {
	fmt.Println("Delete ", del)
	gl.game.RemovePlayer(del)
}

func (gl *GameLoop) handleUpdate(upd models.ClientMessage) {
	player, exist := gl.game.GetPlayer(upd.Id)
	if !exist {
		return
	}
	player.A = upd.A
	player.MoveX = upd.MX
	player.MoveY = upd.MY

	if upd.S && player.ShootTimer <= 0 {
		gl.spawnBullet(player)
		player.ShootTimer = models.ShootCooldown
	}
}

func (gl *GameLoop) spawnBullet(player models.PlayerState) {
	localX := models.PlayerVisualSize / 2.0
	localY := (models.PlayerVisualSize / 2.0) - (models.BulletWidth + (models.PlayerVisualSize * 0.1))

	rotatedDX := localX*math.Cos(player.A) - localY*math.Sin(player.A)
	rotatedDY := localX*math.Sin(player.A) + localY*math.Cos(player.A)
	bullet := models.Bullet{
		X:         player.X + rotatedDX,
		Y:         player.Y + rotatedDY,
		Direction: player.A,
		Life:      models.BulletLife,
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
	activeBullets := make([]models.Bullet, 0)
	for _, bullet := range gl.game.GetAllBullets() {
		bullet.Life--
		bullet.X += math.Cos(bullet.Direction) * models.MaxBulletSpeed
		bullet.Y += math.Sin(bullet.Direction) * models.MaxBulletSpeed
		if bullet.Life > 0 {
			activeBullets = append(activeBullets, bullet)
		}
	}
	gl.game.SetBullets(activeBullets)
}

func (gl *GameLoop) updatePlayers() {
	for _, player := range gl.game.GetAllPlayers() {
		gl.NormaliseDirection(&player.MoveX, &player.MoveY)

		player.X += player.MoveX * float64(models.TickTime) * models.PlayerSpeed
		player.Y += player.MoveY * float64(models.TickTime) * models.PlayerSpeed
	}
}

func (gl *GameLoop) NormaliseDirection(moveX, moveY *float64) {
	var vectorLength = math.Sqrt(*moveX**moveX + *moveY**moveY)

	if vectorLength != 0 {
		*moveX /= vectorLength
		*moveY /= vectorLength
	}
}

func (gl *GameLoop) createSnapshot() []models.PlayerState {
	snapshot := make([]models.PlayerState, 0, len(gl.game.GetAllPlayers()))
	for _, player := range gl.game.GetAllPlayers() {
		snapshot = append(snapshot, player)
	}
	return snapshot
}

func (gl *GameLoop) broadcast(message models.ServerMessage) {
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
