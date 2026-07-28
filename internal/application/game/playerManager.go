package game

import (
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/state"
)

type PlayerManager struct {
	gameState *state.GameState
}

func NewPlayerManager(state *state.GameState) *PlayerManager {
	return &PlayerManager{gameState: state}
}

func (pm *PlayerManager) UpdatePlayer(upd domain.ClientGameMessage) {
	player, exist := pm.gameState.GetPlayer(upd.Id)
	if !exist || player == nil {
		return
	}

	player.Angle = upd.Angle
	player.MoveX = upd.MX
	player.MoveY = upd.MY

	if pm.gameState.IsGameActive() && player.Health > 0 && upd.IsShoot && player.CooldownTimer <= 0 {
		pm.AddBullet(player)
		player.CooldownTimer = domain.ShootCooldown
	}
	if player.Health < 0 && player.RebornTimer != 0 {
		player.RebornTimer--
	} else if player.Health < 0 && player.RebornTimer == 0 {
		player.Health = domain.MaxPlayerHealth
	}
}

func (pm *PlayerManager) AddBullet(player *domain.PlayerGameState) {
	bullets := player.Weapon(player)
	pm.gameState.SetBullets(append(pm.gameState.GetAllBullets(), bullets...))
}
