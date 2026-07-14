package game

import (
	"encoding/json"
	"fmt"
	"gamedevRooms/internal/models"
	"log"

	"github.com/gorilla/websocket"
)

func broadcastAbsolute(message models.AbsoluteServerMessage) {
	message.Type = "absolute"
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}

	Mutex.Lock()
	conns := make(map[string]*models.PlayerState)
	for id, player := range Clients {
		conns[id] = player
	}
	Mutex.Unlock()
	var deadClients map[string]*models.PlayerState
	for id, player := range conns {
		player.Mu.Lock()
		err := player.Conn.WriteMessage(websocket.TextMessage, data)
		player.Mu.Unlock()
		if err != nil {
			deadClients[id] = player
		}
	}
	if len(deadClients) > 0 {
		Mutex.Lock()
		for id, deadClient := range deadClients {
			delete(Clients, id)
			fmt.Println("Client disconnected. Id: ", id)
			err = deadClient.Conn.Close()
			if err != nil {
				log.Println(err)
			}
		}
		Mutex.Unlock()
	}
}
