package broadcast

import (
	"encoding/json"
	"gamedevRooms/internal/recovery"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const writeDeadline = 3 * time.Second

type BroadcastService struct {
	addConnChan    chan connectionEvent
	removeConnChan chan string
	broadcastChan  chan broadcastEvent
	toPlayerChan   chan playerEvent
	connections    map[string]*websocket.Conn
}

type connectionEvent struct {
	playerId string
	conn     *websocket.Conn
}

type broadcastEvent struct {
	message interface{}
}

type playerEvent struct {
	playerId string
	message  interface{}
}

func NewBroadcastService() *BroadcastService {
	bs := &BroadcastService{
		addConnChan:    make(chan connectionEvent),
		removeConnChan: make(chan string),
		broadcastChan:  make(chan broadcastEvent),
		toPlayerChan:   make(chan playerEvent),
		connections:    make(map[string]*websocket.Conn),
	}
	go bs.run()
	return bs
}

func (bs *BroadcastService) run() {
	defer recovery.Recover()
	for {
		select {
		case event := <-bs.addConnChan:
			bs.connections[event.playerId] = event.conn
			log.Printf("Добавлено соединение для игрока %s", event.playerId)

		case playerId := <-bs.removeConnChan:
			if conn, exists := bs.connections[playerId]; exists {
				conn.Close()
				delete(bs.connections, playerId)
				log.Printf("Удалено соединение для игрока %s", playerId)
			}

		case event := <-bs.broadcastChan:
			bs.sendToAll(event.message)

		case event := <-bs.toPlayerChan:
			bs.sendToPlayer(event.playerId, event.message)
		}
	}
}

func (bs *BroadcastService) AddConnection(playerId string, conn *websocket.Conn) {
	bs.addConnChan <- connectionEvent{
		playerId: playerId,
		conn:     conn,
	}
}

func (bs *BroadcastService) RemoveConnection(playerId string) {
	bs.removeConnChan <- playerId
}

func (bs *BroadcastService) BroadcastToAll(message interface{}) {
	bs.broadcastChan <- broadcastEvent{message: message}
}

func (bs *BroadcastService) BroadcastToPlayer(playerId string, message interface{}) {
	bs.toPlayerChan <- playerEvent{
		playerId: playerId,
		message:  message,
	}
}

func (bs *BroadcastService) sendToAll(message interface{}) {
	defer recovery.Recover()
	data, err := json.Marshal(message)
	if err != nil {
		log.Println("Ошибка маршалинга:", err)
		return
	}

	for id, conn := range bs.connections {
		conn.SetWriteDeadline(time.Now().Add(writeDeadline))
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Printf("Ошибка отправки игроку %s: %v", id, err)
			bs.RemoveConnection(id)
		}
	}
}

func (bs *BroadcastService) sendToPlayer(playerId string, message interface{}) {
	defer recovery.Recover()

	var data []byte
	var err error
	if b, ok := message.([]byte); ok {
		data = b
	} else {
		data, err = json.Marshal(message)
		if err != nil {
			log.Println("Ошибка маршалинга:", err)
			return
		}
	}
	conn, exists := bs.connections[playerId]
	if !exists {
		log.Printf("Соединение для игрока %s не найдено", playerId)
		return
	}
	conn.SetWriteDeadline(time.Now().Add(writeDeadline))
	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Printf("Ошибка отправки игроку %s: %v", playerId, err)
		bs.RemoveConnection(playerId)
	}
}
