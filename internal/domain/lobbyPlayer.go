package domain

type LobbyPlayer struct {
	Nickname string
	Id       string
	Ready    bool
}

func NewLobbyPlayer(nickname string) *LobbyPlayer {
	return &LobbyPlayer{
		Nickname: nickname,
		Id:       GenerateID(),
		Ready:    false,
	}
}
