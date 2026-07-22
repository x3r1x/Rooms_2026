package model

type Object struct {
	Id      string  `json:"id"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Width   float64 `json:"width"`
	Height  float64 `json:"height"`
	Type    string  `json:"type"`
	IsSolid bool    `json:"isSolid"`
	//IsDestroyable bool    `json:"isDestroyable"`
	//Health        float64 `json:"health"`
}
