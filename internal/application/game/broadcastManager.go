package game

import (
	"encoding/json"
	"gamedevRooms/internal/adapters/broadcast"
	"gamedevRooms/internal/domain"
	"time"
)

type BroadcastManager struct {
	state            *GameState
	broadcastService *broadcast.BroadcastService
}

func NewBroadcastManager(state *GameState, broadcastService *broadcast.BroadcastService) *BroadcastManager {
	return &BroadcastManager{
		state:            state,
		broadcastService: broadcastService,
	}
}

func (gs *GameService) createRoomSnapshot(roomId string) ([]domain.PlayerGameState, []domain.Bullet) {
	players := make([]domain.PlayerGameState, 0)
	bullets := make([]domain.Bullet, 0)
	for _, player := range gs.gameState.GetAllPlayers() {
		if gs.gameState.GetPlayerRoom(player.Id) == roomId {
			players = append(players, *player)
		}
	}
	for _, bullet := range gs.gameState.GetAllBullets() {
		owner, exists := gs.gameState.GetPlayer(bullet.OwnerId)
		if exists && gs.gameState.GetPlayerRoom(owner.Id) == roomId {
			bullets = append(bullets, bullet)
		}
	}
	return players, bullets
}

func (gs *GameService) broadcast() {
	playersByRoom := gs.gameState.GetPlayersByRoom()
	bullets := gs.gameState.GetAllBullets()
	gameTime := float64(time.Since(gs.gameState.GetGameStartTime()).Milliseconds())
	statistic := gs.getInGameStatistics()
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
			gs.broadcastService.BroadcastToPlayer(p.Id, data)
		}
	}
}

func (gs *GameService) broadcastFinal(message domain.ServerEndMessage) {
	gs.broadcastService.BroadcastToAll(message)
}

func (gs *GameService) getInGameStatistics() []domain.PlayerStatistic {
	stats := make([]domain.PlayerStatistic, 0, len(gs.gameState.GetAllPlayers()))
	for _, player := range gs.gameState.GetAllPlayers() {
		stats = append(stats, domain.PlayerStatistic{
			Id:     player.Id,
			Kills:  player.BodyCount,
			Deaths: player.DeathCount,
			Hp:     player.Health,
		})
	}
	return stats
}

func (gs *GameService) getFinalStatistics() []domain.PlayerFinalState {
	stats := make([]domain.PlayerFinalState, 0, len(gs.gameState.GetAllPlayers()))
	for _, player := range gs.gameState.GetAllPlayers() {
		stats = append(stats, domain.PlayerFinalState{
			Id:     player.Id,
			Deaths: player.DeathCount,
			Kills:  player.BodyCount,
		})
	}
	return stats
}
