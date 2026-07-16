package models

type Bullet struct {
	Id        string  `json:"-"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Direction float64 `json:"direction"`
	Life      float64 `json:"life"`
	OwnerId   string  `json:"ownerid"`
}
