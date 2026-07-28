package game

import (
	"gamedevRooms/internal/adapters/broadcast"
	"gamedevRooms/internal/adapters/collision"
	_map "gamedevRooms/internal/adapters/map"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/ports"
	"log"
	"time"
)

type GameService struct {
	gameState        ports.GameStateProvider
	collisionService *collision.CollisionService
	mapManager       *_map.MapManager
	broadcastService *broadcast.BroadcastService
	updateChan       chan domain.ClientGameMessage
	deleteChan       chan string
	finishChan       chan bool
	onGameEnd        func()
}

func NewGameService(
	gameState ports.GameStateProvider,
	collisionService *collision.CollisionService,
	mapManager *_map.MapManager,
	broadcastService *broadcast.BroadcastService,
	onGameEnd func(),
) *GameService {
	return &GameService{
		gameState:        gameState,
		collisionService: collisionService,
		mapManager:       mapManager,
		broadcastService: broadcastService,
		updateChan:       make(chan domain.ClientGameMessage),
		deleteChan:       make(chan string),
		finishChan:       make(chan bool, 1),
		onGameEnd:        onGameEnd,
	}
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
			gs.gameState.UpdatePlayer(upload)
		case <-ticker.C:
			gs.gameState.IncrementTick()
			gs.updateShooterTimers()
			gs.updateBullets()
			gs.updatePlayers()
			remaining := gs.gameState.GetRemainingSeconds()
			if remaining <= 0 {
				gs.Stop()
			}
			gs.broadcast()
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
	stats := gs.getFinalStatistics()
	log.Println(stats)
	if gs.gameState.GetCountOfPlayers() > 0 {
		gs.broadcastFinal(domain.ServerEndMessage{
			State:  domain.FinalGameState,
			Result: stats,
		})
	}
	if gs.onGameEnd != nil {
		gs.onGameEnd()
	}
}
