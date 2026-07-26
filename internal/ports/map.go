package ports

import (
	"gamedevRooms/internal/adapters/map/generator"
	"gamedevRooms/internal/domain"
)

type MapManager interface {
	GetRoomMessages() map[string]domain.RoomMessage
	GetRoomInfo(roomId string) *generator.Room
	GetExit(roomId, direction string) string
	LoadMapObjects(gameState GameStateProvider)
}
