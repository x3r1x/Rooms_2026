package domain

type Object struct {
	Id      string  `json:"id"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Width   float64 `json:"width"`
	Height  float64 `json:"height"`
	Type    string  `json:"type"`
	IsSolid bool    `json:"isSolid"`
	RoomId  string  `json:"roomId"`
}

func (o *Object) GetPoints() []Point {
	return []Point{
		{o.X, o.Y},
		{o.X + o.Width, o.Y},
		{o.X + o.Width, o.Y + o.Height},
		{o.X, o.Y + o.Height},
	}
}

func (o *Object) GetBounds() (minX, minY, maxX, maxY float64) {
	return o.X, o.Y, o.X + o.Width, o.Y + o.Height
}

func (o *Object) ContainsPoint(x, y float64) bool {
	return x >= o.X && x <= o.X+o.Width &&
		y >= o.Y && y <= o.Y+o.Height
}

func (o *Object) Intersects(other *Object) bool {
	return o.X < other.X+other.Width &&
		o.X+o.Width > other.X &&
		o.Y < other.Y+other.Height &&
		o.Y+o.Height > other.Y
}
