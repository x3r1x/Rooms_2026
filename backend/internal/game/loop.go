package game

import (
	"gamedevRooms/internal/models"
	"math"
	"sync"
	"time"
)

var ServerMessagePool = sync.Pool{
	New: func() interface{} {
		return &models.ServerMessage{
			Players: make([]models.PlayerState, 0, 100),
		}
	},
}

func UpdateBullets(player *models.PlayerState) {
	writeIdx := 0
	for readIdx := range player.Bullets {
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

func StartGameLoop() {
	ticker := time.NewTicker(16 * time.Millisecond)
	//TODO: мапа на клиентов, чтобы избавиться от дублирования.
	//TODO: мапа на пули, для регулирования при попадании в стену и удалении из последовательности середины

	for range ticker.C {
		ProcessCommands()

		msg := ServerMessagePool.Get().(*models.ServerMessage)
		msg.Players = msg.Players[:0]

		for _, player := range Clients {
			UpdateBullets(player)
			msg.Players = append(msg.Players, models.PlayerState{
				X:         player.X,
				Y:         player.Y,
				Direction: player.Direction,
				Id:        player.Id,
				Bullets:   player.Bullets,
			})
		}
		broadcast(*msg)
		ServerMessagePool.Put(msg)
	}
}
