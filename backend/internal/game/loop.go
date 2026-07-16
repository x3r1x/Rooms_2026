package game

//func (r *GameRoom) GameLoop(ctx context.Context) {
//	ticker := time.NewTicker(16 * time.Millisecond)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ctx.Done():
//			return
//
//		case id := <-r.RegisterChan:
//			r.players[id] = Player{X: 400, Y: 300, HP: 100}
//
//		case id := <-r.LeaveChan:
//			delete(r.players, id)
//
//		case input := <-r.InputChan:
//			r.handleInput(input)
//
//		case <-ticker.C:
//			r.tickCount++
//			r.updatePhysics()
//			r.broadcastState()
//		}
//	}
//}
