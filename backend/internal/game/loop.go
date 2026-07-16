package game

import (
	"gamedevRooms/internal/models"
	"time"
)

func GoGameLoop(gameState *models.GameState) {
	ticker := time.NewTicker(16 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {

		case reg := <-models.Game.RegisterChan:
			models.Game.Players[reg.Id] = &models.PlayerState{
				Id:         reg.Id,
				X:          reg.X,
				Y:          reg.Y,
				Direction:  reg.Direction,
				Bullets:    reg.Bullets,
				Connection: reg.Connection,
			}
		case del := <-models.Game.LeaveChan:
			delete(gameState.Players, del)

		case upload := <-models.Game.InputChan:
			if player, exist := models.Game.Players[upload.Player.Id]; exist {
				player.X = upload.Player.X
				player.Y = upload.Player.Y
				player.Direction = upload.Player.Direction
				player.Bullets = upload.Player.Bullets
			}
		case <-ticker.C:
			gameState.TickCount++
			//gameState.updatePhysics()
			broadcast(models.ServerMessage{Players: takeSnapshot()})
		}
	}
}
