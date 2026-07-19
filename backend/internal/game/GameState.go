package game

import (
	"fmt"
	"gamedevRooms/internal/model"
	"log"
	"math"
)

type GameState struct {
	players   map[string]*model.PlayerState
	bullets   []model.Bullet
	tickCount uint64
}

func NewGameState() *GameState {
	return &GameState{
		players: make(map[string]*model.PlayerState),
		bullets: make([]model.Bullet, 0),
	}
}

func (gs *GameState) GetAllBullets() []model.Bullet {
	return gs.bullets
}

func (gs *GameState) GetAllPlayers() map[string]*model.PlayerState {
	return gs.players
}

func (gs *GameState) GetPlayer(id string) (*model.PlayerState, bool) {
	player, exist := gs.players[id]
	return player, exist
}

func (gs *GameState) SetBullets(bullets []model.Bullet) {
	gs.bullets = bullets
}

func (gs *GameState) IncrementTick() {
	gs.tickCount++
}

func (gs *GameState) AddPlayer(player *model.PlayerState) {
	log.Println("Register ", player.Id)
	gs.players[player.Id] = player
}

func (gs *GameState) RemovePlayer(playerId string) {
	fmt.Println("Delete ", playerId)
	delete(gs.players, playerId)
}

func (gs *GameState) UpdatePlayer(upd model.ClientMessage) {
	player, exist := gs.GetPlayer(upd.Id)
	if !exist {
		return
	}

	player.Angle = upd.Angle
	player.MoveX = upd.MX
	player.MoveY = upd.MY

	if player.Health > 0 && upd.IsShoot && player.ShootTimer <= 0 {
		gs.addBullet(player)
		player.ShootTimer = model.ShootCooldown
	} else if player.Health < 0 && player.RebornTimer != 0 {
		player.RebornTimer--
	} else if player.Health < 0 && player.RebornTimer == 0 {
		player.Health = model.MaxPlayerHealth
	}
}

func (gs *GameState) addBullet(player *model.PlayerState) {
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
	gs.bullets = append(gs.bullets, bullet)
}
