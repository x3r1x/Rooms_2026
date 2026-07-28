package game

import (
	"gamedevRooms/internal/adapters/broadcast"
	"gamedevRooms/internal/adapters/collision"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/state"
	"log"
	"time"
)

type GameService struct {
	gameState        *state.GameState
	playerManager    *PlayerManager
	physicsEngine    *PhysicsEngine
	broadcastManager *BroadcastManager
	updateChan       chan domain.ClientGameMessage
	deleteChan       chan string
	finishChan       chan bool
	onGameEnd        func()
}

func NewGameService(
	gameState *state.GameState,
	collisionService *collision.CollisionService,
	broadcastService *broadcast.BroadcastService,
	onGameEnd func(),
) *GameService {
	pm := NewPlayerManager(gameState)
	pe := NewPhysicsEngine(gameState, collisionService)
	bm := NewBroadcastManager(gameState, broadcastService)
	return &GameService{
		gameState:        gameState,
		playerManager:    pm,
		physicsEngine:    pe,
		broadcastManager: bm,
		updateChan:       make(chan domain.ClientGameMessage),
		deleteChan:       make(chan string),
		finishChan:       make(chan bool, 1),
		onGameEnd:        onGameEnd,
	}
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

func (gs *GameService) Stop() {
	gs.finishChan <- true
}

func (gs *GameService) Run() {
	ticker := time.NewTicker(domain.TickTime * time.Millisecond)
	defer ticker.Stop()

	gs.startGame()

	for {
		select {
		case <-gs.finishChan:
			gs.endGame()
			return
		case del := <-gs.deleteChan:
			gs.handleDelete(del)
		case upload := <-gs.updateChan:
			gs.playerManager.UpdatePlayer(upload)
		case <-ticker.C:
			gs.gameState.IncrementTick()
			gs.physicsEngine.updateShooterTimers()
			gs.physicsEngine.updateBullets()
			gs.physicsEngine.updatePlayers()
			remaining := gs.gameState.GetRemainingSeconds()
			if remaining <= 0 {
				gs.Stop()
			}
			gs.broadcastManager.broadcast()
		}
	}
}

func (gs *GameService) handleDelete(id string) {
	gs.gameState.RemovePlayer(id)
	log.Printf("Игрок %s удален. Осталось: %d", id, len(gs.gameState.GetAllPlayers()))
	if gs.gameState.IsGameActive() && len(gs.gameState.GetAllPlayers()) <= 1 {
		log.Println("Игрок вышел, игра завершена досрочно")
		gs.Stop()
		return
	}
}

func (gs *GameService) startGame() {
	if gs.gameState.IsGameActive() {
		return
	}
	gs.gameState.SetGameActive(true)
}

func (gs *GameService) endGame() {
	if !gs.gameState.IsGameActive() {
		return
	}
	gs.gameState.SetGameActive(false)
	stats := gs.broadcastManager.getFinalStatistics()
	if gs.gameState.GetCountOfPlayers() > 0 {
		gs.broadcastManager.broadcastFinal(domain.ServerEndMessage{
			State:  domain.FinalGameState,
			Result: stats,
		})
	}
	if gs.onGameEnd != nil {
		gs.onGameEnd()
	}
}
