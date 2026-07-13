package websocket

import (
	"encoding/json"
	"fmt"
	"gamedevRooms/internal/game"
	"gamedevRooms/internal/models"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}()
	fmt.Println("New connection established.")

	var currentID string

	for {
		//TODO: добавить коллизии
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			game.Mutex.Lock()
			if currentID != "" {
				delete(game.Clients, currentID)
				fmt.Println("Client disconnected. Id: ", currentID)
			}
			game.Mutex.Unlock()
			break
		}
		var msg models.ClientMessage
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println(err)
			continue
		}

		game.Mutex.Lock()

		if currentID == "" && msg.Player.Id != "" {
			currentID = msg.Player.Id
			game.Clients[currentID] = &models.PlayerState{
				Id:        currentID,
				X:         msg.Player.X,
				Y:         msg.Player.Y,
				Direction: msg.Player.Direction,
				Bullets:   msg.Player.Bullets,
				Conn:      conn,
			}
			game.Mutex.Unlock()
			continue
		}

		if currentID != "" {
			if _, ok := game.Clients[currentID]; !ok {
				game.UpdatePlayerState(currentID, msg.Player.X, msg.Player.Y, msg.Player.Direction, msg.Player.Bullets)
			}
		}

		game.Mutex.Unlock()
	}
}
