package game

import (
	"gamedevRooms/internal/collision"
	"gamedevRooms/internal/model"
	"log"
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

		playerSAT := cs.buildPlayerSAT(player)
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
	playerSAT := cs.buildPlayerSAT(player)

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

func (cs *CollisionService) buildPlayerSAT(player *model.PlayerGameState) collision.SATBox {
	points := collision.GetPlayerPoints(player.X, player.Y, player.Angle)
	normals := collision.GetNormals(points)
	return collision.SATBox{
		Points:  points,
		Normals: normals,
	}
}

func (cs *CollisionService) buildObjectSAT(obj *model.Object) collision.SATBox {
	points := collision.GetObjectPoints(obj.X, obj.Y, obj.Width, obj.Height)
	normals := collision.GetNormals(points)
	return collision.SATBox{
		Points:  points,
		Normals: normals,
	}
}

//func (cs *CollisionService) HandleObjectHit(obj *model.Object, bullet model.Bullet) {
//	if obj == nil {
//		return
//	}
//	if !obj.IsDestroyable {
//		return
//	}
//	obj.Health -= cs.calculateDamage(bullet)
//
//	log.Printf("HIT! Object %s took %.2f damage. Health: %.2f\n")
//	if obj.Health < 0 {
//		cs.state.RemoveObject(obj.Id)
//		log.Printf("Object %s was destroyed", obj.Id)
//	}
//}

func (cs *CollisionService) HandlePlayerHit(player *model.PlayerGameState, bullet model.Bullet) {
	if player == nil {
		return
	}
	player.Health -= cs.calculateDamage(bullet)

	log.Printf("HIT! Player %s took %.2f damage. Health: %.2f\n",
		player.Id, player.Health)

	if player.Health < 0 {
		if killer, exist := cs.state.GetPlayer(bullet.OwnerId); exist {
			killer.BodyCount++
		}
		player.DeathCount++
		player.RebornTimer = model.PlayerRebornTimer
	}
}

func (cs *CollisionService) calculateDamage(bullet model.Bullet) float64 {
	return model.BulletDamage * (bullet.Life/model.BulletLife*model.BulletDamageMulti + 1)
}

func (cs *CollisionService) CheckPlayerExitCollision(player *model.PlayerGameState) (bool, string, string) {
	// Получаем текущую комнату игрока
	currentRoomId := cs.state.GetPlayerRoom(player.Id)
	if currentRoomId == "" {
		return false, "", ""
	}

	// Получаем информацию о комнате
	roomInfo := cs.state.GetRoomManager().GetRoomInfo(currentRoomId)
	if roomInfo == nil {
		return false, "", ""
	}

	// Проверяем каждую сторону выхода
	exitChecks := map[string]func() bool{
		model.TopMarker:    func() bool { return player.Y < 0 },
		model.BottomMarker: func() bool { return player.Y > float64(model.RoomHeight*36) },
		model.LeftMarker:   func() bool { return player.X < 0 },
		model.RightMarker:  func() bool { return player.X > float64(model.RoomWidth*36) },
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
	return float64(model.RoomWidth * 36)
}
