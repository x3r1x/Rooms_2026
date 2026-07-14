package logic

import (
	"gamedevRooms/internal/models"
	"math"
)

func UpdateBullets(bullets []models.Bullet) []models.Bullet {
	activeBullets := make([]models.Bullet, 0, len(bullets))
	for _, bullet := range bullets {
		newBullet := bullet
		newBullet.Life--
		newBullet.X += math.Cos(newBullet.Direction) * models.MAX_BULLET_SPEED
		newBullet.Y += math.Sin(newBullet.Direction) * models.MAX_BULLET_SPEED
		// обдумать условие
		if newBullet.Life > 0 && newBullet.Y >= 0 && newBullet.Y <= models.MAP_SIZE && newBullet.X >= 0 && newBullet.X <= models.MAP_SIZE {
			activeBullets = append(activeBullets, newBullet)
		}
	}
	return activeBullets
}
