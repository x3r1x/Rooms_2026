package domain

const (
	// константы для пуль
	BulletSpeedGun    = 25.0
	BulletSpeedRifle  = 30.0
	BulletSpeedSom    = 10.0
	BulletLifeGun     = 60.0
	BulletLifeRifle   = 20.0
	BulletLifeSom     = 600
	BulletDamageGun   = 20.0
	BulletDamageRifle = 30.0
	BulletDamageSom   = 100.0

	ShootCooldown      = 15
	BulletLength       = 25.0
	BulletWidth        = 5.0
	BulletDamageMulti  = 1.2
	BulletBarrelOffset = 24.0
	RecoilDistance     = 8.0
	KnockbackForce     = 12.0
	// константы для игроков
	PlayerSpawnPointX = 710
	PlayerSpawnPointY = 400
	InitDirection     = 0
	InitValue         = 0
	PlayerVisualSize  = 34.0
	PlayerSpeed       = 0.4
	MaxPlayerHealth   = 100
	PlayerRebornTimer = 300
	// константы общие служебные
	TickTime                     = 16
	GameDuration                 = 20
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
	TileSize                     = 36
	PlayerHalfSize               = PlayerVisualSize / 2.0
	Epsilon                      = 0.001
	// константы для лобби
	WaitingLobbyState   = "w"
	ReadyLobbyState     = "r"
	CountdownLobbyState = "c"
	OngoingGameState    = "o"
	FinalGameState      = "f"

	PlayerGun   = "g"
	PlayerRifle = "r"
	PlayerSom   = "s"
	EmptyPlayer = "e"
)
