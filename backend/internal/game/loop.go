package game

import (
	"gamedevRooms/internal/models"
	"math"
	"time"
)

func StartGameLoop() {
	ticker := time.NewTicker(16 * time.Millisecond)
	for range ticker.C {
		Mutex.Lock()
		//TODO: вынести в функцию
		for _, player := range Clients {
			activeBullets := make([]models.Bullet, 0, len(player.Bullets))
			for _, bullet := range player.Bullets {
				bullet.Life--
				bullet.X += math.Cos(bullet.Direction) * MAX_BULLET_SPEED
				bullet.Y += math.Sin(bullet.Direction) * MAX_BULLET_SPEED
				// обдумать условие
				if bullet.Life > 0 && bullet.Y >= 0 && bullet.Y <= MAP_SIZE && bullet.X >= 0 && bullet.X <= MAP_SIZE {
					activeBullets = append(activeBullets, bullet)
				}
			}
			player.Bullets = activeBullets
		}

		serverMessage := models.ServerMessage{
			Players: make([]models.PlayerState, 0, len(Clients)),
		}
		for _, player := range Clients {
			serverMessage.Players = append(serverMessage.Players, models.PlayerState{
				X:         player.X,
				Y:         player.Y,
				Direction: player.Direction,
				Id:        player.Id,
				Bullets:   player.Bullets,
			})
		}
		Mutex.Unlock()
		broadcast(serverMessage)
	}
}
