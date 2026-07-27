package main

import (
	"embed"
	"fmt"
	"gamedevRooms/internal/adapters/broadcast"
	"gamedevRooms/internal/adapters/collision"
	_map "gamedevRooms/internal/adapters/map"
	"gamedevRooms/internal/adapters/websocket"
	"gamedevRooms/internal/application/factory"
	"gamedevRooms/internal/application/game"
	"gamedevRooms/internal/application/lobby"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:frontend
var frontendFs embed.FS

func main() {
	broadcastService := broadcast.NewBroadcastService()
	gameState := game.NewGameState()
	mapManager := _map.NewMapManager()
	collisionService := collision.NewCollisionService(gameState)
	bulletFactory := factory.NewBulletFactory()

	lobbyService := lobby.NewLobbyService(
		gameState,
		mapManager,
		broadcastService,
	)

	gameService := game.NewGameService(
		gameState,
		collisionService,
		mapManager,
		broadcastService,
		bulletFactory,
		func() {
			lobbyService.HandleGameEnd()
		},
	)

	lobbyService.SetGameService(gameService)
	wsHandler := websocket.NewWebsocketHandler(lobbyService)

	http.HandleFunc("/ws", wsHandler.InitWebsocket)

	staticFiles, err := fs.Sub(frontendFs, "frontend")
	if err != nil {
		log.Fatal("Failed to load frontend files:", err)
	}

	http.Handle("/", http.FileServer(http.FS(staticFiles)))

	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
