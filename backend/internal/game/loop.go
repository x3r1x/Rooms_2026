package game

import (
	"gamedevRooms/internal/models"
	"math"
	"sync"
	"time"
)

var ServerMessagePool = sync.Pool{
	New: func() interface{} {
		return &models.AbsoluteServerMessage{
			Players: make(map[string]models.PlayerState),
		}
	},
}

func UpdateBullets() {
	for id, player := range Clients {
		newBullets := make(map[string]models.Bullet)
		for readIdx, bullet := range player.Bullets {
			bullet.Life--
			if bullet.Life <= 0 {
				continue
			}
			bullet.X += math.Cos(bullet.Direction) * models.MAX_BULLET_SPEED
			bullet.Y += math.Sin(bullet.Direction) * models.MAX_BULLET_SPEED
			if bullet.Y >= 0 && bullet.Y <= models.MAP_SIZE && bullet.X >= 0 && bullet.X <= models.MAP_SIZE {
				newBullets[readIdx] = bullet
			}
		}
		player.Bullets = newBullets
		Clients[id] = player
	}
}

func StartGameLoop() {
	ticker := time.NewTicker(models.TICK_TIME * time.Millisecond)
	//TODO: мапа на клиентов, чтобы избавиться от дублирования.

	for range ticker.C {
		msg := ServerMessagePool.Get().(*models.AbsoluteServerMessage)

		for k := range msg.Players {
			delete(msg.Players, k)
		}

		Mutex.Lock()
		UpdateBullets()
		for id, player := range Clients {
			msg.Players[id] = *player
		}
		Mutex.Unlock()
		broadcastAbsolute(*msg)
		//PreviousClients = Clients
		ServerMessagePool.Put(msg)
	}
}
