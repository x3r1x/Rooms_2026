package game

import (
	"encoding/json"
	"gamedevRooms/internal/model"
	"log"
	"math"
	"time"

	"github.com/gorilla/websocket"
)

type GameLoop struct {
	game             *GameState
	collisionService *CollisionService
	updateChan       chan model.ClientGameMessage
	deleteChan       chan string
	finishChan       chan bool
	stopChan         chan bool
}

func NewGameLoop(game *GameState, finishChan chan bool) *GameLoop {
	return &GameLoop{game: game,
		collisionService: NewCollisionService(game),
		updateChan:       make(chan model.ClientGameMessage),
		deleteChan:       make(chan string),
		finishChan:       finishChan,
		stopChan:         make(chan bool)}
}

func (gl *GameLoop) UpdatePlayer(msg model.ClientGameMessage) {
	if _, exists := gl.game.GetPlayer(msg.Id); !exists {
		return
	}
	gl.updateChan <- msg
}

func (gl *GameLoop) DeletePlayer(id string) {
	gl.deleteChan <- id
}

func (gl *GameLoop) Run() {
	ticker := time.NewTicker(model.TickTime * time.Millisecond)
	defer ticker.Stop()
	defer gl.cleanup()
	gl.startGame()

	for {
		select {
		case del := <-gl.deleteChan:
			gl.handleDelete(del)
			if gl.shouldStop() {
				log.Println("GameLoop: остановка после удаления игрока")
				return
			}
		case upload := <-gl.updateChan:
			gl.game.UpdatePlayer(upload)
		case <-ticker.C:
			if !gl.game.IsGameActive() {
				continue
			}
			gl.game.IncrementTick()
			gl.updateShooterTimers()
			gl.updateBullets()
			gl.updatePlayers()
			remaining := gl.game.GetRemainingSeconds()
			if remaining <= 0 {
				gl.endGame()
				return
			}
			gl.broadcast(model.ServerGameMessage{
				State:   model.OngoingGameState,
				Type:    "a",
				Players: gl.createSnapshot(),
				Bullets: gl.game.GetAllBullets(),
			})
		}
	}
}

func (gl *GameLoop) shouldStop() bool {
	return !gl.game.IsGameActive() || len(gl.game.GetAllPlayers()) == 0
}

func (gl *GameLoop) cleanup() {
	log.Println("GameLoop: очистка ресурсов...")
	gl.game.SetBullets([]model.Bullet{})
	gl.game.SetObjects(make(map[string]*model.Object))
	gl.game.SetGameActive(false)
	if gl.finishChan != nil {
		select {
		case gl.finishChan <- true:
			log.Println("GameLoop: отправлен сигнал о завершении в лобби")
		default:
			log.Println("GameLoop: канал завершения уже содержит сигнал")
		}
	}
}

func (gl *GameLoop) handleDelete(id string) {
	gl.game.RemovePlayer(id)
	log.Printf("Игрок %s удален. Осталось: %d", id, len(gl.game.GetAllPlayers()))
	if gl.game.IsGameActive() && len(gl.game.GetAllPlayers()) <= 1 {
		log.Println("Игрок вышел, игра завершена досрочно")
		gl.endGame()
		select {
		case gl.stopChan <- true:
			log.Println("GameLoop: отправлен сигнал остановки")
		default:
			log.Println("GameLoop: stopChan уже содержит сигнал")
		}
	}
}

func (gl *GameLoop) startGame() {
	if gl.game.IsGameActive() {
		return
	}
	gl.game.SetGameActive(true)
}

func (gl *GameLoop) endGame() {
	if !gl.game.IsGameActive() {
		return
	}
	gl.game.SetGameActive(false)
	stats := gl.getStatistics()
	log.Println(stats)
	if gl.game.GetCountOfPlayers() > 0 {
		gl.broadcastFinal(model.ServerEndMessage{
			State:  model.FinalGameState,
			Result: stats,
		})
	}
}

func (gl *GameLoop) getStatistics() []model.PlayerFinalState {
	stats := make([]model.PlayerFinalState, 0, len(gl.game.GetAllPlayers()))
	for _, player := range gl.game.GetAllPlayers() {
		stats = append(stats, model.PlayerFinalState{
			Nickname: player.Nickname,
			Id:       player.Id,
			Deaths:   player.DeathCount,
			Kills:    player.BodyCount,
		})
	}
	return stats
}

func (gl *GameLoop) updateShooterTimers() {
	for _, player := range gl.game.GetAllPlayers() {
		if player.ShootTimer > 0 {
			player.ShootTimer--
		}
	}
}

