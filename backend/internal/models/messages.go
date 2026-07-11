package models

type ClientMessage struct {
	Player struct {
		X         float64  `json:"x"`
		Y         float64  `json:"y"`
		DIRECTION float64  `json:"direction"`
		ID        string   `json:"id"`
		Bullets   []Bullet `json:"bullets"`
	} `json:"player"`
}

type ServerMessage struct {
	Players []PlayerState `json:"players"`
}
