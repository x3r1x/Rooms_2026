package websocket

import (
	"encoding/json"
	"fmt"
	"gamedevRooms/internal/game"
	"gamedevRooms/internal/models"
	"log"
	"net/http"
	"time"

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
	fmt.Println("New connection established.")

	sendChan := make(chan []byte, 128)
	var currentID string

	go func() {
		defer func() {
			if err := conn.Close(); err != nil {
				log.Println(err)
			}
		}()
		for data := range sendChan {
			err = conn.SetWriteDeadline(time.Now().Add(time.Second * 3))
			if err != nil {
				log.Println(err)
				return
			}
			if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
				return
			}
		}
	}()

	defer func() {
		if currentID != "" {
			game.CommandChan <- game.Command{
				Type: game.DisconnectPlayer,
				ID:   currentID,
			}
			fmt.Println("Disconnected client:", currentID)
		} else {
			close(sendChan)
		}
	}()

	for {
		//TODO: добавить коллизии
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		var msg models.ClientMessage
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println(err)
			continue
		}

		if currentID == "" && msg.Player.Id != "" {
			currentID = msg.Player.Id
			game.CommandChan <- game.Command{
				Type: game.RegisterPlayer,
				ID:   currentID,
				Player: &models.PlayerState{
					Id:        currentID,
					X:         msg.Player.X,
					Y:         msg.Player.Y,
					Direction: msg.Player.Direction,
					Bullets:   []models.Bullet{},
					Conn:      conn,
					SendChan:  sendChan,
				},
			}
			continue
		}

		if currentID != "" {
			cmd := game.Command{
				Type:    game.UpdatePlayer,
				ID:      currentID,
				X:       msg.Player.X,
				Y:       msg.Player.Y,
				Dir:     msg.Player.Direction,
				Bullets: msg.Player.Bullets,
			}
			select {
			case game.CommandChan <- cmd:
			default:
			}
		}
	}
}
