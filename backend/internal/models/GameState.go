package models

type GameState struct {
	Players      map[string]*PlayerState
	Bullets      []Bullet
	InputChan    chan ClientMessage
	RegisterChan chan PlayerState
	LeaveChan    chan string
	TickCount    uint64
}

func NewGameState() *GameState {
	return &GameState{
		Players:      make(map[string]*PlayerState),
		Bullets:      make([]Bullet, 0),
		InputChan:    make(chan ClientMessage, 100),
		RegisterChan: make(chan PlayerState, 100),
		LeaveChan:    make(chan string, 100),
	}
}

// Getters

func (gs *GameState) GetInputChan() chan<- ClientMessage {
	return gs.InputChan
}

func (gs *GameState) GetRegisterChan() chan PlayerState {
	return gs.RegisterChan
}

func (gs *GameState) GetLeaveChan() chan string {
	return gs.LeaveChan
}

func (gs *GameState) GetAllBullets() []Bullet {
	return gs.Bullets
}

func (gs *GameState) GetAllPlayers() map[string]*PlayerState {
	return gs.Players
}

func (gs *GameState) GetPlayer(id string) (*PlayerState, bool) {
	player, exist := gs.Players[id]
	return player, exist
}

// Setters
func (gs *GameState) SetBullets(bullets *[]Bullet) {
	gs.Bullets = *bullets
}

// Utils
func (gs *GameState) IncrementTick() {
	gs.TickCount++
}

func (gs *GameState) AddPlayer(player *PlayerState) {
	gs.Players[player.Id] = player
}

func (gs *GameState) RemovePlayer(player *PlayerState) {
	delete(gs.Players, player.Id)
}
