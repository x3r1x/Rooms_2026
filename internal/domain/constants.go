package domain

const (
	// константы для пуль
	BulletSpeedGun    = 20.0
	BulletSpeedRifle  = 25.0
	BulletSpeedSom    = 3.0
	BulletLifeGun     = 60.0
	BulletLifeRifle   = 40.0
	BulletLifeSom     = 600
	BulletDamageGun   = 12.0
	BulletDamageRifle = 5.0
	BulletDamageSom   = 100.0
	PelletCount       = 6
	SpreadAngle       = 0.4

	ShootCooldown     = 16
	BulletLengthGun   = 16.0
	BulletWidthGun    = 7.0
	BulletLengthRifle = 3.0
	BulletWidthRifle  = 3.0
	BulletLengthSom   = 11.0
	BulletWidthSom    = 3.0
	BulletDamageMulti = 1.2
	ExplosionRadius   = 250
	KnockbackForce    = 12.0
	// константы для игроков
	InitDirection     = 0
	InitValue         = 0
	PlayerVisualSize  = 34.0
	GunOffsetX        = 0.0
	GunOffsetY        = 20.5
	RifleOffsetX      = 12.0
	RifleOffsetY      = 19.0
	SomOffsetX        = 21.0
	SomOffsetY        = 19.0
	PlayerSpeedGun    = 0.6
	PlayerSpeedRifle  = 0.4
	PlayerSpeedSom    = 0.2
	MaxPlayerHealth   = 100
	PlayerRebornTimer = 300
	PlayerShieldTimer = 200
	// константы общие служебные
	TickTime                     = 16
	GameDuration                 = 30
	MaxPlayers                   = 30
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
	RoomPixelWidth               = float64(RoomWidth * int(TileSize))
	RoomPixelHeight              = float64(RoomHeight * int(TileSize))
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
