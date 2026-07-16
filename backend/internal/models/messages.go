package models

type ClientMessage struct {
	Player struct {
		Id        string   `json:"id"`
		X         float64  `json:"x"`
		Y         float64  `json:"y"`
		Direction float64  `json:"direction"`
		Bullets   []Bullet `json:"bullets"`
	} `json:"player"`
}

type ServerMessage struct {
	Players []PlayerState `json:"players"`
}
