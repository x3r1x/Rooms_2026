package game

import (
	"encoding/json"
	"gamedevRooms/internal/adapters/broadcast"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/state"
	"time"
)

type BroadcastManager struct {
	gameState        *state.GameState
	broadcastService *broadcast.BroadcastService
}

func NewBroadcastManager(state *state.GameState, broadcastService *broadcast.BroadcastService) *BroadcastManager {
	return &BroadcastManager{
		gameState:        state,
		broadcastService: broadcastService,
	}
}

func (bm *BroadcastManager) createRoomSnapshot(roomId string) ([]domain.PlayerGameState, []domain.Bullet) {
	players := make([]domain.PlayerGameState, 0)
	bullets := make([]domain.Bullet, 0)
	for _, player := range bm.gameState.GetAllPlayers() {
		if bm.gameState.GetPlayerRoom(player.Id) == roomId {
			players = append(players, *player)
		}
	}
	for _, bullet := range bm.gameState.GetAllBullets() {
		owner, exists := bm.gameState.GetPlayer(bullet.OwnerId)
		if exists && bm.gameState.GetPlayerRoom(owner.Id) == roomId {
			bullets = append(bullets, bullet)
		}
	}
	return players, bullets
}

func (bm *BroadcastManager) broadcast() {
	playersByRoom := bm.gameState.GetPlayersByRoom()
	bullets := bm.gameState.GetAllBullets()
	gameTime := float64(time.Since(bm.gameState.GetGameStartTime()).Milliseconds())
	statistic := bm.getInGameStatistics()
	for _, playersInRoom := range playersByRoom {
		msg := domain.ServerGameMessage{
			State:     domain.OngoingGameState,
			Time:      gameTime,
			Players:   playersInRoom,
			Bullets:   bullets,
			Statistic: statistic,
		}

		data, err := json.Marshal(msg)
		if err != nil {
			continue
		}

		for _, p := range playersInRoom {
			bm.broadcastService.BroadcastToPlayer(p.Id, data)
		}
	}
}

func (bm *BroadcastManager) broadcastFinal(message domain.ServerEndMessage) {
	bm.broadcastService.BroadcastToAll(message)
}

func (bm *BroadcastManager) getInGameStatistics() []domain.PlayerStatistic {
	stats := make([]domain.PlayerStatistic, 0, len(bm.gameState.GetAllPlayers()))
	for _, player := range bm.gameState.GetAllPlayers() {
		stats = append(stats, domain.PlayerStatistic{
			Id:     player.Id,
			Kills:  player.BodyCount,
			Deaths: player.DeathCount,
			Hp:     player.Health,
		})
	}
	return stats
}

func (bm *BroadcastManager) getFinalStatistics() []domain.PlayerFinalState {
	stats := make([]domain.PlayerFinalState, 0, len(bm.gameState.GetAllPlayers()))
	for _, player := range bm.gameState.GetAllPlayers() {
		stats = append(stats, domain.PlayerFinalState{
			Id:     player.Id,
			Deaths: player.DeathCount,
			Kills:  player.BodyCount,
		})
	}
	return stats
}
