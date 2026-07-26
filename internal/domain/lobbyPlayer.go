package domain

type LobbyPlayer struct {
	Nickname    string
	Id          string
	Ready       bool
	PlayerClass string
}

func NewLobbyPlayer(nickname string) *LobbyPlayer {
	return &LobbyPlayer{
		Nickname:    nickname,
		Id:          GenerateID(),
		Ready:       false,
		PlayerClass: EmptyPlayer,
	}
}
