package domain

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type LobbyPlayer struct {
	Nickname   string
	Id         string
	Ready      bool
	Connection *websocket.Conn
}

func NewLobbyPlayer(nickname string) *LobbyPlayer {
	return &LobbyPlayer{
		Nickname: nickname,
		Id:       uuid.New().String(),
		Ready:    false,
	}
}
