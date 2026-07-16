package game

import (
	"gamedevRooms/internal/models"
)

const MAX_BULLET_SPEED = 15.0

func NewGameRoom() *models.GameRoom {
	return &models.GameRoom{
		players: make(map[PlayerID]models.Player),
		//InputChan:    make(chan ClientInput, PlayersCount*16),
		//RegisterChan: make(chan PlayerID, PlayersCount),
		//LeaveChan:    make(chan PlayerID, PlayersCount),
	}
}
