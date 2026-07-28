package game

import (
	"gamedevRooms/internal/adapters/collision"
	"gamedevRooms/internal/domain"
	"log"
	"math"
)

type PhysicsEngine struct {
	state            *GameState
	collisionService *collision.CollisionService
}

func NewPhysicsEngine(state *GameState, collisionService *collision.CollisionService) *PhysicsEngine {
	return &PhysicsEngine{
		state:            state,
		collisionService: collisionService,
	}
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
		} else if bullet.Life <= 0 && bullet.Type == domain.PlayerSom {
			gs.collisionService.TriggerExplosion(bullet)
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
