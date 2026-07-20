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
	updateChan       chan model.ClientMessage
	registerChan     chan *model.PlayerState
	deleteChan       chan string
}

func NewGameLoop(game *GameState) *GameLoop {
	return &GameLoop{game: game,
		collisionService: NewCollisionService(game),
		updateChan:       make(chan model.ClientMessage),
		registerChan:     make(chan *model.PlayerState),
		deleteChan:       make(chan string),
	}
}

func (gl *GameLoop) RegisterPlayer(player *model.PlayerState) {
	gl.registerChan <- player
}

func (gl *GameLoop) UpdatePlayer(msg model.ClientMessage) {
	gl.updateChan <- msg
}

func (gl *GameLoop) DeletePlayer(id string) {
	gl.deleteChan <- id
}
func (gl *GameLoop) Run() {
	ticker := time.NewTicker(model.TickTime * time.Millisecond)
	defer ticker.Stop()

	for {
		select {

		case reg := <-gl.registerChan:
			gl.handleRegister(reg)
		case del := <-gl.deleteChan:
			gl.handleDelete(del)
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

			gl.broadcast(model.ServerMessage{
				Type:    "a",
				Players: gl.createSnapshot(),
				Bullets: gl.game.GetAllBullets(),
			})
		}
	}
}

// === LOBBITOMIA ===

func (gl *GameLoop) handleRegister(player *model.PlayerState) {
	if !gl.game.CanAddPlayer() {
		log.Printf(" Отклонено подключение %s: игра активна или лобби заполнено", player.Id)
		if player.Connection != nil {
			if err := player.Connection.Close(); err != nil {
				log.Println("Подключение невозможно закрыть")
			}
		}
		return
	}

	gl.game.AddPlayer(player)
	log.Printf("Игрок %s подключился. Всего: %d/4", player.Id, len(gl.game.GetAllPlayers()))

	if gl.game.IsLobbyFull() {
		log.Println("Лобби заполнено")
		gl.startGame()
	}
}

func (gl *GameLoop) handleDelete(id string) {
	gl.game.RemovePlayer(id)
	log.Printf("Игрок %s удален. Осталось: %d", id, len(gl.game.GetAllPlayers()))

	if gl.game.IsGameActive() {
		if len(gl.game.GetAllPlayers()) <= 1 {
			log.Println("Игрок вышел, игра завершена досрочно")
			gl.endGame()
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
	winner := gl.getWinner()
	log.Println(winner)
	gl.resetGame()
}

func (gl *GameLoop) getWinner() string {
	var winner string
	maxHealth := -1.0

	for _, player := range gl.game.GetAllPlayers() {
		if player.Health > maxHealth {
			maxHealth = player.Health
			winner = player.Id
		}
	}

	if winner == "" {
		return "Никто"
	}
	return winner
}

func (gl *GameLoop) resetGame() {
	gl.game.SetBullets([]model.Bullet{})
	for _, player := range gl.game.GetAllPlayers() {
		player.Health = model.MaxPlayerHealth
		player.X = model.PlayerSpawnPointX
		player.Y = model.PlayerSpawnPointY
		player.Angle = model.InitDirection
		player.RebornTimer = 0
		player.ShootTimer = 0
	}
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
			hit, player := gl.collisionService.CheckBulletCollision(bullet)
			if hit && player.Health > 0 {
				gl.collisionService.HandleHit(player, bullet)
			} else {
				activeBullets = append(activeBullets, bullet)
			}
		}
	}
	gl.game.SetBullets(activeBullets)
}

func (gl *GameLoop) updatePlayers() {
	for _, player := range gl.game.GetAllPlayers() {
		vectorLength := math.Sqrt(player.MoveX*player.MoveX + player.MoveY*player.MoveY)

		if vectorLength != 0 {
			player.MoveX /= vectorLength
			player.MoveY /= vectorLength
		}

		player.X += player.MoveX * model.TickTime * model.PlayerSpeed
		player.Y += player.MoveY * model.TickTime * model.PlayerSpeed
	}
}

func (gl *GameLoop) createSnapshot() []model.PlayerState {
	snapshot := make([]model.PlayerState, 0, len(gl.game.GetAllPlayers()))
	for _, player := range gl.game.GetAllPlayers() {
		snapshot = append(snapshot, *player)
	}
	return snapshot
}

func (gl *GameLoop) broadcast(message model.ServerMessage) {
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
