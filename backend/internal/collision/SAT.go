package collision

type Point struct {
	X, Y float64
}

type Vector struct {
	X, Y float64
}

type Rect struct {
	X, Y, W, H float64
}

type Polygon struct {
	Points  []Point
	Normals []Vector
}

type Projection struct {
	Min, Max float64
}
