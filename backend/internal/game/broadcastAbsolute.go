package game

import (
	"encoding/json"
	"gamedevRooms/internal/models"
	"log"
)

func broadcastAbsolute(message models.AbsoluteServerMessage) {
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
