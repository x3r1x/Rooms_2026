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

				existingID := make(map[string]struct{}, len(cmd.Bullets))
				for _, b := range p.Bullets {
					existingID[b.Id] = struct{}{}
				}
				for _, cb := range cmd.Bullets {
					if _, exist := existingID[cb.Id]; !exist {
						newBullet := models.Bullet{
							Id:        cb.Id,
							X:         cmd.X,
							Y:         cmd.Y,
							Direction: cb.Direction,
							StartX:    cmd.X,
							StartY:    cmd.Y,
							Owner:     cb.Owner,
							Life:      BULLET_LIFE,
						}
						p.Bullets = append(p.Bullets, newBullet)
						existingID[cb.Id] = struct{}{}
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
