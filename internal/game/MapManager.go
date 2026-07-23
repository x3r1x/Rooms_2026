package game

import (
	mapModel "gamedevRooms/internal/gameMap"
	"gamedevRooms/internal/model"
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
		tileSize:  36,
	}

	mm.loadMapObjects()

	return mm
}

func (mm *MapManager) loadMapObjects() {
	objects := make(map[string]*model.Object)

	for _, room := range mm.gameMap.GetGameMap() {
		matrix := room.GetMatrix()
		for i, tile := range matrix {
			if tile == nil {
				continue
			}

			isSolid := false
			for _, prop := range tile.Properties {
				if prop.Name == "solid" && prop.Value == true {
					isSolid = true
					break
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
					//IsDestroyable: false,
					//Health:        0,
					Type: "wall",
				}
				objects[obj.Id] = obj
			}
		}
	}
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
