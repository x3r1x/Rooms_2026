package game

import (
	"gamedevRooms/internal/models"
)

var (
	Clients     = make(map[string]*models.PlayerState)
	CommandChan = make(chan Command, 1000)
)

const (
	MAX_BULLET_SPEED = 22.5
	MAP_SIZE         = 2000.0
	BULLET_LIFE      = 60
)

type CommandType int

const (
	RegisterPlayer CommandType = iota
	UpdatePlayer
	DisconnectPlayer
)

type Command struct {
	Type      CommandType
	ID        string
	Player    *models.PlayerState
	X, Y, Dir float64
	Bullets   []models.Bullet
}

func ProcessCommands() {
	for len(CommandChan) > 0 {
		cmd := <-CommandChan
		switch cmd.Type {
		case RegisterPlayer:
			Clients[cmd.ID] = cmd.Player
		case UpdatePlayer:
			if p, ok := Clients[cmd.ID]; ok {
				p.X = cmd.X
				p.Y = cmd.Y
				p.Direction = cmd.Dir

				for i := range cmd.Bullets {
					cb := &cmd.Bullets[i]
					if cb.Id == "" {
						continue
					}
					if _, exist := p.BulletIndex[cb.Id]; exist {
						continue
					}
					p.Bullets = append(p.Bullets, models.Bullet{
						Id:        cb.Id,
						X:         cb.X,
						Y:         cb.Y,
						Direction: cb.Direction,
						StartX:    cmd.X,
						StartY:    cmd.Y,
						Owner:     cb.Owner,
						Life:      BULLET_LIFE,
					})
					p.BulletIndex[cb.Id] = len(p.Bullets) - 1
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
