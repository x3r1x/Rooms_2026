package models

type GameState struct {
	Players      map[string]*PlayerState
	Bullets      []Bullet
	InputChan    chan ClientMessage
	RegisterChan chan PlayerState
	LeaveChan    chan string
	TickCount    uint64
}

//func NewGameState() *GameState {
//	return &GameState{
//		Players:      make(map[string]*PlayerState),
//		InputChan:    make(chan ClientMessage),
//		RegisterChan: make(chan PlayerState),
//		LeaveChan:    make(chan string),
//	}
//}

var Game = &GameState{
	Players:      make(map[string]*PlayerState),
	Bullets:      make([]Bullet),
	RegisterChan: make(chan PlayerState),
	InputChan:    make(chan ClientMessage, 100),
	LeaveChan:    make(chan string, 100),
}
