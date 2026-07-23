package factory

import (
	"gamedevRooms/internal/model"
	"math"
)

func BulletFactory(player *model.PlayerGameState) model.Bullet {
	radius := model.PlayerVisualSize / 2.0

	barrelLength := radius + 15.0

	barrelX := barrelLength * math.Cos(player.Angle)
	barrelY := barrelLength * math.Sin(player.Angle)
	bullet := model.Bullet{
		X:         player.X + barrelX,
		Y:         player.Y + barrelY,
		Direction: player.Angle,
		Life:      model.BulletLife,
		OwnerId:   player.Id,
	}
	return bullet
}
