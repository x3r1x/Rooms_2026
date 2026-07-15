package game

import (
	"fmt"
	"gamedevRooms/internal/models"
	"math"
	"sync/atomic"
)

var (
	Clients     = make(map[string]*models.PlayerState)
	CommandChan = make(chan Command, 1000)
)

const (
	MAX_BULLET_SPEED = 22.5
	MAP_SIZE         = 2000.0
	BULLET_LIFE      = 60
	PLAYER_SPEED     = 20
)

type CommandType int

const (
	RegisterPlayer CommandType = iota
	UpdatePlayer
	DisconnectPlayer
)

type Command struct {
	Type           CommandType
	ID             string
	Player         *models.PlayerState
	DirX, DirY     float64
	DeltaVisualDir float64
	NewBulletsDir  []float64
}

var bulletCounter uint64

func nextBulletId(playerId string) string {
	n := atomic.AddUint64(&bulletCounter, 1)
	return fmt.Sprintf("%s_b_%d", playerId, n)
}

func ProcessCommands() {
	for len(CommandChan) > 0 {
		cmd := <-CommandChan
		switch cmd.Type {
		case RegisterPlayer:
			Clients[cmd.ID] = cmd.Player
		case UpdatePlayer:
			if p, ok := Clients[cmd.ID]; ok {
				p.Direction = math.Atan2(cmd.DirY, cmd.DirX)
				p.Direction += cmd.DeltaVisualDir

				for _, angle := range cmd.NewBulletsDir {
					id := nextBulletId(p.Id)
					p.Bullets[id] = models.Bullet{
						Id:        id,
						X:         p.X,
						Y:         p.Y,
						Direction: angle,
						StartX:    p.X,
						StartY:    p.Y,
						Owner:     p.Id,
						Life:      BULLET_LIFE,
					}
				}
			}
		case DisconnectPlayer:
			if p, ok := Clients[cmd.ID]; ok {
				delete(Clients, cmd.ID)
				close(p.SendChan)
			}
		}
	}
}