func (gl *GameLoop) updateBullets() {
	activeBullets := make([]model.Bullet, 0)
	for _, bullet := range gl.game.GetAllBullets() {
		bullet.Life--
		bullet.X += math.Cos(bullet.Direction) * model.MaxBulletSpeed
		bullet.Y += math.Sin(bullet.Direction) * model.MaxBulletSpeed

		if bullet.Life > 0 {
			// ИСПРАВЛЕНИЕ: Сохраняем объект, в который попала пуля
			hit, player, obj := gl.collisionService.CheckBulletCollision(bullet)
			if hit {
				if player != nil && player.Health > 0 {
					gl.collisionService.HandlePlayerHit(player, bullet)
				} else if obj != nil {
					log.Printf("Пуля попала в стену ID: %s", obj.Id)
				}
				continue
			}
			activeBullets = append(activeBullets, bullet)
		}
	}
	gl.game.SetBullets(activeBullets)
}

func (gl *GameLoop) updatePlayers() {
	for _, player := range gl.game.GetAllPlayers() {
		if player.Health <= 0 {
			if player.RebornTimer > 0 {
				player.RebornTimer--
			} else if player.RebornTimer == 0 {
				player.Health = model.MaxPlayerHealth
				player.RebornTimer = model.PlayerRebornTimer
			}
			continue
		}

		vectorLength := math.Sqrt(player.MoveX*player.MoveX + player.MoveY*player.MoveY)
		var moveX, moveY float64
		if vectorLength != 0 {
			moveX = player.MoveX / vectorLength
			moveY = player.MoveY / vectorLength
		}

		deltaX := moveX * model.TickTime * model.PlayerSpeed
		deltaY := moveY * model.TickTime * model.PlayerSpeed

		nextX := player.X + deltaX
		nextY := player.Y + deltaY

		if gl.canMoveTo(nextX, nextY, player) {
			player.X = nextX
			player.Y = nextY
		}

		if hit, _ := gl.collisionService.CheckPlayerObjectCollision(player); hit {
			gl.collisionService.ResolvePlayerCollisionSmooth(player)
		}

		if hit, direction, targetRoomId := gl.collisionService.CheckPlayerExitCollision(player); hit {
			gl.handleRoomTransition(player, direction, targetRoomId)
		}
	}
}

func (gl *GameLoop) canMoveTo(nextX, nextY float64, player *model.PlayerGameState) bool {
	oldX, oldY := player.X, player.Y

	bestX, bestY := oldX, oldY
	moved := false

	player.X = nextX
	if hit, _ := gl.collisionService.CheckPlayerObjectCollision(player); !hit {
		bestX = nextX
		moved = true
	}
	player.X = oldX

	player.Y = nextY
	if hit, _ := gl.collisionService.CheckPlayerObjectCollision(player); !hit {
		bestY = nextY
		moved = true
	}
	player.Y = oldY

	player.X = nextX
	player.Y = nextY
	if hit, _ := gl.collisionService.CheckPlayerObjectCollision(player); !hit {
		bestX = nextX
		bestY = nextY
		moved = true
	} else {
		player.X = oldX
		player.Y = oldY
	}

	player.X = bestX
	player.Y = bestY

	if hit, _ := gl.collisionService.CheckPlayerObjectCollision(player); hit {
		gl.collisionService.ResolvePlayerCollisionSmooth(player)
	}

	return moved
}
func (gl *GameLoop) handleRoomTransition(player *model.PlayerGameState, direction, targetRoomId string) {
	roomPixelWidth := float64(model.RoomWidth * int(model.TileSize))
	roomPixelHeight := float64(model.RoomHeight * int(model.TileSize))
	halfSize := model.PlayerHalfSize

	switch direction {
	case model.TopMarker:
		player.Y = roomPixelHeight - halfSize - 1
	case model.BottomMarker:
		player.Y = halfSize + 1
	case model.LeftMarker:
		player.X = roomPixelWidth - halfSize - 1
	case model.RightMarker:
		player.X = halfSize + 1
	}
	gl.game.SetPlayerRoom(player.Id, targetRoomId)
}

func (gl *GameLoop) createSnapshot() []model.PlayerGameState {
	snapshot := make([]model.PlayerGameState, 0, len(gl.game.GetAllPlayers()))
	for _, player := range gl.game.GetAllPlayers() {
		snapshot = append(snapshot, *player)
	}
	return snapshot
}

func (gl *GameLoop) broadcast(message model.ServerGameMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}
	for id, p := range gl.game.GetAllPlayers() {
		if p.Connection == nil {
			continue
		}
		err := p.Connection.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Ошибка отправки: ", err)
			gl.deleteChan <- id
		}
	}
}

func (gl *GameLoop) broadcastFinal(message model.ServerEndMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}
	for id, p := range gl.game.GetAllPlayers() {
		if p.Connection == nil {
			continue
		}
		err := p.Connection.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Ошибка отправки: ", err)
			gl.deleteChan <- id
		}
	}
}
