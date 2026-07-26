package domain

import "math"

type Bullet struct {
	Id        string  `json:"id"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Direction float64 `json:"a"`
	Life      float64 `json:"life"`
	OwnerId   string  `json:"oId"`
	Speed     float64 `json:"-"`
	Damage    float64 `json:"-"`
	Type      string  `json:"-"`
}

func (b *Bullet) Move() {
	b.X += math.Cos(b.Direction) * b.Speed
	b.Y += math.Sin(b.Direction) * b.Speed
}

func (b *Bullet) GetPoints() []Point {
	cosAngle := math.Cos(b.Direction)
	sinAngle := math.Sin(b.Direction)
	halfW := BulletWidth / 2.0
	halfH := BulletLength / 2.0

	localPoints := []Point{
		{-halfW, -halfH},
		{halfW, -halfH},
		{halfW, halfH},
		{-halfW, halfH},
	}

	points := make([]Point, 4)
	for i, lp := range localPoints {
		points[i] = Point{
			X: b.X + (lp.X*cosAngle - lp.Y*sinAngle),
			Y: b.Y + (lp.X*sinAngle + lp.Y*cosAngle),
		}
	}
	return points
}
