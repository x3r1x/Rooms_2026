package domain

import (
	"math"
	"math/rand/v2"
)

type WeaponFunc func(player *PlayerGameState) []Bullet

type PlayerGameState struct {
	Id                string     `json:"id"`
	Nickname          string     `json:"-"`
	Health            float64    `json:"h"`
	X                 float64    `json:"x"`
	Y                 float64    `json:"y"`
	Angle             float64    `json:"a"`
	Speed             float64    `json:"-"`
	MoveX             float64    `json:"mx"`
	MoveY             float64    `json:"my"`
	CooldownTimer     int        `json:"-"`
	RebornTimer       int        `json:"rt"`
	PlayerShield      bool       `json:"ps"`
	PlayerShieldTimer int        `json:"-"`
	BodyCount         int        `json:"-"`
	DeathCount        int        `json:"-"`
	RoomId            string     `json:"room_id"`
	PlayerClass       string     `json:"pc"`
	Weapon            WeaponFunc `json:"-"`
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
		Id:                id,
		Nickname:          nickname,
		Health:            MaxPlayerHealth,
		X:                 InitValue,
		Y:                 InitValue,
		Angle:             InitDirection,
		MoveX:             InitValue,
		MoveY:             InitValue,
		CooldownTimer:     InitValue,
		PlayerShield:      false,
		PlayerShieldTimer: InitValue,
		RebornTimer:       InitValue,
		BodyCount:         InitValue,
		DeathCount:        InitValue,
		PlayerClass:       playerClass,
	}
	p.setupSpeedByClass()
	p.setupWeaponByClass()
	return p
}

func (p *PlayerGameState) setupSpeedByClass() {
	switch p.PlayerClass {
	case PlayerGun:
		p.Speed = PlayerSpeed
	case PlayerRifle:
		p.Speed = PlayerSpeed
	case PlayerSom:
		p.Speed = PlayerSpeed
	}
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

func (p *PlayerGameState) SetCooldown() {
	switch p.PlayerClass {
	case PlayerGun:
		p.CooldownTimer = ShootCooldown - 1
	case PlayerRifle:
		p.CooldownTimer = ShootCooldown*2 - 1
	case PlayerSom:
		p.CooldownTimer = ShootCooldown*10 - 1
	}
}

func createGunBullets(player *PlayerGameState) []Bullet {
	rotatedX := GunOffsetX*math.Cos(player.Angle) - GunOffsetY*math.Sin(player.Angle)
	rotatedY := GunOffsetX*math.Sin(player.Angle) + GunOffsetY*math.Cos(player.Angle)

	startX := player.X + rotatedX
	startY := player.Y + rotatedY

	return []Bullet{{
		Id:        GenerateID(),
		X:         startX,
		Y:         startY,
		Direction: player.Angle,
		Life:      BulletLifeGun,
		Damage:    BulletDamageGun,
		Speed:     BulletSpeedGun,
		OwnerId:   player.Id,
		Type:      PlayerGun,
		RoomId:    player.RoomId,
	}}
}

func createRifleBullets(player *PlayerGameState) []Bullet {
	pellets := make([]Bullet, 0, PelletCount)
	rotatedX := RifleOffsetX*math.Cos(player.Angle) - RifleOffsetY*math.Sin(player.Angle)
	rotatedY := RifleOffsetX*math.Sin(player.Angle) + RifleOffsetY*math.Cos(player.Angle)

	for i := 0; i < PelletCount; i++ {
		offset := (rand.Float64() - 0.5) * SpreadAngle
		direction := player.Angle + offset
		startX := player.X + rotatedX
		startY := player.Y + rotatedY

		pellets = append(pellets, Bullet{
			Id:        GenerateID(),
			X:         startX,
			Y:         startY,
			Direction: direction,
			Life:      BulletLifeRifle,
			Damage:    BulletDamageRifle,
			Speed:     BulletSpeedRifle,
			OwnerId:   player.Id,
			Type:      PlayerRifle,
			RoomId:    player.RoomId,
		})
	}

	return pellets
}

func createSomBullet(player *PlayerGameState) []Bullet {
	rotatedX := SomOffsetX*math.Cos(player.Angle) - SomOffsetY*math.Sin(player.Angle)
	rotatedY := SomOffsetX*math.Sin(player.Angle) + SomOffsetY*math.Cos(player.Angle)

	startX := player.X + rotatedX
	startY := player.Y + rotatedY

	return []Bullet{{
		Id:        GenerateID(),
		X:         startX,
		Y:         startY,
		Direction: player.Angle,
		Life:      math.Round(BulletLifeSom * rand.Float64()),
		Damage:    BulletDamageSom,
		Speed:     BulletSpeedSom,
		OwnerId:   player.Id,
		Type:      PlayerSom,
		RoomId:    player.RoomId,
	}}
}
