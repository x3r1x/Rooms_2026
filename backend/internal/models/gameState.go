package models

type GameState struct {
	Players       map[string]*PlayerState
	InputChan     chan ClientMessage
	RegisterChan  chan PlayerState
	LeaveChan     chan string
	BroadcastChan chan []PlayerState
	TickCount     uint64
}

func NewGameState() *GameState {
	return &GameState{
		Players:      make(map[string]*PlayerState),
		InputChan:    make(chan ClientMessage),
		RegisterChan: make(chan PlayerState),
		LeaveChan:    make(chan string),
		//BroadcastChan: make(chan []PlayerState),
	}
}

var Game = &GameState{
	Players:      make(map[string]*PlayerState),
	RegisterChan: make(chan PlayerState),
	InputChan:    make(chan ClientMessage, 100),
	LeaveChan:    make(chan string, 100),
	//BroadcastChan: make(chan []PlayerState, 1),
}
