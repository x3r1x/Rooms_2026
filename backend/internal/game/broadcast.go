package game

import (
	"encoding/json"
	"fmt"
	"gamedevRooms/internal/models"
	"log"

	"github.com/gorilla/websocket"
)

func broadcast(message models.ServerMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}

	Mutex.Lock()
	conns := make([]*models.PlayerState, 0, len(Clients))
	for _, player := range Clients {
		conns = append(conns, player)
	}
	Mutex.Unlock()
	var deadClients []*models.PlayerState
	for _, player := range conns {
		player.Mu.Lock()
		err := player.Conn.WriteMessage(websocket.TextMessage, data)
		player.Mu.Unlock()
		if err != nil {
			deadClients = append(deadClients, player)
		}
	}
	if len(deadClients) > 0 {
		Mutex.Lock()
		for _, deadClient := range deadClients {
			delete(Clients, deadClient.ID)
			fmt.Println("Client disconnected. ID: ", deadClient.ID)
			err = deadClient.Conn.Close()
			if err != nil {
				log.Println(err)
			}
		}
		Mutex.Unlock()
	}
}
