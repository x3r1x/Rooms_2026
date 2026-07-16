package game

import (
	"encoding/json"
	"gamedevRooms/internal/models"
	"log"

	"github.com/gorilla/websocket"
)

func broadcast(message models.ServerMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}
	for id, p := range models.Game.Players {
		err := p.Connection.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Ошибка отправки: ", err)
			models.Game.LeaveChan <- id
		}
	}
}
