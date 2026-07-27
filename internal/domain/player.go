package domain

type PlayerGameState struct {
	Id           string  `json:"id"`
	Health       float64 `json:"h"`
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	Angle        float64 `json:"a"`
	MoveX        float64 `json:"mx"`
	MoveY        float64 `json:"my"`
	RebornTimer  int     `json:"rt"`
	RoomId       string  `json:"room_id"`
	PlayerClass  string  `json:"pc"`
	Nickname     string  `json:"-"`
	ShootTimer   int     `json:"-"`
	BodyCount    int     `json:"-"`
	DeathCount   int     `json:"-"`
	BulletSpeed  float64 `json:"-"`
	BulletLife   float64 `json:"-"`
	BulletDamage float64 `json:"-"`
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
