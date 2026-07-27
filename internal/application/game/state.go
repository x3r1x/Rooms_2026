package game

import (
	"fmt"
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
	playerRooms   map[string]string
	roomPlayers   map[string][]string
	roomManager   interface{}
}

func NewGameState() *GameState {
	return &GameState{
		players:      make(map[string]*domain.PlayerGameState),
		objects:      make(map[string]*domain.Object),
		bullets:      make([]domain.Bullet, 0),
		gameDuration: domain.GameDuration * time.Second,
		playerRooms:  make(map[string]string),
		roomPlayers:  make(map[string][]string),
	}
}

// ====ROOM====
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

func (gs *GameState) GetRoomManager() interface{} {
	return gs.roomManager
}

func (gs *GameState) SetRoomManager(mm interface{}) {
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

// ==================
// ===== OBJECT =====
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

// ==================

// === LOBBITOMIA ===

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

func (gs *GameState) UpdatePlayer(upd domain.ClientGameMessage) {
	player, exist := gs.GetPlayer(upd.Id)
	if !exist || player == nil {
		return
	}

	player.Angle = upd.Angle
	player.MoveX = upd.MX
	player.MoveY = upd.MY

	if gs.isGameActive && player.Health > 0 && upd.IsShoot && player.CooldownTimer <= 0 {
		gs.AddBullet(player)
		player.CooldownTimer = domain.ShootCooldown
	}
	if player.Health < 0 && player.RebornTimer != 0 {
		player.RebornTimer--
	} else if player.Health < 0 && player.RebornTimer == 0 {
		player.Health = domain.MaxPlayerHealth
	}
}

func (gs *GameState) AddBullet(player *domain.PlayerGameState) {
	bullets := player.Weapon(player)
	gs.bullets = append(gs.bullets, bullets...)
}

func (gs *GameState) GetCountOfPlayers() int {
	return len(gs.players)
}

func (gs *GameState) GetPlayersByRoom() map[string][]*domain.PlayerGameState {
	rooms := make(map[string][]*domain.PlayerGameState)
	for id, roomId := range gs.playerRooms {
		if player, exists := gs.players[id]; exists {
			rooms[roomId] = append(rooms[roomId], player)
		}
	}
	return rooms
}
