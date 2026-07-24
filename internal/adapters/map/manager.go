package

import (
	"gamedevRooms/internal/adapters/map/generator"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/game"
	"log"
	"strconv"

	"github.com/google/uuid"
)
map

import (
	"gamedevRooms/internal/adapters/map/generator"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/game"
	"log"
	"strconv"

	"github.com/google/uuid"
)

import (
	"gamedevRooms/internal/model"
)

type MapManager struct {
	gameMap   *generator.Map
	gameState *game.GameState
	tileSize  float64
}

func NewMapManager(gameState *game.GameState) *MapManager {
	//gameMap := mapModel.NewMap(gameState.GetCountOfPlayers()/3 + 1)
	gameMap := generator.NewMap(1)
	mm := &MapManager{
		gameMap:   gameMap,
		gameState: gameState,
		tileSize:  domain.TileSize,
	}
	mm.loadMapObjects()
	return mm
}

func (mm *MapManager) loadMapObjects() {
	objects := make(map[string]*model.Object)
	solidCount := 0

	for _, room := range mm.gameMap.GetGameMap() {
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
				obj := &model.Object{
					Id:      uuid.New().String(),
					X:       float64(col) * mm.tileSize,
					Y:       float64(row) * mm.tileSize,
					Width:   mm.tileSize,
					Height:  mm.tileSize,
					IsSolid: true,
					Type:    "wall",
				}
				objects[obj.Id] = obj
				solidCount++
			}
		}
	}

	log.Printf("Загружено %d твердых объектов (стен) из карты", solidCount)
	mm.gameState.SetObjects(objects)
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
