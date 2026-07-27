package game

import (
	"encoding/json"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/ports"
	"log"
	"math"
	"time"
)

type GameService struct {
	gameState        ports.GameStateProvider
	collisionService ports.CollisionService
	mapManager       ports.MapManager
	broadcastService ports.BroadcastService
	updateChan       chan domain.ClientGameMessage
	deleteChan       chan string
	finishChan       chan bool
	onGameEnd        func()
}

func NewGameService(
	gameState ports.GameStateProvider,
	collisionService ports.CollisionService,
	mapManager ports.MapManager,
	broadcastService ports.BroadcastService,
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

func (gs *GameService) updateShooterTimers() {
	for _, player := range gs.gameState.GetAllPlayers() {
		if player.CooldownTimer > 0 {
			player.CooldownTimer--
		}
	}
}

func (gs *GameService) updateBullets() {
	activeBullets := make([]domain.Bullet, 0)
	for _, bullet := range gs.gameState.GetAllBullets() {
		if bullet.Life > 0 {
			bullet.Move()
			bullet.Life--
			hit, player, obj := gs.collisionService.CheckBulletCollision(bullet)
			if hit {
				if player != nil && player.Health > 0 {
					gs.collisionService.HandlePlayerHit(player, bullet)
				} else if obj != nil || bullet.Type == domain.PlayerSom {
					if bullet.Type == domain.PlayerSom {
						gs.collisionService.TriggerExplosion(bullet)
					}
					log.Printf("Пуля попала в стену ID: %s", obj.Id)
				}
				continue
			}
			activeBullets = append(activeBullets, bullet)
		}
	}
	gs.gameState.SetBullets(activeBullets)
}

func (gs *GameService) updatePlayers() {
	for _, player := range gs.gameState.GetAllPlayers() {
		if player.Health <= 0 {
			if player.RebornTimer > 0 {
				player.RebornTimer--
			} else if player.RebornTimer == 0 {
				player.Health = domain.MaxPlayerHealth
				player.RebornTimer = domain.PlayerRebornTimer
				player.X = domain.PlayerSpawnPointX
				player.Y = domain.PlayerSpawnPointY
			}
			continue
		}

		if player.MoveX == 0 && player.MoveY == 0 {
			continue
		}

		vectorLength := math.Sqrt(player.MoveX*player.MoveX + player.MoveY*player.MoveY)
		var moveX, moveY float64
		if vectorLength != 0 {
			moveX = player.MoveX / vectorLength
			moveY = player.MoveY / vectorLength
		}

		deltaX := moveX * domain.TickTime * domain.PlayerSpeed
		deltaY := moveY * domain.TickTime * domain.PlayerSpeed

		player.X += deltaX
		if hit, _ := gs.collisionService.CheckPlayerObjectCollision(player); hit {
			player.X -= deltaX
		}

		player.Y += deltaY
		if hit, _ := gs.collisionService.CheckPlayerObjectCollision(player); hit {
			player.Y -= deltaY
		}

		if hit, _ := gs.collisionService.CheckPlayerObjectCollision(player); hit {
			gs.collisionService.ResolvePlayerCollisionSmooth(player)
		}

		if hit, direction, targetRoomId := gs.collisionService.CheckPlayerExitCollision(player); hit {
			gs.handleRoomTransition(player, direction, targetRoomId)
		}
	}
}

func (gs *GameService) handleRoomTransition(player *domain.PlayerGameState, direction, targetRoomId string) {
	roomPixelWidth := float64(domain.RoomWidth * int(domain.TileSize))
	roomPixelHeight := float64(domain.RoomHeight * int(domain.TileSize))
	halfSize := domain.PlayerHalfSize

	switch direction {
	case domain.TopMarker:
		player.Y = roomPixelHeight - halfSize - 1
	case domain.BottomMarker:
		player.Y = halfSize + 1
	case domain.LeftMarker:
		player.X = roomPixelWidth - halfSize - 1
	case domain.RightMarker:
		player.X = halfSize + 1
	}
	gs.gameState.SetPlayerRoom(player.Id, targetRoomId)
}
