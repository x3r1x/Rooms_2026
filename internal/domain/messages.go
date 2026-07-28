package domain

type ClientRegisterMessage struct {
	Nickname string `json:"n"`
}

type ClientReadyStateMessage struct {
	Ready       bool   `json:"r"`
	PlayerClass string `json:"pc"`
}

type ClientGameMessage struct {
	Id      string  `json:"id"`
	MX      float64 `json:"mx"`
	MY      float64 `json:"my"`
	Angle   float64 `json:"a"`
	IsShoot bool    `json:"s"`
}

type LobbyPlayerMessage struct {
	Nickname string `json:"n"`
	Id       string `json:"id"`
	Ready    bool   `json:"r"`
}

type ServerLobbyMessage struct {
	State   string               `json:"s"`
	OwnId   string               `json:"oId"`
	Players []LobbyPlayerMessage `json:"p"`
}

type RoomMessage struct {
	ExitTop    string `json:"eT"`
	ExitLeft   string `json:"eL"`
	ExitBottom string `json:"eB"`
	ExitRight  string `json:"eR"`
	Id         string `json:"id"`
	BorderType int    `json:"bT"`
}

type ServerReadyMessage struct {
	State     string                 `json:"s"`
	Countdown float64                `json:"c"`
	Map       map[string]RoomMessage `json:"m,omitempty"`
}

type ServerGameMessage struct {
	State     string             `json:"s"`
	Time      float64            `json:"t"`
	Players   []*PlayerGameState `json:"p"`
	Bullets   []*Bullet          `json:"b"`
	Statistic []PlayerStatistic  `json:"stat"`
}

type ServerCountdownMessage struct {
	State     string `json:"s"`
	Countdown int    `json:"c"`
}

type ServerEndMessage struct {
	State  string             `json:"s"`
	Result []PlayerFinalState `json:"r"`
}
