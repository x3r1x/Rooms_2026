package model

const (
	// константы для пуль
	MaxBulletSpeed = 22.5
	BulletLife     = 60
	ShootCooldown  = 15
	BulletWidth    = 5.0
	BulletDamage   = 10
	// константы для игроков
	PlayerSpawnPointX = 500
	PlayerSpawnPointY = 500
	InitDirection     = 0
	PlayerVisualSize  = 40.0
	PlayerSpeed       = float64(0.2)
	PlayerHalfSize    = 40.0
	MaxPlayerHealth   = 100
	//константы для окружения и полей
	CellSize = 40.0
	// общие служебные координаты
	TickTime = 16
)
