package models

type ClientMessage struct {
	Player struct {
		Id string  `json:"id"`
		MX float32 `json:"mx"`
		MY float32 `json:"my"`
		A  float32 `json:"a"`
		S  bool    `json:"s"`
	} `json:"player"`
}

type ServerMessage struct {
	Type    string        `json:"type"`
	Players []PlayerState `json:"players"`
	Bullets []Bullet      `json:"bullets"`
}
