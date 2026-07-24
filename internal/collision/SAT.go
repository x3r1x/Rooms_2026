package collision

import (
	"gamedevRooms/internal/domain"
	"math"
)

type SATBox struct {
	Points  []Point
	Normals []Point
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
	return (max1+domain.Epsilon >= min2) && (max2+domain.Epsilon >= min1)
}
