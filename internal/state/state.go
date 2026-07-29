package state

import (
	"fmt"
	_map "gamedevRooms/internal/adapters/map"
	"gamedevRooms/internal/domain"
	"log"
	"time"
)

type GameState struct {
	players       map[string]*domain.PlayerGameState
	objects       map[string]*domain.Object
	bullets       []domain.Bullet
	tickCount     uint64
	isGameActive  bool
	gameStartTime time.Time
	gameDuration  time.Duration
	kills         []*domain.Kill
	playerRooms   map[string]string
	roomPlayers   map[string][]string
	roomManager   *_map.MapManager
}

func NewGameState() *GameState {
	return &GameState{
		players:      make(map[string]*domain.PlayerGameState),
		objects:      make(map[string]*domain.Object),
		bullets:      make([]domain.Bullet, 0),
		gameDuration: domain.GameDuration * time.Second,
		playerRooms:  make(map[string]string),
		kills:        make([]*domain.Kill, 0),
		roomPlayers:  make(map[string][]string),
	}
}

func (gs *GameState) GetPlayerRoom(id string) string {
	if roomId, exists := gs.playerRooms[id]; exists {
		return roomId
	}
	return ""
}

func (gs *GameState) SetPlayerRoom(id, roomId string) {
	if oldRoom, exists := gs.playerRooms[id]; exists {
		gs.removePlayerFromRoom(oldRoom, id)
	}

	if player, exists := gs.players[id]; exists {
		player.RoomId = roomId
	}
	gs.playerRooms[id] = roomId
	gs.addPlayerToRoom(roomId, id)
}

func (gs *GameState) GetGameStartTime() time.Time {
	return gs.gameStartTime
}

func (gs *GameState) GetRoomManager() *_map.MapManager {
	return gs.roomManager
}

func (gs *GameState) SetRoomManager(mm *_map.MapManager) {
	gs.roomManager = mm
}

func (gs *GameState) addPlayerToRoom(roomId, playerId string) {
	gs.roomPlayers[roomId] = append(gs.roomPlayers[roomId], playerId)
}

func (gs *GameState) removePlayerFromRoom(roomId, playerId string) {
	players := gs.roomPlayers[roomId]
	for i, id := range players {
		if id == playerId {
			gs.roomPlayers[roomId] = append(players[:i], players[i+1:]...)
			break
		}
	}
}

func (gs *GameState) GetObjects() map[string]*domain.Object {
	return gs.objects
}

func (gs *GameState) SetObjects(objects map[string]*domain.Object) {
	gs.objects = objects
}

func (gs *GameState) AddObject(obj *domain.Object) {
	gs.objects[obj.Id] = obj
}

func (gs *GameState) RemoveObject(id string) {
	if _, exist := gs.objects[id]; exist {
		delete(gs.objects, id)
	}
}

func (gs *GameState) IsGameActive() bool {
	return gs.isGameActive
}

func (gs *GameState) SetGameActive(active bool) {
	gs.isGameActive = active
	if active {
		gs.gameStartTime = time.Now()
		log.Println("=== Game is active ===")
	}
}

func (gs *GameState) GetRemainingSeconds() int {
	if !gs.isGameActive {
		return int(gs.gameDuration.Seconds())
	}
	elapsed := time.Since(gs.gameStartTime)
	if elapsed >= gs.gameDuration {
		return 0
	}
	return int((gs.gameDuration - elapsed).Seconds())
}

// ==================

func (gs *GameState) AddKill(killerId, victimId string) {
	gs.kills = append(gs.kills, domain.NewKill(killerId, victimId, float64(time.Since(gs.gameStartTime).Milliseconds())))
}

func (gs *GameState) GetKills() []*domain.Kill {
	return gs.kills
}

func (gs *GameState) ClearKills() {
	gs.kills = gs.kills[:0]
}

func (gs *GameState) GetAllBullets() []domain.Bullet {
	return gs.bullets
}

func (gs *GameState) GetAllPlayers() map[string]*domain.PlayerGameState {
	return gs.players
}

func (gs *GameState) GetPlayer(id string) (*domain.PlayerGameState, bool) {
	player, exist := gs.players[id]
	return player, exist
}

func (gs *GameState) SetBullets(bullets []domain.Bullet) {
	gs.bullets = bullets
}

func (gs *GameState) IncrementTick() {
	gs.tickCount++
}

func (gs *GameState) AddPlayer(player *domain.PlayerGameState) {
	log.Println("Register ", player.Id)
	gs.players[player.Id] = player
}

func (gs *GameState) RemovePlayer(playerId string) {
	fmt.Println("Delete ", playerId)
	delete(gs.players, playerId)
}

func (gs *GameState) GetCountOfPlayers() int {
	return len(gs.players)
}

func (gs *GameState) GetPlayersByRoom() map[string][]domain.PlayerGameState {
	rooms := make(map[string][]domain.PlayerGameState)
	for id, roomId := range gs.playerRooms {
		if player, exists := gs.players[id]; exists {
			rooms[roomId] = append(rooms[roomId], *player)
		}
	}
	return rooms
}

func (gs *GameState) GetBulletsByRoom() map[string][]domain.Bullet {
	rooms := make(map[string][]domain.Bullet)

	for id := range gs.roomPlayers {
		rooms[id] = gs.getBulletsInRoom(id)
	}

	return rooms
}

func (gs *GameState) getBulletsInRoom(roomId string) []domain.Bullet {
	bulletsInRoom := make([]domain.Bullet, 0)

	for _, bullet := range gs.bullets {
		if bullet.RoomId == roomId {
			bulletsInRoom = append(bulletsInRoom, bullet)
		}
	}

	return bulletsInRoom
}
