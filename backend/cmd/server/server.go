package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type BulletSpawn struct {
	VX float64 `json:"vx"`
	VY float64 `json:"vy"`
}

type ClientMessage struct {
	Player struct {
		X         float64       `json:"x"`
		Y         float64       `json:"y"`
		DIRECTION float64       `json:"direction"`
		ID        string        `json:"id"`
		Bullets   []BulletSpawn `json:"bullets"`
	} `json:"player"`
}

type BulletState struct {
	ID    string  `json:"id"`
	Owner string  `json:"owner"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	VX    float64 `json:"vx"`
	VY    float64 `json:"vy"`
	Life  int     `json:"-"`
}

type PlayerState struct {
	X         float64         `json:"x"`
	Y         float64         `json:"y"`
	DIRECTION float64         `json:"direction"`
	ID        string          `json:"id"`
	Bullets   []BulletState   `json:"bullets"`
	Conn      *websocket.Conn `json:"-"`
	Mu        sync.Mutex      `json:"-"`
}

type ServerMessage struct {
	Players []PlayerState `json:"players"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	clients     = make(map[string]*PlayerState)
	bullets     = make(map[string]*BulletState)
	mutex       sync.Mutex
	BULLET_TIME = 60
)

const MAX_BULLET_SPEED = 15.0

// TODO: уменьшить объем обмена, чтобы уменьшить нагрузку
// TODO: постепенно перевести на обмен данных между клиентом и сервером на сырые байты.
// TODO: разбить на файлы, для нормального поддержания
// TODO: нужна обратная мапа для подключений для упрощения поиска и мониторинга
// TODO: коллизия для пуль и их непосредственная обработка
// TODO: создать отдельный цикл на обработку пуль и движения, наконец добиться нормального распределения
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
		//TODO: добавить коллизии
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
		for _, spawn := range msg.Player.Bullets {
			var spawnX, spawnY float64
			if state, ok := clients[currentID]; ok {
				spawnX = state.X
				spawnY = state.Y
			} else {
				spawnX = msg.Player.X

			}
			//TODO: отредачить создание пуль, на атомик
			//TODO: решить проблему спавна
			bulletID := fmt.Sprintf("%s_b_%d", msg.Player.ID, time.Now().UnixNano())
			bullets[bulletID] = &BulletState{
				ID:    bulletID,
				Owner: msg.Player.ID,
				X:     spawnX,
				Y:     spawnY,
				VX:    spawn.VX,
				VY:    spawn.VY,
				Life:  BULLET_TIME,
			}
			if state, ok := clients[currentID]; ok {
				state.X = msg.Player.X
				state.Y = msg.Player.Y
				state.DIRECTION = msg.Player.DIRECTION
			}
		}

		mutex.Unlock()
	}
}

func broadcast(message ServerMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}

	mutex.Lock()
	conns := make([]*PlayerState, 0, len(clients))
	for _, player := range clients {
		conns = append(conns, player)
	}
	mutex.Unlock()
	deadClients := []*PlayerState{}
	for _, player := range conns {
		player.Mu.Lock()
		err := player.Conn.WriteMessage(websocket.TextMessage, data)
		player.Mu.Unlock()
		if err != nil {
			deadClients = append(deadClients, player)
		}
	}
	if len(deadClients) > 0 {
		mutex.Lock()
		for _, deadClient := range deadClients {
			delete(clients, deadClient.ID)
			fmt.Println("Client disconnected. ID: ", deadClient.ID)
			err = deadClient.Conn.Close()
			if err != nil {
				log.Println(err)
			}
		}
		mutex.Unlock()
	}
}

func gameLoop() {
	ticker := time.NewTicker(16 * time.Millisecond)
	for range ticker.C {
		mutex.Lock()
		for id, b := range bullets {
			b.X += b.VX
			b.Y += b.VY
			b.Life--
			if b.Life <= 0 {
				delete(bullets, id)
			}
		}
		serverMessage := ServerMessage{
			Players: make([]PlayerState, 0, len(clients)),
		}
		for _, player := range clients {
			myBullets := make([]BulletState, 0)
			for _, bullet := range bullets {
				if bullet.Owner == player.ID {
					myBullets = append(myBullets, *bullet)
				}
			}
			serverMessage.Players = append(serverMessage.Players, PlayerState{
				X: player.X, Y: player.Y, DIRECTION: player.DIRECTION, ID: player.ID, Bullets: myBullets,
			})

		}
		mutex.Unlock()
		broadcast(serverMessage)
	}
}

// scp -i "C:\Users\maxim\.ssh\id_ed25519.pub" server yc-user@84.201.159.214:/home/yc-user/
func main() {
	go gameLoop()
	http.HandleFunc("/ws", handleWebsocket)
	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
