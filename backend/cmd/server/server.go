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
		X  float64 `json:"x"`
		Y  float64 `json:"y"`
		ID string  `json:"id"`
	} `json:"player"`
}

type PlayerState struct {
	X    float64         `json:"x"`
	Y    float64         `json:"y"`
	ID   string          `json:"id"`
	Conn *websocket.Conn `json:"-"`
}

type ServerMessage struct {
	Players []PlayerState `json:"players"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[string]*PlayerState)
var mutex sync.Mutex

// TODO: постепенно перевести на обмен данных между клиентом и сервером на сырые байты.
// TODO: уменьшить обмен, чтобы уменьшить нагрузку
// TODO: разбить на файлы, для нормального поддержания
// TODO: нужна обратная мапа для подключений для упрощения поиска и мониторинга
// TODO: коллизия для пуль и их непосредственная обработка
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

	var currentID string

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			mutex.Lock()
			if currentID != "" {
				delete(clients, currentID)
				fmt.Println("Client disconnected. ID: ", currentID)
			}
			mutex.Unlock()
			break
		}
		var msg ClientMessage
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println(err)
			continue
		}

		mutex.Lock()
		if currentID == "" && msg.Player.ID != "" {
			currentID = msg.Player.ID
			clients[currentID] = &PlayerState{
				ID:   currentID,
				Conn: conn,
			}
		}
		if state, ok := clients[currentID]; ok {
			state.X = msg.Player.X
			state.Y = msg.Player.Y
		}
		serverMessage := ServerMessage{
			Players: make([]PlayerState, 0, len(clients)),
		}
		for _, player := range clients {
			serverMessage.Players = append(serverMessage.Players, PlayerState{
				X: player.X, Y: player.Y, ID: player.ID,
			})
		}
		mutex.Unlock()
		broadcast(serverMessage)
	}
}

func broadcast(message ServerMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}

	mutex.Lock()
	conns := make([]*websocket.Conn, 0, len(clients))
	for _, player := range clients {
		conns = append(conns, player.Conn)
	}
	mutex.Unlock()

	for _, conn := range conns {
		err := conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			mutex.Lock()
			for id, p := range clients {
				if p.Conn == conn {
					delete(clients, id)

					fmt.Println("Client disconnected. ID: ", id)
					break
				}
			}
			mutex.Unlock()
			err = conn.Close()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// scp -i "C:\Users\maxim\.ssh\id_ed25519.pub" server yc-user@84.201.159.214:/home/yc-user/
func main() {
	http.HandleFunc("/ws", handleWebsocket)
	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
