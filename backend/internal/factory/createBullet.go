package factory

import (
	"gamedevRooms/internal/models"
	"math"
)

func CreateBullet(player *models.PlayerState, direction float64) models.Bullet {
	const localX = models.PLAYER_VISUAL_SIZE / 2
	const localY = models.PLAYER_VISUAL_SIZE/2 - (models.BULLET_WIDTH + (models.PLAYER_VISUAL_SIZE * 0.1))

	var rotatedX = localX*math.Cos(direction) - localY*math.Sin(direction)
	var rotatedY = localX*math.Sin(direction) + localY*math.Cos(direction)

	return models.Bullet{
		X:         player.X + rotatedX,
		Y:         player.Y + rotatedY,
		Direction: direction,
		Life:      models.BULLET_LIFE,
	}
}
