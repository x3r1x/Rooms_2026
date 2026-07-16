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

func InitWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("Ошибка закрытия связи: ", err)
		}
	}()
	fmt.Println("New connection established.")
	HandleWebsocket(conn)
}

func HandleWebsocket(conn *websocket.Conn) {

	var currentID string
	var GameState models.GameState

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения: ", err)
			break
		}
		var msg models.ClientMessage
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println("Ошибка сериализации: ", err)
			continue
		}

		if currentID == "" && msg.Player.Id != "" {
			currentID = msg.Player.Id
			GameState.Players[currentID] = &models.PlayerState{
				Id:        currentID,
				X:         msg.Player.X,
				Y:         msg.Player.Y,
				Direction: msg.Player.Direction,
				Bullets:   msg.Player.Bullets,
			}
		}

		if state, ok := game.Clients[currentID]; ok {
			state.X = msg.Player.X
			state.Y = msg.Player.Y
			state.Direction = msg.Player.Direction
			state.Bullets = msg.Player.Bullets
		}

	}
}
