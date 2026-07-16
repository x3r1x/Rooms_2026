package game

import "gamedevRooms/internal/models"

func takeSnapshot() []models.PlayerState {
	snapshot := make([]models.PlayerState, 0, len(models.Game.Players))
	for _, player := range models.Game.Players {
		snapshot = append(snapshot, *player)
	}
	return snapshot
}
