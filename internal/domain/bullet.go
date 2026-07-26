package domain

import "math"

type Bullet struct {
	Id        string  `json:"id"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Direction float64 `json:"a"`
	Life      float64 `json:"life"`
	OwnerId   string  `json:"oId"`
}

func (b *Bullet) IsActive() bool {
	return b.Life > 0
}

func (b *Bullet) Update() {
	if b.Life > 0 {
		b.Life--
		b.Move()
	}
}

func (b *Bullet) Move() {
	b.X += math.Cos(b.Direction) * MaxBulletSpeed
	b.Y += math.Sin(b.Direction) * MaxBulletSpeed
}

func (b *Bullet) CalculateDamage() float64 {
	lifeRatio := b.Life / BulletLife
	return BulletDamage * (lifeRatio*BulletDamageMulti + 1)
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
