package Lobby

import (
	"gamedevRooms/internal/game"
	"gamedevRooms/internal/model"
	"time"
)

type Lobby struct {
	state        string
	players      map[string]*LobbyPlayer
	playersReady int
	gameLoop     *game.GameLoop
}

func NewLobby() *Lobby {
	return &Lobby{
		players: make(map[string]*LobbyPlayer),
	}
}

func (l *Lobby) GetGameLoop() *game.GameLoop {
	return l.gameLoop
}

func (l *Lobby) GetState() string {
	return l.state
}

func (l *Lobby) SetUserReadyState(userId *string, readyState bool) {
	l.players[*userId].Ready = readyState
	l.playersReady++
}

func (l *Lobby) SetState(newState string) {
	l.state = newState
}

func (l *Lobby) AddUser(nickname string) string {
	newPlayer := NewLobbyPlayer(nickname)
	l.players[newPlayer.Id] = newPlayer

	return newPlayer.Id
}

func (l *Lobby) CheckIfEveryoneReady() bool {
	return l.playersReady == len(l.players)
}

func (l *Lobby) RemoveUser(id string) {
	delete(l.players, id)
}

func (l *Lobby) Run() {
	ticker := time.NewTicker(model.TickTime * time.Millisecond)
	defer ticker.Stop()

	//for {
	//
	//}
}

func (l *Lobby) StartGameLoop() {
	gameState := game.NewGameState()
	l.gameLoop = game.NewGameLoop(gameState)
	go l.gameLoop.Run()
}
