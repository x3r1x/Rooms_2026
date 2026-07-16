package game

import (
	"gamedevRooms/internal/models"
	"math"
	"time"
)

func GoGameLoop() {
	ticker := time.NewTicker(16 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {

		case reg := <-models.Game.RegisterChan:
			models.Game.Players[reg.Id] = &models.PlayerState{
				Id:         reg.Id,
				X:          reg.X,
				Y:          reg.Y,
				Direction:  reg.Direction,
				Bullets:    make([]models.Bullet, 0),
				Connection: reg.Connection,
			}
		case del := <-models.Game.LeaveChan:
			delete(models.Game.Players, del)

		case upload := <-models.Game.InputChan:
			if player, exist := models.Game.Players[upload.Player.Id]; exist {
				player.X = upload.Player.X
				player.Y = upload.Player.Y
				player.Direction = upload.Player.Direction
			}
		case <-ticker.C:
			models.Game.TickCount++
			updateBullets()
			snapshot := createSnapshot()
			broadcast(models.ServerMessage{Players: snapshot})
		}
	}
}

func createSnapshot() []models.PlayerState {
	snapshot := make([]models.PlayerState, 0, len(models.Game.Players))
	for _, player := range models.Game.Players {
		snapshot = append(snapshot, *player)
	}
	return snapshot
}

func updateBullets() {
	for _, player := range models.Game.Players {
		activeBullets := make([]models.Bullet, 0)
		for _, bullet := range player.Bullets {
			bullet.Life--
			bullet.X += math.Cos(bullet.Direction) * models.MAX_BULLET_SPEED
			bullet.Y += math.Sin(bullet.Direction) * models.MAX_BULLET_SPEED
			if bullet.Life > 0 {
				activeBullets = append(activeBullets, bullet)
			}
		}
		player.Bullets = activeBullets
	}
}
