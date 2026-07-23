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
			if remaining%10 == 0 {
				log.Printf("Осталось времени: %d секунд", remaining)
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

	close(gl.updateChan)
	close(gl.deleteChan)
	close(gl.stopChan)

	if gl.finishChan != nil {
		select {
		case gl.finishChan <- true:
			log.Println("GameLoop: отправлен сигнал о завершении в лобби")
		default:
			log.Println("GameLoop: канал завершения уже содержит сигнал")
		}
	}
}

// === LOBBITOMIA ===

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
	if gl.game.GetCountOfPlayers() > 0 {
		gl.broadcastFinal(model.ServerEndMessage{
			State:  model.FinalGameState,
			Result: stats,
		})
	}
}

func (gl *GameLoop) getStatistics() []model.PlayerFinalState {
	stats := make([]model.PlayerFinalState, len(gl.game.GetAllPlayers()))
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

// ==================

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
			hit, player, _ := gl.collisionService.CheckBulletCollision(bullet)
			if hit {
				if player.Health > 0 {
					gl.collisionService.HandlePlayerHit(player, bullet)
				} else {
					//	gl.collisionService.HandleObjectHit(obj, bullet)
					log.Println("Bah in object")
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
		oldX, oldY := player.X, player.Y
		vectorLength := math.Sqrt(player.MoveX*player.MoveX + player.MoveY*player.MoveY)

		if vectorLength != 0 {
			player.MoveX /= vectorLength
			player.MoveY /= vectorLength
		}

		newX := player.X + player.MoveX*model.TickTime*model.PlayerSpeed
		newY := player.Y + player.MoveY*model.TickTime*model.PlayerSpeed
		player.X = newX
		if hit, _ := gl.collisionService.CheckPlayerObjectCollision(player); hit {
			player.X = oldX
		}
		player.Y = newY
		if hit, _ := gl.collisionService.CheckPlayerObjectCollision(player); hit {
			player.Y = oldY
		}
	}
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
