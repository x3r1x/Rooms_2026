package ports

import "gamedevRooms/internal/domain"

type MapManager interface {
	GetRoomMessages() map[string]domain.RoomMessage
	GetRoomInfo(roomId string) interface{}
	GetExit(roomId, direction string) string
	LoadMapObjects()
}
