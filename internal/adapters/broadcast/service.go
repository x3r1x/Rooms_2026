package broadcast

import (
	"encoding/json"
	"gamedevRooms/internal/recovery"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const writeDeadline = 5 * time.Second

type BroadcastService struct {
	mu             sync.RWMutex
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
		removeConnChan: make(chan string, 100),
		broadcastChan:  make(chan broadcastEvent),
		toPlayerChan:   make(chan playerEvent, 1000),
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
			bs.mu.Lock()
			bs.connections[event.playerId] = event.conn
			bs.mu.Unlock()
			log.Printf("Добавлено соединение для игрока %s", event.playerId)

		case playerId := <-bs.removeConnChan:
			if conn, exists := bs.connections[playerId]; exists {
				if err := conn.Close(); err != nil {
					log.Printf("Ошибка закрытия соединения для игрока %s: %v", playerId, err)
				}
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
	bs.mu.RLock()
	conns := make(map[string]*websocket.Conn, len(bs.connections))
	for k, v := range bs.connections {
		conns[k] = v
	}
	bs.mu.RUnlock()

	for id, conn := range conns {
		go func(pid string, c *websocket.Conn, d []byte) {
			defer recovery.Recover()
			if err := c.SetWriteDeadline(time.Now().Add(writeDeadline)); err != nil {
				bs.RemoveConnection(pid)
				return
			}
			if err := c.WriteMessage(websocket.TextMessage, d); err != nil {
				bs.RemoveConnection(pid)
			}
		}(id, conn, data)
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
	bs.mu.RLock()
	conn, exists := bs.connections[playerId]
	bs.mu.RUnlock()
	if !exists {
		log.Printf("Соединение для игрока %s не найдено", playerId)
		return
	}
	go func(c *websocket.Conn, pid string, d []byte) {
		defer recovery.Recover()
		if err := c.SetWriteDeadline(time.Now().Add(writeDeadline)); err != nil {
			bs.RemoveConnection(pid)
			return
		}
		if err := c.WriteMessage(websocket.TextMessage, d); err != nil {
			bs.RemoveConnection(pid)
		}
	}(conn, playerId, data)
}
