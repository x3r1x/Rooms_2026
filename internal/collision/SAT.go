package collision

import (
	"gamedevRooms/internal/model"
	"math"
)

type Point struct {
	X, Y float64
}

type SATBox struct {
	Points  []Point
	Normals []Point
}

// GetPlayerPoints рассчитывает точки относительно центра игрока
func GetPlayerPoints(centerX, centerY, angle float64) []Point {
	cosAngle := math.Cos(angle)
	sinAngle := math.Sin(angle)
	halfSize := model.PlayerHalfSize

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

// GetBulletPoints рассчитывает точки пули относительно её центра (x, y)
func GetBulletPoints(x, y, angle float64) []Point {
	cosAngle := math.Cos(angle)
	sinAngle := math.Sin(angle)
	halfW := model.BulletWidth / 2.0
	halfH := model.BulletLength / 2.0

	localPoints := []Point{
		{-halfW, -halfH},
		{halfW, -halfH},
		{halfW, halfH},
		{-halfW, halfH},
	}

	points := make([]Point, 4)
	for i, lp := range localPoints {
		points[i] = Point{
			X: x + (lp.X*cosAngle - lp.Y*sinAngle),
			Y: y + (lp.X*sinAngle + lp.Y*cosAngle),
		}
	}
	return points
}

// GetObjectPoints рассчитывает точки объекта (стены) от левого верхнего угла (x, y)
// Это соответствует логике фронтенда и данным из Tiled
func GetObjectPoints(x, y, width, height float64) []Point {
	return []Point{
		{x, y},
		{x + width, y},
		{x + width, y + height},
		{x, y + height},
	}
}

// GetRectNormals возвращает стандартные нормали для axis-aligned прямоугольника
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

func CheckCollisionSAT(box1, box2 SATBox) bool {
	if !checkAxes(box1.Points, box2.Points, box1.Normals) {
		return false
	}
	return checkAxes(box1.Points, box2.Points, box2.Normals)
}

func checkAxes(points1, points2 []Point, normals []Point) bool {
	for _, axis := range normals {
		min1, max1 := getMinMax(points1, axis)
		min2, max2 := getMinMax(points2, axis)
		if !isOverlapping(min1, max1, min2, max2) {
			return false
		}
	}
	return true
}

func getMinMax(points []Point, axis Point) (float64, float64) {
	minimum := math.Inf(1)
	maximum := math.Inf(-1)
	length := math.Sqrt(axis.X*axis.X + axis.Y*axis.Y)
	if length > 0 {
		axis.X /= length
		axis.Y /= length
	}
	for _, p := range points {
		dot := p.X*axis.X + p.Y*axis.Y
		if dot < minimum {
			minimum = dot
		}
		if dot > maximum {
			maximum = dot
		}
	}
	return minimum, maximum
}

func isOverlapping(min1, max1, min2, max2 float64) bool {
	return (max1 >= min2) && (max2 >= min1)
}
