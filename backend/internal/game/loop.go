package game

import (
	"context"
	"gamedevRooms/internal/models"
	"time"
)

func GameLoop(gameState *models.GameState, ctx context.Context) {
	ticker := time.NewTicker(16 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case reg := <-gameState.RegisterChan:
			models.Game.Players[reg.Id] = &models.PlayerState{
				Id:         reg.Id,
				X:          reg.X,
				Y:          reg.Y,
				Direction:  reg.Direction,
				Bullets:    reg.Bullets,
				Connection: reg.Connection,
			}
		case id := <-gameState.LeaveChan:
			delete(gameState.Players, id)

		//case upload := <-gameState.InputChan:
		//	//gameState.handleInput(input)

		case <-ticker.C:
			gameState.TickCount++
			//gameState.updatePhysics()
			//gameState.BroadcastState()
		}
	}
}
