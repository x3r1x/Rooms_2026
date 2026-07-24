package game

import (
	"gamedevRooms/internal/collision"
	"gamedevRooms/internal/model"
	"log"
	"math"
)

type CollisionService struct {
	state *GameState
}

func NewCollisionService(state *GameState) *CollisionService {
	return &CollisionService{state: state}
}

func (cs *CollisionService) CheckBulletCollision(bullet model.Bullet) (bool, *model.PlayerGameState, *model.Object) {
	bulletSAT := cs.buildBulletSAT(bullet)

	for _, player := range cs.state.GetAllPlayers() {
		if player.Id == bullet.OwnerId {
			continue
		}
		if player.Health <= 0 {
			continue
		}
		playerSAT := cs.buildPlayerSAT(player, true)
		if collision.CheckCollisionSAT(bulletSAT, playerSAT) {
			return true, player, nil
		}
	}

	for _, object := range cs.state.GetObjects() {
		if !object.IsSolid {
			continue
		}
		objectSAT := cs.buildObjectSAT(object)
		if collision.CheckCollisionSAT(bulletSAT, objectSAT) {
			return true, nil, object
		}
	}

	return false, nil, nil
}

func (cs *CollisionService) CheckPlayerObjectCollision(player *model.PlayerGameState) (bool, *model.Object) {
	playerSAT := cs.buildPlayerSAT(player, false)

	for _, object := range cs.state.GetObjects() {
		if !object.IsSolid {
			continue
		}
		objectSAT := cs.buildObjectSAT(object)
		if collision.CheckCollisionSAT(playerSAT, objectSAT) {
			return true, object
		}
	}
	return false, nil
}

func (cs *CollisionService) buildBulletSAT(bullet model.Bullet) collision.SATBox {
	points := collision.GetBulletPoints(bullet.X, bullet.Y, bullet.Direction)
	normals := collision.GetNormals(points)
	return collision.SATBox{
		Points:  points,
		Normals: normals,
	}
}

func (cs *CollisionService) buildPlayerSAT(player *model.PlayerGameState, rotated bool) collision.SATBox {
	var points []collision.Point
	var normals []collision.Point

	if rotated {
		points = collision.GetPlayerPoints(player.X, player.Y, player.Angle)
		normals = collision.GetNormals(points)
	} else {
		halfSize := model.PlayerHalfSize
		points = collision.GetAxisAlignedPoints(player.X, player.Y, halfSize)
		normals = collision.GetRectNormals()
	}

	return collision.SATBox{
		Points:  points,
		Normals: normals,
	}
}

func (cs *CollisionService) buildObjectSAT(obj *model.Object) collision.SATBox {
	points := collision.GetObjectPoints(obj.X, obj.Y, obj.Width, obj.Height)
	normals := collision.GetRectNormals()
	return collision.SATBox{
		Points:  points,
		Normals: normals,
	}
}

func (cs *CollisionService) HandlePlayerHit(player *model.PlayerGameState, bullet model.Bullet) {
	if player == nil {
		return
	}
	player.Health -= cs.calculateDamage(bullet)
	log.Printf("HIT! Player %s took damage. Health: %.2f", player.Id, player.Health)

	//knockbackX := math.Cos(bullet.Direction) * cs.calculateKnockbackKoef(bullet)
	//knockbackY := math.Sin(bullet.Direction) * cs.calculateKnockbackKoef(bullet)
	//player.X += knockbackX
	//player.Y += knockbackY

	//if hit, _ := cs.CheckPlayerObjectCollision(player); hit {
	//	cs.ResolvePlayerCollisionSmooth(player)
	//}

	if player.Health < 0 {
		if killer, exist := cs.state.GetPlayer(bullet.OwnerId); exist {
			killer.BodyCount++
		}
		player.DeathCount++
		player.RebornTimer = model.PlayerRebornTimer
	}
}

func (cs *CollisionService) calculateDamage(bullet model.Bullet) float64 {
	return math.Round(model.BulletDamage * (bullet.Life/model.BulletLife*model.BulletDamageMulti + 1))
}

//func (cs *CollisionService) calculateKnockbackKoef(bullet model.Bullet) float64 {
//	return model.KnockbackForce * (bullet.Life/model.BulletLife*model.BulletDamageMulti + 1)
//}

func (cs *CollisionService) CheckPlayerExitCollision(player *model.PlayerGameState) (bool, string, string) {
	currentRoomId := cs.state.GetPlayerRoom(player.Id)
	if currentRoomId == "" {
		return false, "", ""
	}

	roomInfo := cs.state.GetRoomManager().GetRoomInfo(currentRoomId)
	if roomInfo == nil {
		return false, "", ""
	}

	roomPixelWidth := float64(model.RoomWidth * int(model.TileSize))
	roomPixelHeight := float64(model.RoomHeight * int(model.TileSize))
	halfSize := model.PlayerHalfSize

	exitChecks := map[string]func() bool{
		model.TopMarker:    func() bool { return player.Y-halfSize < 0 },
		model.BottomMarker: func() bool { return player.Y+halfSize > roomPixelHeight },
		model.LeftMarker:   func() bool { return player.X-halfSize < 0 },
		model.RightMarker:  func() bool { return player.X+halfSize > roomPixelWidth },
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
	return float64(model.RoomWidth * int(model.TileSize))
}

func (cs *CollisionService) ResolvePlayerCollisionSmooth(player *model.PlayerGameState) bool {
	hit, obj := cs.CheckPlayerObjectCollision(player)
	if !hit {
		return false
	}

	playerHalf := model.PlayerHalfSize

	rightOverlap := (player.X + playerHalf) - obj.X
	leftOverlap := (obj.X + obj.Width) - (player.X - playerHalf)

	bottomOverlap := (player.Y + playerHalf) - obj.Y
	topOverlap := (obj.Y + obj.Height) - (player.Y - playerHalf)

	minOverlapX := math.Min(rightOverlap, leftOverlap)
	minOverlapY := math.Min(bottomOverlap, topOverlap)

	if minOverlapX < minOverlapY {
		if rightOverlap < leftOverlap {
			player.X = obj.X - playerHalf - model.Epsilon
		} else {
			player.X = obj.X + obj.Width + playerHalf + model.Epsilon
		}
	} else {
		if bottomOverlap < topOverlap {
			player.Y = obj.Y - playerHalf - model.Epsilon
		} else {
			player.Y = obj.Y + obj.Height + playerHalf + model.Epsilon
		}
	}

	return true
}
