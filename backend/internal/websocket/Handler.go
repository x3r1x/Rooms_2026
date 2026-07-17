package websocket

import (
	"encoding/json"
	"fmt"
	"gamedevRooms/internal/game"
	"gamedevRooms/internal/model"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	gameState *game.GameState
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWebsocketHandler(gs *game.GameState) *WebsocketHandler {
	return &WebsocketHandler{gameState: gs}
}

func (wsh *WebsocketHandler) InitWebsocket(w http.ResponseWriter, r *http.Request) {
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
	wsh.HandleWebsocket(conn)
}

func (wsh *WebsocketHandler) HandleWebsocket(conn *websocket.Conn) {
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
				wsh.gameState.DeleteChan <- msg.Id
			}
			continue
		}

		if currentId == "" && msg.Id != "" {
			isRegistered = true
			currentId = msg.Id
			fmt.Println("Регистрация пользователя")
			wsh.gameState.RegisterChan <- &model.PlayerState{
				Id:         msg.Id,
				X:          model.PlayerSpawnPointX,
				Y:          model.PlayerSpawnPointY,
				Angle:      model.InitDirection,
				MoveX:      model.InitDirection,
				MoveY:      model.InitDirection,
				Connection: conn,
			}
		} else if currentId != "" {
			wsh.gameState.UpdateChan <- model.ClientMessage{
				Id:      msg.Id,
				MX:      msg.MX,
				MY:      msg.MY,
				Angle:   msg.Angle,
				IsShoot: msg.IsShoot,
			}
		}
	}
	if currentId != "" {
		wsh.gameState.DeleteChan <- currentId
	}
}
