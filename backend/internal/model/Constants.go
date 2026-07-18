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
	PlayerVisualSize  = 40.0
	PlayerSpeed       = float64(0.2)
	MaxPlayerHealth   = 100
	PlayerRebornTimer = 300
	// общие служебные координаты
	TickTime = 16
)
