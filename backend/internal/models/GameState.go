package models

type GameState struct {
	players      map[string]PlayerState
	bullets      []Bullet
	UpdateChan   chan ClientMessage
	RegisterChan chan PlayerState
	DeleteChan   chan string
	tickCount    uint64
}

func NewGameState() *GameState {
	return &GameState{
		players:      make(map[string]PlayerState),
		bullets:      make([]Bullet, 0),
		UpdateChan:   make(chan ClientMessage, 100),
		RegisterChan: make(chan PlayerState, 100),
		DeleteChan:   make(chan string, 100),
	}
}

// Getters

//func (gs *GameState) GetUpdateChan() chan ClientMessage {
//	return gs.UpdateChan
//}
//
//func (gs *GameState) GetRegisterChan() chan PlayerState {
//	return gs.RegisterChan
//}
//
//func (gs *GameState) GetDeleteChan() chan string {
//	return gs.DeleteChan
//}

func (gs *GameState) GetAllBullets() []Bullet {
	return gs.bullets
}

func (gs *GameState) GetAllPlayers() map[string]PlayerState {
	return gs.players
}

func (gs *GameState) GetPlayer(id string) (PlayerState, bool) {
	player, exist := gs.players[id]
	return player, exist
}

// Setters
func (gs *GameState) SetBullets(bullets []Bullet) {
	gs.bullets = bullets
}

// Utils
func (gs *GameState) IncrementTick() {
	gs.tickCount++
}

func (gs *GameState) AddPlayer(player PlayerState) {
	gs.players[player.Id] = player
}

func (gs *GameState) RemovePlayer(playerId string) {
	delete(gs.players, playerId)
}

func (gs *GameState) AddBullet(bullet Bullet) {
	gs.bullets = append(gs.bullets, bullet)
}
