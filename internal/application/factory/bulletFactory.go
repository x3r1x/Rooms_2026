package factory

import (
	"gamedevRooms/internal/domain"
	"math"
)

type BulletFactory struct {
	BulletDamage int
	BulletLife   int
}

func NewBulletFactory() *BulletFactory {
	return &BulletFactory{}
}

func (bf *BulletFactory) CreateBullet(player *domain.PlayerGameState) domain.Bullet {

	barrelLength := domain.PlayerHalfSize + domain.BulletBarrelOffset

	barrelX := player.X + barrelLength*math.Cos(player.Angle)
	barrelY := player.Y + barrelLength*math.Sin(player.Angle)
	sideOffset := domain.BulletBarrelOffset
	sideX := barrelX + math.Cos(player.Angle+math.Pi/2)*sideOffset
	sideY := barrelY + math.Sin(player.Angle+math.Pi/2)*sideOffset
	bullet := domain.Bullet{
		Id:        domain.GenerateID(),
		X:         sideX,
		Y:         sideY,
		Direction: player.Angle,
		Life:      player.BulletLife,
		Damage:    player.BulletDamage,
		Speed:     player.BulletSpeed,
		OwnerId:   player.Id,
	}
	//recoilX := -math.Cos(player.Angle) * model.RecoilDistance
	//recoilY := -math.Sin(player.Angle) * model.RecoilDistance
	//player.X += recoilX
	//player.Y += recoilY
	return bullet
}
