package models

import "encoding/json"

type ClientMessage struct {
	PlayerInterpolation struct {
		Direction            Vector2   `json:"direction"`
		DeltaVisualDirection float64   `json:"deltaDirection"`
		ID                   string    `json:"id"`
		NewBulletsDirection  []float64 `json:"newBulletsDirection"`
	}
}

type Vector2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ServerMessage struct {
	Type    string           `json:"type"`
	Players []PlayerSnapshot `json:"players"`
}

func MarshallJson(message *ServerMessage) ([]byte, error) {
	playerMap := make(map[string]PlayerSnapshot, len(message.Players))
	for _, player := range message.Players {
		playerMap[player.Id] = player
	}

	alias := struct {
		Type    string                    `json:"type"`
		Players map[string]PlayerSnapshot `json:"players"`
	}{
		Type:    message.Type,
		Players: playerMap,
	}
	return json.Marshal(alias)
}
