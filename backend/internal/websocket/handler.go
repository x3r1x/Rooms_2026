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

		if currentID == "" && msg.PlayerInterpolation.Id != "" {
			currentID = msg.PlayerInterpolation.Id
			game.Clients[currentID] = &models.PlayerState{
				X:         models.PLAYER_START_X,
				Y:         models.PLAYER_START_Y,
				Direction: models.PLAYER_START_DIRECTION,
				Bullets:   map[string]models.Bullet{},
				Conn:      conn,
			}
			game.Mutex.Unlock()
			continue
		}

		if currentID != "" {
			if _, ok := game.Clients[currentID]; ok {
				var playerMovementDirection game.PlayerMovementDirection

				playerMovementDirection = game.PlayerMovementDirection{
					X: float64(msg.PlayerInterpolation.Direction.X),
					Y: float64(msg.PlayerInterpolation.Direction.Y),
				}

				game.UpdatePlayerState(currentID, playerMovementDirection, msg.PlayerInterpolation.DeltaVisualDirection,
					msg.PlayerInterpolation.NewBulletsDirection)
			}
		}

		game.Mutex.Unlock()
	}
}
