package models

import "github.com/gorilla/websocket"

type PlayerState struct {
	Id         string          `json:"id"`
	X          float64         `json:"x"`
	Y          float64         `json:"y"`
	Direction  float64         `json:"direction"`
	Bullets    []Bullet        `json:"bullets"`
	Connection *websocket.Conn `json:"-"`
}
