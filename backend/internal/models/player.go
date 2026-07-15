package models

import (
	"github.com/gorilla/websocket"
)

type PlayerState struct {
	X         float64           `json:"x"`
	Y         float64           `json:"y"`
	Direction float64           `json:"direction"`
	Id        string            `json:"id"`
	Bullets   map[string]Bullet `json:"bullets"`
	Conn      *websocket.Conn   `json:"-"`
	SendChan  chan<- []byte     `json:"-"`
}

type PlayerSnapshot struct {
	Id        string   `json:"id"`
	X         float64  `json:"x"`
	Y         float64  `json:"y"`
	Direction float64  `json:"movementDirection"`
	Bullets   []Bullet `json:"bullets"`
}
