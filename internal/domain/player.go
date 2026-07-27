package domain

import (
	"math"
	"math/rand/v2"
)

type WeaponFunc func(player *PlayerGameState) []Bullet

type PlayerGameState struct {
	Id            string     `json:"id"`
	Nickname      string     `json:"-"`
	Health        float64    `json:"h"`
	X             float64    `json:"x"`
	Y             float64    `json:"y"`
	Angle         float64    `json:"a"`
	MoveX         float64    `json:"mx"`
	MoveY         float64    `json:"my"`
	CooldownTimer int        `json:"-"`
	RebornTimer   int        `json:"rt"`
	BodyCount     int        `json:"-"`
	DeathCount    int        `json:"-"`
	RoomId        string     `json:"room_id"`
	PlayerClass   string     `json:"pc"`
	Weapon        WeaponFunc `json:"-"`
}

type PlayerStatistic struct {
	Id     string  `json:"id"`
	Hp     float64 `json:"h"`
	Kills  int     `json:"k"`
	Deaths int     `json:"d"`
}

type PlayerFinalState struct {
	Id     string `json:"id"`
	Kills  int    `json:"k"`
	Deaths int    `json:"d"`
}

func NewPlayerGameState(id, nickname, playerClass string) *PlayerGameState {
	p := &PlayerGameState{
		Id:            id,
		Nickname:      nickname,
		Health:        MaxPlayerHealth,
		X:             PlayerSpawnPointX,
		Y:             PlayerSpawnPointY,
		Angle:         InitDirection,
		MoveX:         InitValue,
		MoveY:         InitValue,
		CooldownTimer: InitValue,
		RebornTimer:   InitValue,
		BodyCount:     InitValue,
		DeathCount:    InitValue,
		PlayerClass:   playerClass,
	}

	p.setupWeaponByClass()
	return p
}

func (p *PlayerGameState) setupWeaponByClass() {
	switch p.PlayerClass {
	case PlayerGun:
		p.Weapon = createGunBullets
	case PlayerRifle:
		p.Weapon = createRifleBullets
	case PlayerSom:
		p.Weapon = createSomBullet
	}
}

func createGunBullets(player *PlayerGameState) []Bullet {
	barrelLength := PlayerHalfSize + BulletBarrelOffset
	barrelX := player.X + barrelLength*math.Cos(player.Angle)
	barrelY := player.Y + barrelLength*math.Sin(player.Angle)

	sideOffset := BulletBarrelOffset
	sideX := barrelX + math.Cos(player.Angle+math.Pi/2)*sideOffset
	sideY := barrelY + math.Sin(player.Angle+math.Pi/2)*sideOffset

	return []Bullet{{
		Id:        GenerateID(),
		X:         sideX,
		Y:         sideY,
		Direction: player.Angle,
		Life:      BulletLifeGun,
		Damage:    BulletDamageGun,
		Speed:     BulletSpeedGun,
		OwnerId:   player.Id,
		Type:      PlayerGun,
	}}
}

func createRifleBullets(player *PlayerGameState) []Bullet {
	pellets := make([]Bullet, 0, PelletCount)
	barrelLength := PlayerHalfSize + BulletBarrelOffset

	for i := 0; i < PelletCount; i++ {
		offset := (rand.Float64() - 0.5) * SpreadAngle
		direction := player.Angle + offset
		barrelX := player.X + barrelLength*math.Cos(direction)
		barrelY := player.Y + barrelLength*math.Sin(direction)

		pellets = append(pellets, Bullet{
			Id:        GenerateID(),
			X:         barrelX,
			Y:         barrelY,
			Direction: direction,
			Life:      BulletLifeRifle,
			Damage:    BulletDamageRifle,
			Speed:     BulletSpeedRifle,
			OwnerId:   player.Id,
			Type:      PlayerRifle,
		})
	}

	return pellets
}

func createSomBullet(player *PlayerGameState) []Bullet {
	barrelLength := PlayerHalfSize + BulletBarrelOffset
	barrelX := player.X + barrelLength*math.Cos(player.Angle)
	barrelY := player.Y + barrelLength*math.Sin(player.Angle)

	return []Bullet{{
		Id:        GenerateID(),
		X:         barrelX,
		Y:         barrelY,
		Direction: player.Angle,
		Life:      math.Round(BulletLifeSom * rand.Float64()),
		Damage:    BulletDamageSom,
		Speed:     BulletSpeedSom,
		OwnerId:   player.Id,
		Type:      PlayerSom,
	}}
}
