package ports

type BroadcastService interface {
	AddConnection(playerId string, conn interface{})
	RemoveConnection(playerId string)
	BroadcastToAll(message interface{})
	BroadcastToPlayer(playerId string, message interface{})
}
