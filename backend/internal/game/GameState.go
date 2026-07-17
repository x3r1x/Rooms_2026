package game

import "gamedevRooms/internal/model"

type GameState struct {
	players      map[string]*model.PlayerState
	bullets      []model.Bullet
	UpdateChan   chan model.ClientMessage
	RegisterChan chan *model.PlayerState
	DeleteChan   chan string
	tickCount    uint64
}

func NewGameState() *GameState {
	return &GameState{
		players:      make(map[string]*model.PlayerState),
		bullets:      make([]model.Bullet, 0),
		UpdateChan:   make(chan model.ClientMessage, 100),
		RegisterChan: make(chan *model.PlayerState, 100),
		DeleteChan:   make(chan string, 100),
	}
}

func (gs *GameState) GetAllBullets() []model.Bullet {
	return gs.bullets
}

func (gs *GameState) GetAllPlayers() map[string]*model.PlayerState {
	return gs.players
}

func (gs *GameState) GetPlayer(id string) (*model.PlayerState, bool) {
	player, exist := gs.players[id]
	return player, exist
}

func (gs *GameState) SetBullets(bullets []model.Bullet) {
	gs.bullets = bullets
}

func (gs *GameState) IncrementTick() {
	gs.tickCount++
}

func (gs *GameState) AddPlayer(player *model.PlayerState) {
	gs.players[player.Id] = player
}

func (gs *GameState) RemovePlayer(playerId string) {
	delete(gs.players, playerId)
}

func (gs *GameState) AddBullet(bullet model.Bullet) {
	gs.bullets = append(gs.bullets, bullet)
}
