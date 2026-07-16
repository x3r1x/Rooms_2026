package models

import "github.com/gorilla/websocket"

type PlayerState struct {
	Id         string          `json:"id"`
	X          float64         `json:"x"`
	Y          float64         `json:"y"`
	A          float64         `json:"a"`
	MoveX      float64         `json:"mx"`
	MoveY      float64         `json:"my"`
	Connection *websocket.Conn `json:"-"`
	ShootTimer int             `json:"shotTimer"`
}
