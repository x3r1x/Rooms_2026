package game

import (
	mapModel "gamedevRooms/internal/gameMap"
	"gamedevRooms/internal/model"
	"log"
	"strconv"

	"github.com/google/uuid"
)

type MapManager struct {
	gameMap   *mapModel.Map
	gameState *GameState
	tileSize  float64
}

func NewMapManager(gameState *GameState) *MapManager {
	gameMap := mapModel.NewMap(gameState.GetCountOfPlayers()/3 + 1)
	mm := &MapManager{
		gameMap:   gameMap,
		gameState: gameState,
		tileSize:  model.TileSize, // Используем константу из model
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
			// ИСПРАВЛЕНИЕ: Ищем свойство "blocksPlayer", как указано в tileInfo.json
			for _, prop := range tile.Properties {
				if prop.Name == "blocksPlayer" {
					if val, ok := prop.Value.(bool); ok && val {
						isSolid = true
						break
					}
				}
			}

			if isSolid {
				row := i / model.RoomWidth
				col := i % model.RoomWidth
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

func (mm *MapManager) GetRoomMessages() map[string]model.RoomMessage {
	result := make(map[string]model.RoomMessage)
	for id, room := range mm.gameMap.GetGameMap() {
		barrierNum := 1
		barrierStr := room.GetBarrierType()
		if len(barrierStr) > 7 {
			if num, err := strconv.Atoi(barrierStr[7:]); err == nil {
				barrierNum = num
			}
		}
		result[id] = model.RoomMessage{
			ExitTop:    room.GetExit(model.TopMarker),
			ExitLeft:   room.GetExit(model.LeftMarker),
			ExitBottom: room.GetExit(model.BottomMarker),
			ExitRight:  room.GetExit(model.RightMarker),
			Id:         room.GetId(),
			BorderType: barrierNum,
		}
	}
	return result
}

// GetRoomInfo возвращает комнату по её ID
func (mm *MapManager) GetRoomInfo(roomId string) *mapModel.Room {
	if mm.gameMap == nil {
		return nil
	}
	rooms := mm.gameMap.GetGameMap()
	if rooms == nil {
		return nil
	}
	return rooms[roomId]
}

// GetExit возвращает ID соседней комнаты в указанном направлении
func (mm *MapManager) GetExit(roomId, direction string) string {
	room := mm.GetRoomInfo(roomId)
	if room == nil {
		return ""
	}
	return room.GetExit(direction)
}
