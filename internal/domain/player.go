package domain

import "math"

type PlayerGameState struct {
	Id          string  `json:"id"`
	Nickname    string  `json:"-"`
	Health      float64 `json:"h"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Angle       float64 `json:"a"`
	MoveX       float64 `json:"mx"`
	MoveY       float64 `json:"my"`
	ShootTimer  int     `json:"-"`
	RebornTimer int     `json:"rt"`
	BodyCount   int     `json:"-"`
	DeathCount  int     `json:"-"`
	RoomId      string  `json:"room_id"`
	PlayerClass string  `json:"pc"`
}

type PlayerFinalState struct {
	Nickname string `json:"n"`
	Id       string `json:"id"`
	Kills    int    `json:"k"`
	Deaths   int    `json:"d"`
}

func (p *PlayerGameState) IsAlive() bool {
	return p.Health > 0
}

func (p *PlayerGameState) IsDead() bool {
	return p.Health <= 0
}

func (p *PlayerGameState) TakeDamage(damage float64) {
	if damage < 0 {
		return
	}
	p.Health -= damage
	if p.Health < 0 {
		p.Health = 0
	}
}

func (p *PlayerGameState) CanShoot() bool {
	return p.IsAlive() && p.ShootTimer == 0
}

func (p *PlayerGameState) StartShootCooldown() {
	p.ShootTimer = ShootCooldown
}

func (p *PlayerGameState) UpdateShootTimer() {
	if p.ShootTimer > 0 {
		p.ShootTimer--
	}
}

func (p *PlayerGameState) UpdateRebornTimer() {
	if p.IsDead() && p.RebornTimer > 0 {
		p.RebornTimer--
	}
}

func (p *PlayerGameState) CanReborn() bool {
	return p.IsDead() && p.RebornTimer == 0
}

func (p *PlayerGameState) Reborn() {
	p.Health = MaxPlayerHealth
	p.RebornTimer = PlayerRebornTimer
	p.X = PlayerSpawnPointX
	p.Y = PlayerSpawnPointY
}

func (p *PlayerGameState) AddKill() {
	p.BodyCount++
}

func (p *PlayerGameState) AddDeath() {
	p.DeathCount++
}

func (p *PlayerGameState) GetMoveVector() (float64, float64) {
	length := math.Sqrt(p.MoveX*p.MoveX + p.MoveY*p.MoveY)
	if length == 0 {
		return 0, 0
	}
	return p.MoveX / length, p.MoveY / length
}

func (p *PlayerGameState) GetBarrelPosition() (float64, float64) {
	barrelLength := PlayerHalfSize + BulletBarrelOffset
	x := p.X + barrelLength*math.Cos(p.Angle)
	y := p.Y + barrelLength*math.Sin(p.Angle)
	return x, y
}
