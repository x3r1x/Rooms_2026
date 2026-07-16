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
	var currentId string = ""
	var isRegistered bool = false
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения: ", err)
			break
		}

		var msg models.ClientMessage
		fmt.Println(string(p))
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println("Ошибка сериализации при чтении пакета: ", err)
			if isRegistered {
				log.Println("Удаляем пользователя: ", msg.Id)
				models.Game.LeaveChan <- msg.Id
			}
			continue
		}

		if currentId == "" && msg.Id != "" {
			isRegistered = true
			currentId = msg.Id
			fmt.Println("Регистрация пользователя")
			models.Game.RegisterChan <- models.PlayerState{
				Id:         msg.Id,
				X:          models.PlayerSpawnPointX,
				Y:          models.PlayerSpawnPointY,
				A:          models.InitDirection,
				MoveX:      models.InitDirection,
				MoveY:      models.InitDirection,
				Connection: conn,
			}
		} else if currentId != "" {
			models.Game.InputChan <- models.ClientMessage{
				Id: msg.Id,
				MX: msg.MX,
				MY: msg.MY,
				A:  msg.A,
				S:  msg.S,
			}
		}
	}
	if currentId != "" {
		models.Game.LeaveChan <- currentId
	}
}
