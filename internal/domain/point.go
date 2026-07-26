package domain

import "math"

type Point struct {
	X, Y float64
}

// Todo: убрать отюда
func GetPlayerPoints(centerX, centerY, angle float64) []Point {
	cosAngle := math.Cos(angle)
	sinAngle := math.Sin(angle)
	halfSize := PlayerHalfSize

	localPoints := []Point{
		{-halfSize, -halfSize},
		{halfSize, -halfSize},
		{halfSize, halfSize},
		{-halfSize, halfSize},
	}

	points := make([]Point, 4)
	for i, lp := range localPoints {
		points[i] = Point{
			X: centerX + (lp.X*cosAngle - lp.Y*sinAngle),
			Y: centerY + (lp.X*sinAngle + lp.Y*cosAngle),
		}
	}
	return points
}

func GetAxisAlignedPoints(centerX, centerY, halfSize float64) []Point {
	return []Point{
		{centerX - halfSize, centerY - halfSize},
		{centerX + halfSize, centerY - halfSize},
		{centerX + halfSize, centerY + halfSize},
		{centerX - halfSize, centerY + halfSize},
	}
}

func GetRectNormals() []Point {
	return []Point{
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
	}
}

func GetNormals(points []Point) []Point {
	normals := make([]Point, len(points))
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]
		edge := Point{X: p2.X - p1.X, Y: p2.Y - p1.Y}
		normals[i] = Point{X: -edge.Y, Y: edge.X}
	}
	return normals
}
