package collision

import (
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/ports"
	"log"
	"math"
)

type CollisionService struct {
	state ports.GameStateProvider
}

func NewCollisionService(state ports.GameStateProvider) *CollisionService {
	return &CollisionService{state: state}
}

func (cs *CollisionService) CheckBulletCollision(bullet domain.Bullet) (bool, *domain.PlayerGameState, *domain.Object) {
	bulletSAT := cs.buildBulletSAT(bullet)

	_, exists := cs.state.GetPlayer(bullet.OwnerId)
	bulletRoomId := ""
	if exists {
		bulletRoomId = cs.state.GetPlayerRoom(bullet.OwnerId)
	}

	for _, player := range cs.state.GetAllPlayers() {
		if player.Id == bullet.OwnerId {
			continue
		}
		if player.Health <= 0 {
			continue
		}
		if cs.state.GetPlayerRoom(player.Id) != bulletRoomId {
			continue
		}
		playerSAT := cs.buildPlayerSAT(player, true)
		if CheckCollisionSAT(bulletSAT, playerSAT) {
			return true, player, nil
		}
	}

	for _, object := range cs.state.GetObjects() {
		if !object.IsSolid {
			continue
		}
		if object.RoomId != "" && object.RoomId != bulletRoomId {
			continue
		}
		objectSAT := cs.buildObjectSAT(object)
		if CheckCollisionSAT(bulletSAT, objectSAT) {
			return true, nil, object
		}
	}

	return false, nil, nil
}

func (cs *CollisionService) CheckPlayerObjectCollision(player *domain.PlayerGameState) (bool, *domain.Object) {
	playerSAT := cs.buildPlayerSAT(player, false)

	for _, object := range cs.state.GetObjects() {
		if !object.IsSolid {
			continue
		}
		if player.RoomId != object.RoomId {
			continue
		}
		objectSAT := cs.buildObjectSAT(object)
		if CheckCollisionSAT(playerSAT, objectSAT) {
			return true, object
		}
	}
	return false, nil
}

func (cs *CollisionService) buildBulletSAT(bullet domain.Bullet) SATBox {
	points := bullet.GetPoints()
	normals := domain.GetNormals(points)
	return SATBox{
		Points:  points,
		Normals: normals,
	}
}

func (cs *CollisionService) buildPlayerSAT(player *domain.PlayerGameState, rotated bool) SATBox {
	var points []domain.Point
	var normals []domain.Point

	if rotated {
		points = domain.GetPlayerPoints(player.X, player.Y, player.Angle)
		normals = domain.GetNormals(points)
	} else {
		halfSize := domain.PlayerHalfSize
		points = domain.GetAxisAlignedPoints(player.X, player.Y, halfSize)
		normals = domain.GetRectNormals()
	}

	return SATBox{
		Points:  points,
		Normals: normals,
	}
}

func (cs *CollisionService) buildObjectSAT(obj *domain.Object) SATBox {
	points := obj.GetPoints()
	normals := domain.GetRectNormals()
	return SATBox{
		Points:  points,
		Normals: normals,
	}
}

func (cs *CollisionService) HandlePlayerHit(player *domain.PlayerGameState, bullet domain.Bullet) {
	if player == nil {
		return
	}
	if bullet.Type == domain.PlayerSom {
		cs.handleExplosion(bullet)
		return
	}
	player.Health -= cs.calculateDamage(bullet)
	log.Printf("HIT! Player %s took damage. Health: %.2f", player.Id, player.Health)

	knockbackX := math.Cos(bullet.Direction) * cs.calculateKnockbackKoef(bullet)
	knockbackY := math.Sin(bullet.Direction) * cs.calculateKnockbackKoef(bullet)
	player.X += knockbackX
	player.Y += knockbackY

	if hit, _ := cs.CheckPlayerObjectCollision(player); hit {
		cs.ResolvePlayerCollisionSmooth(player)
	}

	if player.Health < 0 {
		if killer, exist := cs.state.GetPlayer(bullet.OwnerId); exist {
			killer.BodyCount++
			cs.state.AddKill(killer.Id, player.Id)
		}
		player.DeathCount++
		player.RebornTimer = domain.PlayerRebornTimer
	}
}

func (cs *CollisionService) TriggerExplosion(bullet domain.Bullet) {
	cs.handleExplosion(bullet)
}

