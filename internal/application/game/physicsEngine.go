package game

import (
	"gamedevRooms/internal/adapters/collision"
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/recovery"
	"gamedevRooms/internal/state"
	"math"
)

type PhysicsEngine struct {
	gameState        *state.GameState
	collisionService *collision.CollisionService
}

func NewPhysicsEngine(state *state.GameState, collisionService *collision.CollisionService) *PhysicsEngine {
	return &PhysicsEngine{
		gameState:        state,
		collisionService: collisionService,
	}
}

func (pe *PhysicsEngine) updateShooterTimers() {
	for _, player := range pe.gameState.GetAllPlayers() {
		if player.CooldownTimer > 0 {
			player.CooldownTimer--
		}
	}
}

func (pe *PhysicsEngine) updateBullets() {
	defer recovery.Recover()
	activeBullets := make([]domain.Bullet, 0)
	for _, bullet := range pe.gameState.GetAllBullets() {
		if bullet.Life > 0 {
			bullet.Move()
			bullet.Life--
			hit, player, obj := pe.collisionService.CheckBulletCollision(bullet)
			if hit {
				if player != nil && player.Health > 0 {
					pe.collisionService.HandlePlayerHit(player, bullet)
				} else if obj != nil || bullet.Type == domain.PlayerSom {
					if bullet.Type == domain.PlayerSom {
						pe.collisionService.TriggerExplosion(bullet)
					}
				}
				continue
			}
			activeBullets = append(activeBullets, bullet)
		} else if bullet.Life <= 0 && bullet.Type == domain.PlayerSom {
			pe.collisionService.TriggerExplosion(bullet)
		}
	}
	pe.gameState.SetBullets(activeBullets)
}

func (pe *PhysicsEngine) updatePlayers() {
	defer recovery.Recover()
	for _, player := range pe.gameState.GetAllPlayers() {
		if player.Health <= 0 {
			if player.RebornTimer > 0 {
				player.RebornTimer--
			} else if player.RebornTimer == 0 {
				player.Health = domain.MaxPlayerHealth
				player.PlayerShield = true
				player.CooldownTimer = domain.PlayerShieldTimer
				player.PlayerShieldTimer = domain.PlayerShieldTimer
			}
			continue
		}

		if player.PlayerShield && player.PlayerShieldTimer > 0 {
			player.PlayerShieldTimer--
			if player.PlayerShieldTimer == 0 {
				player.PlayerShield = false
			}
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
		if hit, _ := pe.collisionService.CheckPlayerObjectCollision(player); hit {
			player.X -= deltaX
		}

		player.Y += deltaY
		if hit, _ := pe.collisionService.CheckPlayerObjectCollision(player); hit {
			player.Y -= deltaY
		}

		if hit, _ := pe.collisionService.CheckPlayerObjectCollision(player); hit {
			pe.collisionService.ResolvePlayerCollisionSmooth(player)
		}

		if hit, direction, targetRoomId := pe.collisionService.CheckPlayerExitCollision(player); hit {
			pe.handleRoomTransition(player, direction, targetRoomId)
		}
	}
}

func (pe *PhysicsEngine) handleRoomTransition(player *domain.PlayerGameState, direction, targetRoomId string) {
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
	pe.gameState.SetPlayerRoom(player.Id, targetRoomId)
}
