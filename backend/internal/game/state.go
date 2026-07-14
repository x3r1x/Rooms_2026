package game

import (
	"gamedevRooms/internal/factory"
	"gamedevRooms/internal/models"
	"math"
	"sync"

	"github.com/google/uuid"
)

var (
	Clients         = make(map[string]*models.PlayerState)
	PreviousClients = make(map[string]*models.PlayerState)
	Mutex           sync.RWMutex
)

type PlayerMovementDirection struct {
	X float64
	Y float64
}

func UpdatePlayerState(id string, movementDirection PlayerMovementDirection, deltaVisualDirection float64, newBulletDirections []float64) {

	state, ok := Clients[id]
	if !ok {
		return
	}
	normalizeMovementDirection(&movementDirection)
	state.X += movementDirection.X * models.PLAYER_SPEED * models.TICK_TIME
	state.Y += movementDirection.Y * models.PLAYER_SPEED * models.TICK_TIME

	state.Direction += deltaVisualDirection

	for _, dir := range newBulletDirections {
		state.Bullets[uuid.New().String()] = factory.CreateBullet(state, dir)
	}
}

func normalizeMovementDirection(movementDirection *PlayerMovementDirection) {
	var vectorLength = math.Sqrt(movementDirection.X*movementDirection.X + movementDirection.Y*movementDirection.Y)

	if vectorLength != 0 {
		*movementDirection = PlayerMovementDirection{
			X: movementDirection.X / vectorLength,
			Y: movementDirection.Y / vectorLength,
		}
	}
}
