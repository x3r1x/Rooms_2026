package _map

import (
	"encoding/json"
	"gamedevRooms/internal/adapters/map/generator"
	"gamedevRooms/internal/domain"
	"log"
	"strconv"

	"github.com/google/uuid"
)

type MapManager struct {
	gameMap  *generator.Map
	tileSize float64
}

func NewMapManager() *MapManager {
	return &MapManager{
		gameMap:  nil,
		tileSize: domain.TileSize,
	}
}

type HitboxRect struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}

func (mm *MapManager) LoadMapObjects(playerCount int) map[string]*domain.Object {
	mm.gameMap = generator.NewMap(playerCount/2 + 1)
	objects := make(map[string]*domain.Object)

	for roomId, room := range mm.gameMap.GetGameMap() {
		matrix := room.GetMatrix()
		for i, tile := range matrix {
			if tile == nil {
				continue
			}

			blocksPlayer := false
			hitboxStr := ""
			for _, prop := range tile.Properties {

				switch prop.Name {
				case "blocksPlayer":
					if val, ok := prop.Value.(bool); ok && val {
						blocksPlayer = true
					}
				case "hitbox":
					if val, ok := prop.Value.(string); ok {
						hitboxStr = val
					}
				}
			}
			if !blocksPlayer {
				continue
			}

			row := i / domain.RoomWidth
			col := i % domain.RoomWidth

			baseX := float64(col) * mm.tileSize
			baseY := float64(row) * mm.tileSize

			var objs []*domain.Object
			if hitboxStr != "" {
				hitboxes := mm.parseHitbox(hitboxStr)
				if len(hitboxes) > 0 {
					objs = mm.createObjectsFromHitboxes(baseX, baseY, hitboxes, roomId)
				}
			}

			if len(objs) == 0 {
				objs = []*domain.Object{mm.createFullTileObject(baseX, baseY, roomId)}
			}

			for _, obj := range objs {
				objects[obj.Id] = obj
			}
		}
	}
	return objects
}

func (mm *MapManager) parseHitbox(hitboxStr string) []HitboxRect {
	if hitboxStr == "" {
		return nil
	}
	var hitboxes []HitboxRect
	log.Println(hitboxStr)
	err := json.Unmarshal([]byte(hitboxStr), &hitboxes)
	if err != nil {
		log.Printf("Ошибка парсинга hitbox: %v", err)
		return nil
	}
	return hitboxes
}

func (mm *MapManager) createObjectsFromHitboxes(baseX, baseY float64, hitboxes []HitboxRect, roomId string) []*domain.Object {
	var objects []*domain.Object
	for _, hb := range hitboxes {
		if hb.W <= 0 || hb.H <= 0 {
			continue
		}
		if hb.W == mm.tileSize && hb.H == mm.tileSize && hb.X == 0 && hb.Y == 0 {
			continue
		}
		obj := &domain.Object{
			Id:      uuid.New().String(),
			X:       baseX + hb.X,
			Y:       baseY + hb.Y,
			Width:   hb.W,
			Height:  hb.H,
			IsSolid: true,
			Type:    "wall",
			RoomId:  roomId,
		}
		objects = append(objects, obj)
	}
	if len(objects) == 0 {
		objects = append(objects, mm.createFullTileObject(baseX, baseY, roomId))
	}
	return objects
}

func (mm *MapManager) createFullTileObject(baseX, baseY float64, roomId string) *domain.Object {
	return &domain.Object{
		Id:      uuid.New().String(),
		X:       baseX,
		Y:       baseY,
		Width:   mm.tileSize,
		Height:  mm.tileSize,
		IsSolid: true,
		Type:    "wall",
		RoomId:  roomId,
	}
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
