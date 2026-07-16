package websocket

import (
	"encoding/json"
	"fmt"
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
	var currentId string
	var isRegistered bool
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения: ", err)
			break
		}

		var msg models.ClientMessage
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println("Ошибка сериализации при чтении пакета: ", err)
			if isRegistered {
				log.Println("Удаляем пользователя: ", msg.Player.Id)
				models.Game.LeaveChan <- msg.Player.Id
			}
			continue
		}

		if currentId == "" && msg.Player.Id != "" {
			isRegistered = true
			currentId = msg.Player.Id
			fmt.Println("Регистрация пользователя")
			models.Game.RegisterChan <- models.PlayerState{
				Id:         msg.Player.Id,
				X:          msg.Player.X,
				Y:          msg.Player.Y,
				Direction:  msg.Player.Direction,
				Bullets:    msg.Player.Bullets,
				Connection: conn,
			}
		}
		if currentId != "" {
			models.Game.InputChan <- models.ClientMessage{
				Player: msg.Player,
			}
		}
	}
	if currentId != "" {
		models.Game.LeaveChan <- currentId
	}
}
