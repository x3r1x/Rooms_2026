package models

type Bullet struct {
	Id        string  `json:"-"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Direction float64 `json:"direction"`
	StartX    float64 `json:"-"`
	StartY    float64 `json:"-"`
	Owner     string  `json:"-"`
	Life      float64 `json:"life"`
}
