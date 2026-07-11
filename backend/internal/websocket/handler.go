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
				fmt.Println("Client disconnected. ID: ", currentID)
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

		if currentID == "" && msg.Player.ID != "" {
			currentID = msg.Player.ID
			game.Clients[currentID] = &models.PlayerState{
				ID:        currentID,
				X:         msg.Player.X,
				Y:         msg.Player.Y,
				DIRECTION: msg.Player.DIRECTION,
				Bullets:   msg.Player.Bullets,
				Conn:      conn,
			}
		}

		if state, ok := game.Clients[currentID]; ok {
			state.X = msg.Player.X
			state.Y = msg.Player.Y
			state.DIRECTION = msg.Player.DIRECTION
			state.Bullets = msg.Player.Bullets
		}

		game.Mutex.Unlock()
	}
}
