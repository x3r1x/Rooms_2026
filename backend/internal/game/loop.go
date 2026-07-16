package game

import (
	"fmt"
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
			fmt.Println("Register ", reg)
			models.Game.Players[reg.Id] = &models.PlayerState{
				Id:         reg.Id,
				X:          reg.X,
				Y:          reg.Y,
				A:          reg.A,
				Connection: reg.Connection,
				MoveX:      reg.MoveX,
				MoveY:      reg.MoveY,
			}
		case del := <-models.Game.LeaveChan:
			delete(models.Game.Players, del)

		case upload := <-models.Game.InputChan:
			if player, exist := models.Game.Players[upload.Id]; exist {
				player.A = upload.A
				if upload.S && player.ShootTimer <= 0 {
					spawnBullet(player)
					player.ShootTimer = models.ShootCooldown
				}
			}
		case <-ticker.C:
			models.Game.TickCount++
			updateShooterTimers()
			updateBullets()
			snapshot := createSnapshot()
			broadcast(models.ServerMessage{
				Type:    "a",
				Players: snapshot,
				Bullets: getAllBullets(),
			})
		}
	}
}

func spawnBullet(player *models.PlayerState) {
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
	models.Game.Bullets = append(models.Game.Bullets, bullet)
}

func updateShooterTimers() {
	for _, player := range models.Game.Players {
		if player.ShootTimer > 0 {
			player.ShootTimer--
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
