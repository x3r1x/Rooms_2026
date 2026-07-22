package lobby

import (
	"gamedevRooms/internal/game"
	"gamedevRooms/internal/model"
	"log"
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

func (l *Lobby) SetUserReadyState(userId string, readyState bool) {
	l.players[userId].Ready = readyState
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
	return l.playersReady == len(l.players) && len(l.players) >= model.MinCountOfPlayers
}

func (l *Lobby) RemoveUser(id string) {
	player, exists := l.players[id]
	if !exists {
		return
	}

	if player.Ready {
		l.playersReady--
	}

	delete(l.players, id)
	log.Printf("Игрок %s покинул лобби. Осталось: %d", id, len(l.players))
}

func (l *Lobby) StartGame() {
	if l.state == model.OngoingGameState || len(l.players) == 0 {
		return
	}

	l.gameLoop = game.NewGameLoop(game.NewGameState())
	go l.gameLoop.Run()

	for _, player := range l.players {
		if player.Ready {
			l.gameLoop.RegisterPlayer(&model.PlayerGameState{
				Id:          player.Id,
				Health:      model.MaxPlayerHealth,
				X:           model.PlayerSpawnPointX,
				Y:           model.PlayerSpawnPointY,
				Angle:       model.InitDirection,
				MoveX:       model.InitValue,
				MoveY:       model.InitValue,
				Connection:  player.Connection,
				ShootTimer:  0,
				RebornTimer: 0,
				BodyCount:   0,
				DeathCount:  0,
			})
		}
	}
	for id := range l.players {
		delete(l.players, id)
	}
	l.playersReady = 0
}
