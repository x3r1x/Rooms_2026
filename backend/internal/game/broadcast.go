package game

import (
	"encoding/json"
	"gamedevRooms/internal/models"
	"log"
)

func broadcast(message models.ServerMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}
}
