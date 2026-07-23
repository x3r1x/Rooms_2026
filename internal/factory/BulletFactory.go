package factory

import (
	"gamedevRooms/internal/model"
	"math"

	"github.com/google/uuid"
)

func BulletFactory(player *model.PlayerGameState) model.Bullet {

	barrelLength := model.PlayerHalfSize + model.BulletBarrelOffset

	barrelX := player.X + barrelLength*math.Cos(player.Angle)
	barrelY := player.Y + barrelLength*math.Sin(player.Angle)
	sideOffset := 24.0
	sideX := barrelX + math.Cos(player.Angle+math.Pi/2)*sideOffset
	sideY := barrelY + math.Sin(player.Angle+math.Pi/2)*sideOffset
	bullet := model.Bullet{
		Id:        uuid.New().String(),
		X:         sideX,
		Y:         sideY,
		Direction: player.Angle,
		Life:      model.BulletLife,
		OwnerId:   player.Id,
	}
	return bullet
}
