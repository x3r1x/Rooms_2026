package game

import (
	"gamedevRooms/internal/models"
	"time"
)

func StartGameLoop() {
	ticker := time.NewTicker(16 * time.Millisecond)
	for range ticker.C {
		Mutex.Lock()
		serverMessage := models.ServerMessage{
			Players: make([]models.PlayerState, 0, len(Clients)),
		}
		for _, player := range Clients {
			serverMessage.Players = append(serverMessage.Players, models.PlayerState{
				X: player.X, Y: player.Y, DIRECTION: player.DIRECTION, ID: player.ID, Bullets: player.Bullets,
			})
		}
		Mutex.Unlock()
		broadcast(serverMessage)
	}
}
