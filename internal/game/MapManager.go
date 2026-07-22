package game

import (
	mapModel "gamedevRooms/internal/gameMap"
	"gamedevRooms/internal/model"

	"github.com/google/uuid"
)

type MapManager struct {
	gameMap   *mapModel.Map
	gameState *GameState
	tileSize  float64
}

func NewMapManager(gameState *GameState) *MapManager {
	gameMap := mapModel.NewMap(7)

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
