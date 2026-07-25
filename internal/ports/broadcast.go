package ports

import "github.com/gorilla/websocket"

type BroadcastService interface {
	AddConnection(playerId string, conn *websocket.Conn)
	RemoveConnection(playerId string)
	BroadcastToAll(message interface{})
	BroadcastToPlayer(playerId string, message interface{})
}
