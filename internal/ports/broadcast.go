package ports

type BroadcastService interface {
	BroadcastToAll(message interface{})
	BroadcastToPlayer(playerId string, message interface{})
	BroadcastToRoom(roomId string, message interface{})
}
