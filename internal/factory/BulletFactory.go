package factory

import (
	"gamedevRooms/internal/model"
	"math"
)

func BulletFactory(player *model.PlayerGameState) model.Bullet {
	localX := model.PlayerVisualSize / 2.0
	localY := (model.PlayerVisualSize / 2.0) - (model.BulletWidth + (model.PlayerVisualSize * 0.1))

	rotatedDX := localX*math.Cos(player.Angle) - localY*math.Sin(player.Angle)
	rotatedDY := localX*math.Sin(player.Angle) + localY*math.Cos(player.Angle)
	bullet := model.Bullet{
		X:         player.X + rotatedDX,
		Y:         player.Y + rotatedDY,
		Direction: player.Angle,
		Life:      model.BulletLife,
		OwnerId:   player.Id,
	}
	return bullet
}
