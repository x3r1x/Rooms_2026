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
				Id:            reg.Id,
				X:             reg.X,
				Y:             reg.Y,
				DirectionLook: reg.DirectionLook,
				Connection:    reg.Connection,
			}
		case del := <-models.Game.LeaveChan:
			delete(models.Game.Players, del)

		case upload := <-models.Game.InputChan:
			if upload.Player.S
			//if player, exist := models.Game.Players[upload.Player.Id]; exist {
			//	player.X = upload.Player.X
			//	player.Y = upload.Player.Y
			//	player.DirectionLook = upload.Player.Direction
			//	player.Bullets = upload.Player.Bullets
			//}
		case <-ticker.C:
			models.Game.TickCount++
			updateBullets()
			snapshot := createSnapshot()
			broadcast(models.ServerMessage{
				Type: "a",
				Players: snapshot,
				Bullets: getAllBullets(),
			})
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
		activeBullets := make([]models.Bullet, 0)
		for _, bullet := range models.Game.Bullets {
			bullet.Life--
			bullet.X += math.Cos(bullet.Direction) * models.MaxBulletSpeed
			bullet.Y += math.Sin(bullet.Direction) * models.MaxBulletSpeed
			if bullet.Life > 0 {
				activeBullets = append(activeBullets, bullet)
			}
		}
		models.Game.Bullets = activeBullets
}

func getAllBullets() []models.Bullet {
	return models.Game.Bullets
}