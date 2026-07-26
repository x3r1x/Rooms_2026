package _map

import (
	"gamedevRooms/internal/adapters/map/generator"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/ports"
	"strconv"

	"github.com/google/uuid"
)

type MapManager struct {
	gameMap  *generator.Map
	tileSize float64
}

func NewMapManager() *MapManager {
	//gameMap := mapModel.NewMap(gameState.GetCountOfPlayers()/3 + 1)
	gameMap := generator.NewMap(4)
	return &MapManager{
		gameMap:  gameMap,
		tileSize: domain.TileSize,
	}
}

func (mm *MapManager) LoadMapObjects(gameState ports.GameStateProvider) {
	objects := make(map[string]*domain.Object)
	solidCount := 0

	for roomId, room := range mm.gameMap.GetGameMap() {
		matrix := room.GetMatrix()
		for i, tile := range matrix {
			if tile == nil {
				continue
			}

			isSolid := false
			for _, prop := range tile.Properties {
				if prop.Name == "blocksPlayer" {
					if val, ok := prop.Value.(bool); ok && val {
						isSolid = true
						break
					}
				}
			}

			if isSolid {
				row := i / domain.RoomWidth
				col := i % domain.RoomWidth
				obj := &domain.Object{
					Id:      uuid.New().String(),
					X:       float64(col) * mm.tileSize,
					Y:       float64(row) * mm.tileSize,
					Width:   mm.tileSize,
					Height:  mm.tileSize,
					IsSolid: true,
					Type:    "wall",
					RoomId:  roomId,
				}
				objects[obj.Id] = obj
				solidCount++
			}
		}
	}
	gameState.SetObjects(objects)
}

func (mm *MapManager) GetRoomMessages() map[string]domain.RoomMessage {
	result := make(map[string]domain.RoomMessage)
	for id, room := range mm.gameMap.GetGameMap() {
		barrierNum := 1
		barrierStr := room.GetBarrierType()
		if len(barrierStr) > 7 {
			if num, err := strconv.Atoi(barrierStr[7:]); err == nil {
				barrierNum = num
			}
		}
		result[id] = domain.RoomMessage{
			ExitTop:    room.GetExit(domain.TopMarker),
			ExitLeft:   room.GetExit(domain.LeftMarker),
			ExitBottom: room.GetExit(domain.BottomMarker),
			ExitRight:  room.GetExit(domain.RightMarker),
			Id:         room.GetId(),
			BorderType: barrierNum,
		}
	}
	return result
}

func (mm *MapManager) GetRoomInfo(roomId string) *generator.Room {
	if mm.gameMap == nil {
		return nil
	}
	rooms := mm.gameMap.GetGameMap()
	if rooms == nil {
		return nil
	}
	return rooms[roomId]
}

func (mm *MapManager) GetExit(roomId, direction string) string {
	room := mm.GetRoomInfo(roomId)
	if room == nil {
		return ""
	}
	return room.GetExit(direction)
}
