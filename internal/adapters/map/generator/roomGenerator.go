package generator

import (
	"fmt"
	"gamedevRooms/internal/domain"
	"math/rand/v2"
	"strconv"

	"github.com/google/uuid"
)

type RoomGenerator struct {
	baseRoom  []*Tile
	baseWalls []*Tile
	exits     map[string][]*Tile
	flaps     map[string][]*Tile
	barriers  [][]*Tile
}

func NewRoomGenerator(jp *JsonParser) *RoomGenerator {
	tiles := jp.ParseTilesInfo()

	return jp.ParseRoomsInfo(tiles)
}

func (rg *RoomGenerator) CreateRoomWithRandomBarrier() *Room {
	var randomBarrierId = rand.IntN(domain.MaxBarrierType) + domain.MinBarrierType

	return NewRoom(
		RoomExits{},
		uuid.New().String(),
		"barrier"+strconv.Itoa(randomBarrierId),
		rg.createRoomMatrixWithRandomBarrier(randomBarrierId),
	)
}

func (rg *RoomGenerator) ProcessRoomExits(room *Room) {
	rg.processExit(room, room.exits.Top, domain.TopMarker)
	rg.processExit(room, room.exits.Left, domain.LeftMarker)
	rg.processExit(room, room.exits.Bottom, domain.BottomMarker)
	rg.processExit(room, room.exits.Right, domain.RightMarker)
}

func (rg *RoomGenerator) processExit(room *Room, exitType, marker string) {
	if exitType == "" {
		room.SetMatrix(rg.mergeMatrix(rg.flaps[marker], room.GetMatrix()))
	} else {
		room.SetMatrix(rg.mergeMatrix(rg.exits[marker], room.GetMatrix()))
	}
}

func (rg *RoomGenerator) createRoomMatrixWithRandomBarrier(barrierNumber int) []*Tile {
	barrierMatrix := rg.getBarrierByNumber(barrierNumber)

	if len(barrierMatrix) != len(rg.baseRoom) || len(rg.baseRoom) != len(rg.baseWalls) {
		fmt.Println("Не получилось создать комнату со случайным барьером из-за несоответствия"+
			" длин матриц:", len(barrierMatrix), len(rg.baseRoom), len(rg.baseWalls))
		return make([]*Tile, 0)
	}

	baseMatrix := rg.mergeMatrix(rg.baseWalls, rg.baseRoom)
	return rg.mergeMatrix(barrierMatrix, baseMatrix)
}

func (rg *RoomGenerator) mergeMatrix(matrix1, matrix2 []*Tile) []*Tile {
	var resultMatrix = make([]*Tile, 0)

	for i := 0; i < len(rg.baseRoom); i++ {
		if matrix1[i] != nil {
			resultMatrix = append(resultMatrix, matrix1[i])
		} else {
			resultMatrix = append(resultMatrix, matrix2[i])
		}
	}

	return resultMatrix
}

func (rg *RoomGenerator) getBarrierByNumber(barrierNumber int) []*Tile {
	if domain.BarriersAmount-barrierNumber < 0 {
		fmt.Println("Попытка взять неправильный индекс барьера:", barrierNumber)
		return make([]*Tile, 0)
	}

	return rg.barriers[domain.BarriersAmount-barrierNumber]
}
