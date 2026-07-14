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
		writeIdx := 0
		for readIdx, _ := range player.Bullets {
			bullet := player.Bullets[readIdx]
			bullet.Life--
			if bullet.Life <= 0 {
				continue
			}
			bullet.X += math.Cos(bullet.Direction) * models.MAX_BULLET_SPEED
			bullet.Y += math.Sin(bullet.Direction) * models.MAX_BULLET_SPEED
			// обдумать условие
			if bullet.Y >= 0 && bullet.Y <= models.MAP_SIZE && bullet.X >= 0 && bullet.X <= models.MAP_SIZE {
				player.Bullets[id] = bullet
				writeIdx++
			}
		}
		//player.Bullets = player.Bullets[:id]
	}

}

func StartGameLoop() {
	ticker := time.NewTicker(models.TICK_TIME * time.Millisecond)
	//TODO: мапа на клиентов, чтобы избавиться от дублирования.
	//TODO: мапа на пули, для регулирования при попадании в стену и удалении из последовательности середины

	for range ticker.C {
		msg := ServerMessagePool.Get().(*models.AbsoluteServerMessage)

		Mutex.Lock()
		//TODO: вынести в функцию
		//TODO: организовать очищение памяти от мусора. опционально
		UpdateBullets()
		for id, player := range Clients {
			msg.Players[id] = *player
		}
		Mutex.Unlock()
		broadcastAbsolute(*msg)
		ServerMessagePool.Put(msg)
		PreviousClients = Clients
	}
}
