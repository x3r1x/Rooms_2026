package domain

type PlayerGameState struct {
	Id           string  `json:"id"`
	Nickname     string  `json:"-"`
	Health       float64 `json:"h"`
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	Angle        float64 `json:"a"`
	MoveX        float64 `json:"mx"`
	MoveY        float64 `json:"my"`
	ShootTimer   int     `json:"-"`
	RebornTimer  int     `json:"rt"`
	BodyCount    int     `json:"-"`
	DeathCount   int     `json:"-"`
	RoomId       string  `json:"room_id"`
	PlayerClass  string  `json:"pc"`
	BulletSpeed  float64 `json:"-"`
	BulletLife   float64 `json:"-"`
	BulletDamage float64 `json:"-"`
}

type PlayerFinalState struct {
	Nickname string `json:"n"`
	Id       string `json:"id"`
	Kills    int    `json:"k"`
	Deaths   int    `json:"d"`
}

func NewPlayerGameState(id, nickname, playerClass string) *PlayerGameState {
	p := &PlayerGameState{
		Id:          id,
		Nickname:    nickname,
		Health:      MaxPlayerHealth,
		X:           PlayerSpawnPointX,
		Y:           PlayerSpawnPointY,
		Angle:       InitDirection,
		MoveX:       InitValue,
		MoveY:       InitValue,
		ShootTimer:  InitValue,
		RebornTimer: InitValue,
		BodyCount:   InitValue,
		DeathCount:  InitValue,
		PlayerClass: playerClass,
	}

	p.setupWeaponStats()
	return p
}

func (p *PlayerGameState) setupWeaponStats() {
	switch p.PlayerClass {
	case PlayerGun:
		p.BulletDamage = BulletDamageGun
		p.BulletSpeed = BulletSpeedGun
		p.BulletLife = BulletLifeGun
	case PlayerRifle:
		p.BulletDamage = BulletDamageRifle
		p.BulletSpeed = BulletSpeedRifle
		p.BulletLife = BulletLifeRifle
	case PlayerSom:
		p.BulletDamage = BulletDamageSom
		p.BulletSpeed = BulletSpeedSom
		p.BulletLife = BulletLifeSom
	}
}
