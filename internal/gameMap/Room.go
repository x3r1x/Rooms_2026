package mapModel

import (
	"fmt"
	"gamedevRooms/internal/model"
)

type RoomExits struct {
	Top    string
	Right  string
	Bottom string
	Left   string
}

type Room struct {
	exits       RoomExits
	id          string
	barrierType string
	matrix      []*Tile
}

func NewRoom(exits RoomExits, id, barrierType string, matrix []*Tile) *Room {
	return &Room{
		exits:       exits,
		id:          id,
		barrierType: barrierType,
		matrix:      matrix,
	}
}

func (r *Room) SetExit(direction, roomId string) {
	switch direction {
	case model.TopMarker:
		r.exits.Top = roomId
	case model.LeftMarker:
		r.exits.Left = roomId
	case model.RightMarker:
		r.exits.Right = roomId
	case model.BottomMarker:
		r.exits.Bottom = roomId
	}
}

func (r *Room) SetMatrix(newMatrix []*Tile) {
	r.matrix = newMatrix
}

func (r *Room) GetId() string {
	return r.id
}

func (r *Room) GetExit(direction string) string {
	switch direction {
	case model.TopMarker:
		return r.exits.Top
	case model.LeftMarker:
		return r.exits.Left
	case model.RightMarker:
		return r.exits.Right
	case model.BottomMarker:
		return r.exits.Bottom
	}

	return ""
}

func (r *Room) GetBarrierType() string {
	return r.barrierType
}

func (r *Room) GetMatrix() []*Tile {
	return r.matrix
}

func (r *Room) PrintMatrix() {
	for i := 0; i < model.RoomHeight; i++ {
		for j := 0; j < model.RoomWidth; j++ {
			if r.matrix[i*model.RoomWidth+j] == nil {
				fmt.Print("- ")
			} else {
				fmt.Print(r.matrix[i*model.RoomWidth+j].Id, " ")
			}
		}

		fmt.Println()
	}
}
