package models

import "github.com/gorilla/websocket"

type PlayerState struct {
	Id            string          `json:"id"`
	X             float64         `json:"x"`
	Y             float64         `json:"y"`
	DirectionLook float64         `json:"directionlook"`
	Connection    *websocket.Conn `json:"-"`
}
