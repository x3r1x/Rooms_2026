package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type ClientMessage struct {
	Player struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"player"`
}

type PlayerState struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Nickname string  `json:"nickname"`
}

type ServerMessage struct {
	Players []PlayerState `json:"players"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]*PlayerState)
var mutex sync.Mutex
var playerCounter int

// TODO: постепенно перевести на обмен данных между клиентом и сервером на сырые байты.
// TODO: уменьшить обмен, чтобы уменьшить нагрузку
// TODO: разбить на файлы, для нормального поддержания

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
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
	playerCounter++
	newState := &PlayerState{
		Nickname: fmt.Sprintf("player%d", playerCounter),
	}
	mutex.Lock()
	clients[conn] = newState
	mutex.Unlock()

	fmt.Println("New user is here. Players on server", len(clients))
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			fmt.Println("Closing connection. Player leave")
			break
		}

		var msg ClientMessage
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println(err)
			continue
		}
		mutex.Lock()
		if state, ok := clients[conn]; ok {
			state.X = msg.Player.X
			state.Y = msg.Player.Y
		}
		serverMsg := ServerMessage{
			Players: make([]PlayerState, 0, len(clients)),
		}
		for _, client := range clients {
			serverMsg.Players = append(serverMsg.Players, *client)
		}
		mutex.Unlock()

		broadcast(serverMsg)
	}
}

func broadcast(message ServerMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println(err)
			client.Close()
			delete(clients, client)
		}
	}
}

// scp -i "C:\Users\maxim\.ssh\id_ed25519.pub" server yc-user@84.201.159.214:/home/yc-user/
func main() {
	http.HandleFunc("/ws", handleWebsocket)
	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
