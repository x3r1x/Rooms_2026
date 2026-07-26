package ports

import "gamedevRooms/internal/domain"

type GameStateProvider interface {
	GetObjects() map[string]*domain.Object
	SetObjects(objects map[string]*domain.Object)
	AddObject(obj *domain.Object)
	RemoveObject(id string)

	GetAllPlayers() map[string]*domain.PlayerGameState
	GetPlayer(id string) (*domain.PlayerGameState, bool)
	AddPlayer(player *domain.PlayerGameState)
	UpdatePlayer(upd domain.ClientGameMessage)
	RemovePlayer(playerId string)

	GetAllBullets() []domain.Bullet
	SetBullets(bullets []domain.Bullet)
	AddBullet(player *domain.PlayerGameState)

	IsGameActive() bool
	SetGameActive(active bool)
	GetRemainingSeconds() int
	IncrementTick()

	GetPlayerRoom(id string) string
	SetPlayerRoom(id, roomId string)
	GetRoomManager() interface{}
	SetRoomManager(roomManager interface{})
	GetCountOfPlayers() int
	GetPlayersByRoom() map[string][]*domain.PlayerGameState
}
