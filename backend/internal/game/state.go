package game

import (
	"gamedevRooms/internal/models"
	"sync"
)

var (
	Clients = make(map[string]*models.PlayerState)
	Mutex   sync.Mutex
)

const MAX_BULLET_SPEED = 22.5
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

	existingID := make(map[string]struct{}, len(state.Bullets))
	for _, b := range state.Bullets {
		existingID[b.Id] = struct{}{}
	}
	for _, cb := range clientBullets {
		if _, exist := existingID[cb.Id]; !exist {
			newBullet := models.Bullet{
				Id:        cb.Id,
				X:         x,
				Y:         y,
				Direction: cb.Direction,
				StartX:    x,
				StartY:    y,
				Owner:     cb.Owner,
				Life:      BULLET_LIFE,
			}
			state.Bullets = append(state.Bullets, newBullet)
			existingID[cb.Id] = struct{}{}
		}
	}
}
