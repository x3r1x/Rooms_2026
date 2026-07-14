package infrastructure

//
//import (
//	"encoding/json"
//	"fmt"
//	"gamedevRooms/internal/game"
//	"gamedevRooms/internal/models"
//	"log"
//
//	"github.com/gorilla/websocket"
//)
//
//func ClientHandler(conn *websocket.Conn, players map[string]*models.PlayerState) {
//	var currentID string
//
//	for {
//		_, p, err := conn.ReadMessage()
//		if err != nil {
//			log.Println(err)
//			game.Mutex.Lock()
//			if currentID != "" {
//				delete(players, currentID)
//				fmt.Println("Client disconnected. Id: ", currentID)
//			}
//			game.Mutex.Unlock()
//			break
//		}
//		var msg models.ClientMessage
//		if err := json.Unmarshal(p, &msg); err != nil {
//			log.Println(err)
//			continue
//		}
//
//		game.Mutex.Lock()
//
//		if currentID == "" && msg.Player.Id != "" {
//			currentID = msg.Player.Id
//			game.Clients[currentID] = &models.PlayerState{
//				Id:        currentID,
//				X:         msg.Player.X,
//				Y:         msg.Player.Y,
//				Direction: msg.Player.Direction,
//				Bullets:   msg.Player.Bullets,
//				Conn:      conn,
//			}
//			game.Mutex.Unlock()
//			continue
//		}
//
//		if currentID != "" {
//			if _, ok := game.Clients[currentID]; ok {
//				game.UpdatePlayerState(currentID, msg.Player.X, msg.Player.Y, msg.Player.Direction, msg.Player.Bullets)
//			}
//		}
//
//		game.Mutex.Unlock()
//	}
//}