func (cs *CollisionService) handleExplosion(bullet domain.Bullet) {
	bulletRoomId := ""
	if owner, exists := cs.state.GetPlayer(bullet.OwnerId); exists {
		bulletRoomId = cs.state.GetPlayerRoom(owner.Id)
	}
	for _, target := range cs.state.GetAllPlayers() {
		if target.Health <= 0 {
			continue
		}
		if cs.state.GetPlayerRoom(target.Id) != bulletRoomId {
			continue
		}

		dx := target.X - bullet.X
		dy := target.Y - bullet.Y
		distance := math.Sqrt(dx*dx + dy*dy)

		if distance < domain.ExplosionRadius {
			damageMultiplayer := (1 - distance/domain.ExplosionRadius) * domain.BulletDamageMulti
			target.Health -= damageMultiplayer * domain.BulletDamageSom
			if target.Health < 0 {
				if killer, exists := cs.state.GetPlayer(bullet.OwnerId); exists {
					if killer.Id != target.Id {
						killer.BodyCount++
					}
					cs.state.AddKill(killer.Id, target.Id)
				}
				target.DeathCount++
				target.RebornTimer = domain.PlayerRebornTimer
			}
		}
	}
}

func (cs *CollisionService) calculateDamage(bullet domain.Bullet) float64 {
	return math.Round(bullet.Damage * domain.BulletDamageMulti)
}

func (cs *CollisionService) calculateKnockbackKoef(bullet domain.Bullet) float64 {
	return domain.KnockbackForce * (bullet.Life/60*domain.BulletDamageMulti + 1)
}

func (cs *CollisionService) CheckPlayerExitCollision(player *domain.PlayerGameState) (bool, string, string) {
	currentRoomId := cs.state.GetPlayerRoom(player.Id)
	if currentRoomId == "" {
		return false, "", ""
	}

	mapManager, ok := cs.state.GetRoomManager().(ports.MapManager)
	if !ok || mapManager == nil {
		log.Printf("RoomManager is not MapManager or is nil")
		return false, "", ""
	}
	roomInfo := mapManager.GetRoomInfo(currentRoomId)
	if roomInfo == nil {
		return false, "", ""
	}

	roomPixelWidth := float64(domain.RoomWidth * int(domain.TileSize))
	roomPixelHeight := float64(domain.RoomHeight * int(domain.TileSize))
	halfSize := domain.PlayerHalfSize

	exitChecks := map[string]func() bool{
		domain.TopMarker:    func() bool { return player.Y-halfSize < 0 },
		domain.BottomMarker: func() bool { return player.Y+halfSize > roomPixelHeight },
		domain.LeftMarker:   func() bool { return player.X-halfSize < 0 },
		domain.RightMarker:  func() bool { return player.X+halfSize > roomPixelWidth },
	}

	for direction, check := range exitChecks {
		if check() {
			targetRoomId := roomInfo.GetExit(direction)
			if targetRoomId != "" {
				return true, direction, targetRoomId
			}
		}
	}
	return false, "", ""
}

func (cs *CollisionService) getRoomPixelSize() float64 {
	return float64(domain.RoomWidth * int(domain.TileSize))
}

func (cs *CollisionService) ResolvePlayerCollisionSmooth(player *domain.PlayerGameState) bool {
	hit, obj := cs.CheckPlayerObjectCollision(player)
	if !hit {
		return false
	}

	playerHalf := domain.PlayerHalfSize

	rightOverlap := (player.X + playerHalf) - obj.X
	leftOverlap := (obj.X + obj.Width) - (player.X - playerHalf)

	bottomOverlap := (player.Y + playerHalf) - obj.Y
	topOverlap := (obj.Y + obj.Height) - (player.Y - playerHalf)

	minOverlapX := math.Min(rightOverlap, leftOverlap)
	minOverlapY := math.Min(bottomOverlap, topOverlap)

	if minOverlapX < minOverlapY {
		if rightOverlap < leftOverlap {
			player.X = obj.X - playerHalf - domain.Epsilon
		} else {
			player.X = obj.X + obj.Width + playerHalf + domain.Epsilon
		}
	} else {
		if bottomOverlap < topOverlap {
			player.Y = obj.Y - playerHalf - domain.Epsilon
		} else {
			player.Y = obj.Y + obj.Height + playerHalf + domain.Epsilon
		}
	}

	return true
}
