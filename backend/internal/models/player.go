package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type PlayerState struct {
	X         float64         `json:"x"`
	Y         float64         `json:"y"`
	Direction float64         `json:"direction"`
	Id        string          `json:"id"`
	Bullets   []Bullet        `json:"bullets"`
	Conn      *websocket.Conn `json:"-"`
	Mu        sync.Mutex      `json:"-"`
}
