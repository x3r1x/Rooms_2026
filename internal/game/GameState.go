package game

import (
	"fmt"
	"gamedevRooms/internal/factory"
	"gamedevRooms/internal/model"
	"log"
	"time"
)

type GameState struct {
	players       map[string]*model.PlayerGameState
	bullets       []model.Bullet
	tickCount     uint64
	isGameActive  bool
	gameStartTime time.Time
	gameDuration  time.Duration
}

func NewGameState() *GameState {
	return &GameState{
		players:      make(map[string]*model.PlayerGameState),
		bullets:      make([]model.Bullet, 0),
		isGameActive: false,
		gameDuration: 60 * time.Second,
	}
}

// === LOBBITOMIA ===

func (gs *GameState) IsGameActive() bool {
	return gs.isGameActive
}

func (gs *GameState) SetGameActive(active bool) {
	gs.isGameActive = active
	if active {
		gs.gameStartTime = time.Now()
		log.Println("=== Game is active ===")
	}
}

func (gs *GameState) GetRemainingSeconds() int {
	if !gs.isGameActive {
		return int(gs.gameDuration.Seconds())
	}
	elapsed := time.Since(gs.gameStartTime)
	if elapsed >= gs.gameDuration {
		return 0
	}
	return int((gs.gameDuration - elapsed).Seconds())
}

func (gs *GameState) IsLobbyEmpty() bool {
	return len(gs.players) == 0
}

func (gs *GameState) IsLobbyFull() bool {
	return len(gs.players) >= 4
}

func (gs *GameState) HasMinimumPlayer() bool {
	return len(gs.players) >= 2
}

func (gs *GameState) CanAddPlayer() bool {
	return !gs.IsGameActive() && !gs.IsLobbyFull()
}

// ==================

func (gs *GameState) GetAllBullets() []model.Bullet {
	return gs.bullets
}

func (gs *GameState) GetAllPlayers() map[string]*model.PlayerGameState {
	return gs.players
}

func (gs *GameState) GetPlayer(id string) (*model.PlayerGameState, bool) {
	player, exist := gs.players[id]
	return player, exist
}

func (gs *GameState) SetBullets(bullets []model.Bullet) {
	gs.bullets = bullets
}

func (gs *GameState) IncrementTick() {
	gs.tickCount++
}

func (gs *GameState) AddPlayer(player *model.PlayerGameState) {
	log.Println("Register ", player.Id)
	gs.players[player.Id] = player
}

func (gs *GameState) RemovePlayer(playerId string) {
	fmt.Println("Delete ", playerId)
	delete(gs.players, playerId)
}

func (gs *GameState) UpdatePlayer(upd model.ClientGameMessage) {
	player, exist := gs.GetPlayer(upd.Id)
	if !exist {
		return
	}

	player.Angle = upd.Angle
	player.MoveX = upd.MX
	player.MoveY = upd.MY

	if gs.isGameActive && player.Health > 0 && upd.IsShoot && player.ShootTimer <= 0 {
		gs.addBullet(player)
		player.ShootTimer = model.ShootCooldown
	}
	if player.Health < 0 && player.RebornTimer != 0 {
		player.RebornTimer--
	} else if player.Health < 0 && player.RebornTimer == 0 {
		player.Health = model.MaxPlayerHealth
	}
}

func (gs *GameState) addBullet(player *model.PlayerGameState) {
	bullet := factory.BulletFactory(player)
	gs.bullets = append(gs.bullets, bullet)
}
