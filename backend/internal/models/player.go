package models

import (
	"github.com/gorilla/websocket"
)

type PlayerState struct {
	X         float64           `json:"x"`
	Y         float64           `json:"y"`
	Direction float64           `json:"direction"`
	Bullets   map[string]Bullet `json:"bullets"`
	Conn      *websocket.Conn   `json:"-"`
	SendChan  chan<- []byte     `json:"-"`
}

type PlayerInterpolation struct {
	DX             float64                        `json:"dx"`
	DY             float64                        `json:"dy"`
	DeltaDirection float64                        `json:"deltaDirection"`
	DeltaBullets   map[string]BulletInterpolation `json:"deltaBullets"`
	NewBullets     map[string]NewBullet           `json:"newBullets"`
	Conn           *websocket.Conn                `json:"-"`
	SendChan       chan<- []byte                  `json:"-"`
}
