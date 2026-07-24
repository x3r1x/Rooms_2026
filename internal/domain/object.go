package domain

type Object struct {
	Id      string  `json:"id"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Width   float64 `json:"width"`
	Height  float64 `json:"height"`
	Type    string  `json:"type"`
	IsSolid bool    `json:"isSolid"`
}

func GetObjectPoints(x, y, width, height float64) []Point {
	return []Point{
		{x, y},
		{x + width, y},
		{x + width, y + height},
		{x, y + height},
	}
}
