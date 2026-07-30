package websocket

import (
	"encoding/json"
	"fmt"
	"gamedevRooms/internal/application/lobby"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/recovery"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const readDeadlineLobby = 90 * time.Second
const readDeadlineGame = 20 * time.Second

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
	fmt.Println("New connection established.")
	wsh.HandleWebsocket(conn)
}

func (wsh *WebsocketHandler) HandleWebsocket(conn *websocket.Conn) {
	defer recovery.Recover()
	var currentId string
	if err := conn.SetReadDeadline(time.Now().Add(readDeadlineLobby)); err != nil {
		return
	}
	defer func() {
		if currentId != "" {
			wsh.lobby.RemovePlayerFromLobby(currentId, false)
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
			if err := conn.SetReadDeadline(time.Now().Add(readDeadlineLobby)); err != nil {
				return
			}
			currentId = wsh.handleWaitingLobbyState(p, currentId, conn)
		case domain.OngoingGameState:
			if err := conn.SetReadDeadline(time.Now().Add(readDeadlineGame)); err != nil {
				return
			}
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
			Id:          id,
			Ready:       readyMsg.Ready,
			PlayerClass: readyMsg.PlayerClass,
		})
		return id
	}

	log.Println("Ошибка сериализации в waitingLobbyState!")

	if id != "" {
		log.Println("Удаляем пользователя: ", id)
		wsh.lobby.RemovePlayerFromLobby(id, true)
	}
	return id
}

func (wsh *WebsocketHandler) handleOngoingGameState(p []byte, currentId string) {
	var msg domain.ClientGameMessage
	if err := json.Unmarshal(p, &msg); err != nil {
		log.Println("Ошибка сериализации во время игры: ", err)
		return
	}

	if currentId == "" {
		log.Println("ОШИБКА: currentId пустой!")
		return
	}

	if currentId != msg.Id {
		log.Printf("Несоответствие ID: currentId=%s, msg.Id=%s", currentId, msg.Id)
		return
	}

	gameService := wsh.lobby.GetGameService()
	if gameService == nil {
		log.Println("ОШИБКА: GameService is nil!")
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
