package domain

import (
	"github.com/google/uuid"
)

type LobbyPlayer struct {
	Nickname string
	Id       string
	Ready    bool
}

func NewLobbyPlayer(nickname string) *LobbyPlayer {
	return &LobbyPlayer{
		Nickname: nickname,
		Id:       uuid.New().String(),
		Ready:    false,
	}
}
