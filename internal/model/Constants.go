package model

const (
	// константы для пуль
	MaxBulletSpeed    = 22.5
	BulletLife        = 60.0
	ShootCooldown     = 15
	BulletLength      = 25.0
	BulletWidth       = 5.0
	BulletDamage      = 10.0
	BulletDamageMulti = 1.2
	// константы для игроков
	PlayerSpawnPointX = 500
	PlayerSpawnPointY = 500
	InitDirection     = 0
	InitValue         = 0
	PlayerVisualSize  = 40.0
	PlayerSpeed       = 0.2
	MaxPlayerHealth   = 100
	PlayerRebornTimer = 300
	// константы общие служебные
	TickTime                     = 16
	GameDuration                 = 60
	MaxCountOfPlayers            = 2
	MinCountOfPlayers            = 2
	MinBarrierType               = 1
	MaxBarrierType               = 7
	BaseRoomIndex                = 0
	BaseWallsIndex               = 1
	ExitTopIndex                 = 2
	ExitLeftIndex                = 3
	ExitBottomIndex              = 4
	ExitRightIndex               = 5
	FlapTopIndex                 = 6
	FlapLeftIndex                = 7
	FlapBottomIndex              = 8
	FlapRightIndex               = 9
	BarriersStartIndex           = 10
	BarriersAmount               = 7
	RoomHeight                   = 21
	RoomWidth                    = 25
	TopMarker                    = "top"
	LeftMarker                   = "left"
	BottomMarker                 = "bottom"
	RightMarker                  = "right"
	ConnectNeighbouredRoomChance = 0.4
	// константы для лобби
	WaitingLobbyState   = "w"
	ReadyLobbyState     = "r"
	CountdownLobbyState = "c"
	OngoingGameState    = "o"
	EndedGameState      = "e"
)
