package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Bullet struct {
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	DIRECTION float64 `json:"direction"`
}

type PlayerState struct {
	X         float64         `json:"x"`
	Y         float64         `json:"y"`
	DIRECTION float64         `json:"direction"`
	ID        string          `json:"id"`
	Bullets   []Bullet        `json:"bullets"`
	Conn      *websocket.Conn `json:"-"`
	Mu        sync.Mutex      `json:"-"`
}
