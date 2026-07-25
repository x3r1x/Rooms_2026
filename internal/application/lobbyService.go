package application

import (
	"gamedevRooms/internal/domain"
	"gamedevRooms/internal/game"
	"gamedevRooms/internal/ports"
)

type Lobby struct {
	state             string
	players           map[string]*domain.LobbyPlayer
	playersReady      int
	gameLoop          *game.GameLoop
	addChan           chan *domain.LobbyPlayer
	readyChan         chan *domain.LobbyPlayer
	removeChan        chan string
	getStateChan      chan chan string
	getGameLoopChan   chan chan *game.GameLoop
	gameFinishChan    chan bool
	gameStateProvider ports.GameStateProvider
	mapManager        ports.MapManager
}
