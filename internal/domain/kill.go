package domain

type Kill struct {
	KillerId string  `json:"kId"`
	VictimId string  `json:"vId"`
	Time     float64 `json:"-"`
}

func NewKill(killerId, victimId string, time float64) *Kill {
	return &Kill{
		KillerId: killerId,
		VictimId: victimId,
		Time:     time,
	}
}
