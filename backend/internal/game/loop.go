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
			Type:    "absolute",
			Players: make([]models.PlayerSnapshot, 0, 100),
		}
	},
}

func UpdateBullets(player *models.PlayerState) {
	for id, bullet := range player.Bullets {
		bullet.Life--
		if bullet.Life <= 0 {
			delete(player.Bullets, id)
			continue
		}
		bullet.X += math.Cos(bullet.Direction) * MAX_BULLET_SPEED
		bullet.Y += math.Sin(bullet.Direction) * MAX_BULLET_SPEED
		// обдумать условие
		if bullet.Y >= 0 && bullet.Y <= MAP_SIZE && bullet.X >= 0 && bullet.X <= MAP_SIZE {
			delete(player.Bullets, id)
			continue
		}
		player.Bullets[id] = bullet
	}
}

func UpdatePlayerPosition(player *models.PlayerState) {
	player.X += math.Cos(player.Direction) * PLAYER_SPEED
	player.Y += math.Sin(player.Direction) * PLAYER_SPEED
	if player.X < 0 {
		player.X = 0
	}
	if player.X > MAP_SIZE {
		player.X = MAP_SIZE
	}
	if player.Y < 0 {
		player.Y = 0
	}
	if player.Y > MAP_SIZE {
		player.Y = MAP_SIZE
	}
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
			UpdatePlayerPosition(player)
			UpdateBullets(player)
			bullets := make([]models.Bullet, 0, len(player.Bullets))
			for _, b := range player.Bullets {
				bullets = append(bullets, b)
			}
			msg.Players = append(msg.Players, models.PlayerSnapshot{
				Id:        player.Id,
				X:         player.X,
				Y:         player.Y,
				Direction: player.Direction,
				Bullets:   bullets,
			})
		}
		broadcast(*msg)
		ServerMessagePool.Put(msg)
	}
}
