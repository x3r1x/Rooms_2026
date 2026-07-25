package ports

import "gamedevRooms/internal/domain"

type BulletFactory interface {
	CreateBullet(player *domain.PlayerGameState) domain.Bullet
}
