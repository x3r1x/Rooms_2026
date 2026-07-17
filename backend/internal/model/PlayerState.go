package model

import "github.com/gorilla/websocket"

type PlayerState struct {
	Id          string          `json:"id"`
	Health      float64         `json:"-"`
	X           float64         `json:"x"`
	Y           float64         `json:"y"`
	Angle       float64         `json:"a"`
	MoveX       float64         `json:"mx"`
	MoveY       float64         `json:"my"`
	Connection  *websocket.Conn `json:"-"`
	ShootTimer  int             `json:"-"`
	RebornTimer int             `json:"-"`
}
