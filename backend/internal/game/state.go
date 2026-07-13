package game

import (
	"gamedevRooms/internal/models"
	"sync"
)

var (
	Clients = make(map[string]*models.PlayerState)
	Mutex   sync.Mutex
)

const MAX_BULLET_SPEED = 1.5
const MAP_SIZE = 2000.0
const BULLET_LIFE = 60

func UpdatePlayerState(id string, x, y, dir float64, clientBullets []models.Bullet) {

	state, ok := Clients[id]
	if !ok {
		return
	}
	state.X = x
	state.Y = y
	state.Direction = dir

	for _, cb := range clientBullets {
		isDuplicate := false
		for _, existing := range state.Bullets {
			if existing.Id == cb.Id {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			newBullet := models.Bullet{
				Id:        cb.Id,
				X:         x,
				Y:         y,
				Direction: cb.Direction,
				StartX:    state.X,
				StartY:    state.Y,
				Owner:     cb.Owner,
				Life:      BULLET_LIFE,
			}
			state.Bullets = append(state.Bullets, newBullet)
		}
	}
}
