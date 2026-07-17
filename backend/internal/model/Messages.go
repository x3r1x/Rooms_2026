package model

type ClientMessage struct {
	Id string  `json:"id"`
	MX float64 `json:"mx"`
	MY float64 `json:"my"`
	A  float64 `json:"a"`
	S  bool    `json:"s"`
}

type ServerMessage struct {
	Type    string        `json:"type"`
	Players []PlayerState `json:"p"`
	Bullets []Bullet      `json:"b"`
}
