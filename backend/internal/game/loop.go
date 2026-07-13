package game

import (
	"gamedevRooms/internal/models"
	"math"
	"time"
)

func StartGameLoop() {
	ticker := time.NewTicker(16 * time.Millisecond)
	serverMessage := models.ServerMessage{
		Players: make([]models.PlayerState, 0, len(Clients)),
	}
	for range ticker.C {
		Mutex.Lock()
		//TODO: вынести в функцию
		for _, player := range Clients {
			writeIdx := 0
			for readIdx, _ := range player.Bullets {
				bullet := &player.Bullets[readIdx]
				bullet.Life--
				if bullet.Life <= 0 {
					continue
				}
				bullet.X += math.Cos(bullet.Direction) * MAX_BULLET_SPEED
				bullet.Y += math.Sin(bullet.Direction) * MAX_BULLET_SPEED
				// обдумать условие
				if bullet.Y >= 0 && bullet.Y <= MAP_SIZE && bullet.X >= 0 && bullet.X <= MAP_SIZE {
					player.Bullets[writeIdx] = *bullet
					writeIdx++
				}
			}
			player.Bullets = player.Bullets[:writeIdx]
		}

		serverMessage.Players = serverMessage.Players[:0]
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
