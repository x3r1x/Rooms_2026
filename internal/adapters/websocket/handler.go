package websocket

import (
	"encoding/json"
	"fmt"
	"gamedevRooms/internal/application/lobby"
	"gamedevRooms/internal/domain"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	lobby *lobby.LobbyService
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWebsocketHandler(l *lobby.LobbyService) *WebsocketHandler {
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
	var currentId string

	defer func() {
		if currentId != "" {
			wsh.lobby.RemovePlayerFromLobby(currentId)
		}
	}()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения: ", err)
			break
		}

		switch wsh.lobby.GetState() {
		case domain.WaitingLobbyState:
			currentId = wsh.handleWaitingLobbyState(p, currentId, conn)
		case domain.OngoingGameState:
			wsh.handleOngoingGameState(p, currentId)
		}
	}
}

func (wsh *WebsocketHandler) handleWaitingLobbyState(p []byte, id string, conn *websocket.Conn) string {
	var registerMsg domain.ClientRegisterMessage

	if err := json.Unmarshal(p, &registerMsg); err == nil && id == "" {
		return wsh.lobby.AddPlayerToLobby(registerMsg.Nickname, conn)
	}

	var readyMsg domain.ClientReadyStateMessage

	if err := json.Unmarshal(p, &readyMsg); err == nil && id != "" {
		wsh.lobby.UpdatePlayerInLobby(&domain.LobbyPlayer{
			Id:    id,
			Ready: readyMsg.Ready,
		})
		return id
	}

	log.Println("Ошибка сериализации в waitingLobbyState!")

	if id != "" {
		log.Println("Удаляем пользователя: ", id)
		wsh.lobby.RemovePlayerFromLobby(id)
	}
	return id
}

func (wsh *WebsocketHandler) handleOngoingGameState(p []byte, currentId string) {
	var msg domain.ClientGameMessage

	if err := json.Unmarshal(p, &msg); err != nil {
		log.Println("Ошибка сериализации во время игры: ", err)
		return
	}

	if currentId != msg.Id {
		return
	}

	gameService := wsh.lobby.GetGameService()
	if gameService == nil {
		return
	}

	gameService.UpdatePlayer(domain.ClientGameMessage{
		Id:      msg.Id,
		MX:      msg.MX,
		MY:      msg.MY,
		Angle:   msg.Angle,
		IsShoot: msg.IsShoot,
	})
}
