package game

import (
	"encoding/json"
	"gamedevRooms/internal/models"
	"log"
)

type SendResult struct {
	player *models.PlayerState
	err    error
}

func broadcast(message models.ServerMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}

	Mutex.RLock()
	for _, player := range Clients {
		select {
		case player.SendChan <- data:
		default:
		}
	}
	Mutex.RUnlock()
}
