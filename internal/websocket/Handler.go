package websocket

import (
	"encoding/json"
	"fmt"
	"gamedevRooms/internal/Lobby"
	"gamedevRooms/internal/model"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	lobby *Lobby.Lobby
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWebsocketHandler(l *Lobby.Lobby) *WebsocketHandler {
	return &WebsocketHandler{
		lobby: l,
	}
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
	currentId := ""

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения: ", err)
			break
		}

		switch wsh.lobby.GetState() {
		case model.WaitingLobbyState:
			wsh.handleWaitingLobbyState(p, &currentId)
		case model.OngoingGameState:
			wsh.handleOngoingGameState(p)
		}
	}
}

func (wsh *WebsocketHandler) handleWaitingLobbyState(p []byte, id *string) {
	var registerMsg model.ClientRegisterMessage

	if err := json.Unmarshal(p, &registerMsg); err == nil {
		*id = wsh.lobby.AddUser(registerMsg.Nickname)
		return
	}

	var readyMsg model.ClientReadyStateMessage

	if err := json.Unmarshal(p, &readyMsg); err == nil {
		wsh.lobby.SetUserReadyState(id, readyMsg.Ready)
		if wsh.lobby.CheckIfEveryoneReady() {
			wsh.lobby.SetState(model.ReadyLobbyState)
			//TODO: send message
		}

		return
	}

	log.Println("Ошибка сериализации в waitingLobbyState!")

	if *id != "" {
		log.Println("Удаляем пользователя: ", *id)
		wsh.lobby.RemoveUser(*id)
	}
}

func (wsh *WebsocketHandler) handleOngoingGameState(p []byte) {
	var msg model.ClientGameMessage

	if err := json.Unmarshal(p, &msg); err != nil {
		log.Println("Ошибка сериализации во время игры: ", err)
		log.Println("Удаляем пользователя ", msg.Id)
		wsh.lobby.GetGameLoop().DeletePlayer(msg.Id)

		return
	}

	wsh.lobby.GetGameLoop().UpdatePlayer(model.ClientGameMessage{
		Id:      msg.Id,
		MX:      msg.MX,
		MY:      msg.MY,
		Angle:   msg.Angle,
		IsShoot: msg.IsShoot,
	})
}
