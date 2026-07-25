package ports

import "gamedevRooms/internal/domain"

type CollisionService interface {
	CheckBulletCollision(bullet domain.Bullet) (bool, *domain.PlayerGameState, *domain.Object)
	CheckPlayerObjectCollision(player *domain.PlayerGameState) (bool, *domain.Object)
	HandlePlayerHit(player *domain.PlayerGameState, bullet domain.Bullet)
	ResolvePlayerCollisionSmooth(player *domain.PlayerGameState) bool
	//CheckPlayerExitCollision(player *domain.PlayerGameState) (bool, string, string)
}
