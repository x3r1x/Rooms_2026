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

func (cs *CollisionService) CheckBulletCollision(bullet model.Bullet) (bool, *model.PlayerGameState) {
	bulletSAT := cs.buildBulletSAT(bullet)

	for _, player := range cs.state.GetAllPlayers() {
		if player.Id == bullet.OwnerId {
			continue
		}

		playerSAT := cs.buildPlayerSAT(player)
		if collision.CheckCollisionSAT(bulletSAT, playerSAT) {
			return true, player
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

func (cs *CollisionService) HandleHit(player *model.PlayerGameState, bullet model.Bullet) {
	damage := model.BulletDamage * (bullet.Life/model.BulletLife*model.BulletDamageMulti + 1)
	player.Health -= damage

	log.Printf("HIT! Player %s took %.2f damage. Health: %.2f\n",
		player.Id, damage, player.Health)

	if player.Health < 0 {
		player.RebornTimer = model.PlayerRebornTimer
	}
}
