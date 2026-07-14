package models

type ClientMessage struct {
	PlayerInterpolation struct {
		Direction struct {
			X int8 `json:"x"`
			Y int8 `json:"y"`
		} `json:"direction"`
		DeltaVisualDirection float64   `json:"deltaVisualDirection"`
		Id                   string    `json:"id"`
		NewBulletsDirection  []float64 `json:"newBulletsDirection"`
	} `json:"playerInterpolation"`
}

type AbsoluteServerMessage struct {
	Type    string                 `json:"type"`
	Players map[string]PlayerState `json:"players"`
}

type InterpolationServerMessage struct {
	Type                 string                         `json:"type"`
	PlayerInterpolations map[string]PlayerInterpolation `json:"playerInterpolations"`
}
