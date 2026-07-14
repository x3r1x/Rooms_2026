package models

type Bullet struct {
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Direction float64 `json:"direction"`
	Life      float64 `json:"life"`
}

type BulletInterpolation struct {
	DX float64 `json:"dx"`
	DY float64 `json:"dy"`
}

type NewBullet struct {
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Direction float64 `json:"direction"`
}
