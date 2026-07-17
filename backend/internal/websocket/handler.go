package websocket

import (
	"encoding/json"
	"fmt"
	"gamedevRooms/internal/model"
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

		var msg model.ClientMessage
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println("Ошибка сериализации при чтении пакета: ", err)
			if isRegistered {
				log.Println("Удаляем пользователя: ", msg.Id)
				model.Game.LeaveChan <- msg.Id
			}
			continue
		}

		if currentId == "" && msg.Id != "" {
			isRegistered = true
			currentId = msg.Id
			fmt.Println("Регистрация пользователя")
			model.Game.RegisterChan <- model.PlayerState{
				Id:         msg.Id,
				X:          model.PlayerSpawnPointX,
				Y:          model.PlayerSpawnPointY,
				A:          model.InitDirection,
				MoveX:      model.InitDirection,
				MoveY:      model.InitDirection,
				Connection: conn,
			}
		} else if currentId != "" {
			model.Game.InputChan <- model.ClientMessage{
				Id: msg.Id,
				MX: msg.MX,
				MY: msg.MY,
				A:  msg.A,
				S:  msg.S,
			}
		}
	}
	if currentId != "" {
		model.Game.LeaveChan <- currentId
	}
}
