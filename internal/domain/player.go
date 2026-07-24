package domain

type PlayerGameState struct {
	Id          string  `json:"id"`
	Nickname    string  `json:"-"`
	Health      float64 `json:"h"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Angle       float64 `json:"a"`
	MoveX       float64 `json:"mx"`
	MoveY       float64 `json:"my"`
	ShootTimer  int     `json:"-"`
	RebornTimer int     `json:"rt"`
	BodyCount   int     `json:"-"`
	DeathCount  int     `json:"-"`
}

type PlayerFinalState struct {
	Nickname string `json:"n"`
	Id       string `json:"id"`
	Kills    int    `json:"k"`
	Deaths   int    `json:"d"`
}
