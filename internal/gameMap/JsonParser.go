package mapModel

import (
	"encoding/json"
	"fmt"
	"gamedevRooms/internal/model"
	"os"
)

type JsonParser struct{}

func NewJsonParser() *JsonParser {
	return &JsonParser{}
}

func (jp *JsonParser) ParseTilesInfo() map[int16]*Tile {
	tilesJson, err := os.Open("gameMap/tileInfo.json")

	if err != nil {
		fmt.Println("Error opening tiles JSON:", err)
		return map[int16]*Tile{}
	}
	defer func(tilesJson *os.File) {
		err := tilesJson.Close()

		if err != nil {
			fmt.Println("Error closing tiles JSON:", err)
		}
	}(tilesJson)

	return jp.decodeTilesJson(tilesJson)
}

func (jp *JsonParser) ParseRoomsInfo(tiles map[int16]*Tile) *RoomGenerator {
	roomsJson, err := os.Open("gameMap/allRoom.json")

	if err != nil {
		fmt.Println("Error opening rooms JSON:", err)
		return &RoomGenerator{
			baseRoom:  nil,
			baseWalls: nil,
			exits:     nil,
			flaps:     nil,
			barriers:  nil,
		}
	}
	defer func(roomsJson *os.File) {
		err := roomsJson.Close()

		if err != nil {
			fmt.Println("Error closing tiles JSON:", err)
		}
	}(roomsJson)

	return jp.decodeRoomsJson(roomsJson, tiles)
}

func (jp *JsonParser) decodeTilesJson(_json *os.File) map[int16]*Tile {
	decoder := json.NewDecoder(_json)

	for {
		token, err := decoder.Token()

		if err != nil {
			fmt.Println("Ошибка при поиске поля tiles:", err)
			break
		}

		if key, ok := token.(string); ok && key == "tiles" {
			return jp.getTilesMap(decoder)
		}
	}
	return map[int16]*Tile{}
}

func (jp *JsonParser) getTilesMap(decoder *json.Decoder) map[int16]*Tile {
	var tilesSlice []Tile

	if err := decoder.Decode(&tilesSlice); err != nil {
		fmt.Println("Ошибка декодирования tiles:", err)
		return map[int16]*Tile{}
	}

	var tilesMap = make(map[int16]*Tile)

	for i := range tilesSlice {
		tilesSlice[i].Id++

		tilesMap[tilesSlice[i].Id] = &tilesSlice[i]
	}

	return tilesMap
}

func (jp *JsonParser) decodeRoomsJson(_json *os.File, tiles map[int16]*Tile) *RoomGenerator {
	decoder := json.NewDecoder(_json)

	for {
		token, err := decoder.Token()

		if err != nil {
			fmt.Println("Ошибка при поиске поля layers:", err)
			break
		}

		if key, ok := token.(string); ok && key == "layers" {
			return jp.getLayersData(decoder, tiles)
		}
	}

	return &RoomGenerator{
		baseRoom:  nil,
		baseWalls: nil,
		exits:     nil,
		flaps:     nil,
		barriers:  nil,
	}
}

type layerDataOnly struct {
	Data []int16 `json:"data"`
}

func (jp *JsonParser) getLayersData(decoder *json.Decoder, tiles map[int16]*Tile) *RoomGenerator {
	bracketToken, err := decoder.Token()

	if err != nil {
		fmt.Println("Ошибка декодирования layers:", err)
		return &RoomGenerator{
			baseRoom:  make([]*Tile, 0),
			baseWalls: make([]*Tile, 0),
			exits:     make(map[string][]*Tile),
			flaps:     make(map[string][]*Tile),
			barriers:  make([][]*Tile, 0),
		}
	}

	if delim, ok := bracketToken.(json.Delim); !ok || delim != '[' {
		fmt.Println("Ожидалось начало массива '['")
		return &RoomGenerator{
			baseRoom:  make([]*Tile, 0),
			baseWalls: make([]*Tile, 0),
			exits:     make(map[string][]*Tile),
			flaps:     make(map[string][]*Tile),
			barriers:  make([][]*Tile, 0),
		}
	}

	var dataArray = make([][]int16, 0)

	for decoder.More() {
		var thisData layerDataOnly

		if err := decoder.Decode(&thisData); err != nil {
			fmt.Println("Ошибка поиска поля data:", err)
			return &RoomGenerator{
				baseRoom:  make([]*Tile, 0),
				baseWalls: make([]*Tile, 0),
				exits:     make(map[string][]*Tile),
				flaps:     make(map[string][]*Tile),
				barriers:  make([][]*Tile, 0),
			}
		}

		dataArray = append(dataArray, thisData.Data)
	}

	return jp.getRoomGenerator(dataArray, tiles)
}

func (jp *JsonParser) getRoomGenerator(data [][]int16, tiles map[int16]*Tile) *RoomGenerator {
	var roomGenerator = RoomGenerator{
		baseRoom:  jp.attachTilesToData(data[model.BaseRoomIndex], tiles),
		baseWalls: jp.attachTilesToData(data[model.BaseWallsIndex], tiles),
		exits:     jp.getRoomGeneratorExits(data, tiles),
		flaps:     jp.getRoomGeneratorFlaps(data, tiles),
		barriers:  jp.getRoomGeneratorBarriers(data, tiles),
	}

	return &roomGenerator
}

func (jp *JsonParser) getRoomGeneratorExits(data [][]int16, tiles map[int16]*Tile) map[string][]*Tile {
	var exits = make(map[string][]*Tile)

	exits[model.TopMarker] = jp.attachTilesToData(data[model.ExitTopIndex], tiles)
	exits[model.LeftMarker] = jp.attachTilesToData(data[model.ExitLeftIndex], tiles)
	exits[model.BottomMarker] = jp.attachTilesToData(data[model.ExitBottomIndex], tiles)
	exits[model.RightMarker] = jp.attachTilesToData(data[model.ExitRightIndex], tiles)

	return exits
}

func (jp *JsonParser) getRoomGeneratorFlaps(data [][]int16, tiles map[int16]*Tile) map[string][]*Tile {
	var flaps = make(map[string][]*Tile)

	flaps[model.TopMarker] = jp.attachTilesToData(data[model.FlapTopIndex], tiles)
	flaps[model.LeftMarker] = jp.attachTilesToData(data[model.FlapLeftIndex], tiles)
	flaps[model.BottomMarker] = jp.attachTilesToData(data[model.FlapBottomIndex], tiles)
	flaps[model.RightMarker] = jp.attachTilesToData(data[model.FlapRightIndex], tiles)

	return flaps
}

func (jp *JsonParser) getRoomGeneratorBarriers(data [][]int16, tiles map[int16]*Tile) [][]*Tile {
	var barriers = make([][]*Tile, 0)

	for i := 1; i <= 7; i++ {
		barriers = append(barriers, jp.attachTilesToData(data[model.BarriersStartIndex-1+i], tiles))
	}

	return barriers
}

func (jp *JsonParser) attachTilesToData(data []int16, tiles map[int16]*Tile) []*Tile {
	var tilesArray = make([]*Tile, model.RoomHeight*model.RoomWidth)

	for id, tileType := range data {
		tilesArray[id] = tiles[tileType]
	}

	return tilesArray
}
