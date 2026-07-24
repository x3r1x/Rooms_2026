package factory

import (
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/model"
	"math"

	"github.com/google/uuid"
)

func BulletFactory(player *model.PlayerGameState) model.Bullet {

	barrelLength := domain.PlayerHalfSize + domain.BulletBarrelOffset

	barrelX := player.X + barrelLength*math.Cos(player.Angle)
	barrelY := player.Y + barrelLength*math.Sin(player.Angle)
	sideOffset := domain.BulletBarrelOffset
	sideX := barrelX + math.Cos(player.Angle+math.Pi/2)*sideOffset
	sideY := barrelY + math.Sin(player.Angle+math.Pi/2)*sideOffset
	bullet := model.Bullet{
		Id:        uuid.New().String(),
		X:         sideX,
		Y:         sideY,
		Direction: player.Angle,
		Life:      domain.BulletLife,
		OwnerId:   player.Id,
	}
	//recoilX := -math.Cos(player.Angle) * model.RecoilDistance
	//recoilY := -math.Sin(player.Angle) * model.RecoilDistance
	//player.X += recoilX
	//player.Y += recoilY
	return bullet
}
