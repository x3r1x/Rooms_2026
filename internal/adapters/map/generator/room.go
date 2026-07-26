package generator

import (
	"fmt"
	"gamedevRooms/internal/domain"
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
	case domain.TopMarker:
		r.exits.Top = roomId
	case domain.LeftMarker:
		r.exits.Left = roomId
	case domain.RightMarker:
		r.exits.Right = roomId
	case domain.BottomMarker:
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
	case domain.TopMarker:
		return r.exits.Top
	case domain.LeftMarker:
		return r.exits.Left
	case domain.RightMarker:
		return r.exits.Right
	case domain.BottomMarker:
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
	for i := 0; i < domain.RoomHeight; i++ {
		for j := 0; j < domain.RoomWidth; j++ {
			if r.matrix[i*domain.RoomWidth+j] == nil {
				fmt.Print("- ")
			} else {
				fmt.Print(r.matrix[i*domain.RoomWidth+j].Id, " ")
			}
		}

		fmt.Println()
	}
}
