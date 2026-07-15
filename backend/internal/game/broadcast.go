package game

import (
	"gamedevRooms/internal/models"
	"log"
)

type SendResult struct {
	player *models.PlayerState
	err    error
}

func broadcast(message models.ServerMessage) {
	data, err := models.MarshallJson(&message)
	if err != nil {
		log.Println(err)
		return
	}

	for _, player := range Clients {
		select {
		case player.SendChan <- data:
		default:
			continue
		}
	}
}
