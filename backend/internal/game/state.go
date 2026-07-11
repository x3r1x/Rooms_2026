package game

import (
	"gamedevRooms/internal/models"
	"sync"
)

var (
	Clients = make(map[string]*models.PlayerState)
	Mutex   sync.Mutex
)

const MAX_BULLET_SPEED = 15.0
