package game

import "gamedevRooms/internal/domain"

type PlayerManager struct {
	state *GameState
}

func NewPlayerManager(state *GameState) *PlayerManager {
	return &PlayerManager{state: state}
}

func (gs *GameState) AddPlayer(player *domain.PlayerGameState) {
	gs.players[player.Id] = player
}

func (gs *GameState) RemovePlayer(playerId string) {
	delete(gs.players, playerId)
}

func (gs *GameService) UpdatePlayer(msg domain.ClientGameMessage) {
	if _, exists := gs.gameState.GetPlayer(msg.Id); !exists {
		return
	}
	gs.updateChan <- msg
}

func (gs *GameService) DeletePlayer(id string) {
	gs.deleteChan <- id
}
