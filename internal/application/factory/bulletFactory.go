package factory

import (
	"gamedevRooms/internal/domain"
	"math"
	"math/rand/v2"
)

type BulletFactory struct {
	BulletDamage int
	BulletLife   int
}

func NewBulletFactory() *BulletFactory {
	return &BulletFactory{}
}

func (bf *BulletFactory) CreateBullet(player *domain.PlayerGameState) []domain.Bullet {
	var bullets []domain.Bullet
	switch player.PlayerClass {
	case domain.PlayerRifle:
		bullets = bf.createShotgunBullets(player)
	default:
		bullets = append(bullets, bf.createSingleBullet(player))
	}
	return bullets
}

func (bf *BulletFactory) createSingleBullet(player *domain.PlayerGameState) domain.Bullet {
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
		Type:      player.PlayerClass,
	}
	//recoilX := -math.Cos(player.Angle) * model.RecoilDistance
	//recoilY := -math.Sin(player.Angle) * model.RecoilDistance
	//player.X += recoilX
	//player.Y += recoilY
	return bullet
}

func (bf *BulletFactory) createShotgunBullets(player *domain.PlayerGameState) []domain.Bullet {

	pellets := make([]domain.Bullet, 0, domain.PelletCount)
	barrelLength := domain.PlayerHalfSize + domain.BulletBarrelOffset

	for i := 0; i < domain.PelletCount; i++ {
		offset := (rand.Float64() - 0.5) * domain.SpreadAngle
		direction := player.Angle + offset

		barrelX := player.X + barrelLength*math.Cos(direction)
		barrelY := player.Y + barrelLength*math.Sin(direction)

		pellets = append(pellets, domain.Bullet{
			Id:        domain.GenerateID(),
			X:         barrelX,
			Y:         barrelY,
			Direction: direction,
			Life:      player.BulletLife,
			OwnerId:   player.Id,
			Speed:     player.BulletSpeed,
			Damage:    domain.BulletDamageRifle,
			Type:      player.PlayerClass,
		})
	}

	return pellets
}
