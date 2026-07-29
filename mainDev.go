package main

import (
	"fmt"
	"gamedevRooms/internal/adapters/broadcast"
	"gamedevRooms/internal/adapters/collision"
	"gamedevRooms/internal/adapters/map"
	"gamedevRooms/internal/adapters/websocket"
	"gamedevRooms/internal/application/game"
	"gamedevRooms/internal/application/lobby"
	"gamedevRooms/internal/recovery"
	"gamedevRooms/internal/state"
	"log"
	"net/http"
)

func main() {
	defer recovery.Recover()
	broadcastService := broadcast.NewBroadcastService()
	gameState := state.NewGameState()
	mapManager := _map.NewMapManager()
	collisionService := collision.NewCollisionService(gameState)

	lobbyService := lobby.NewLobbyService(
		gameState,
		mapManager,
		broadcastService,
	)

	gameService := game.NewGameService(
		gameState,
		collisionService,
		broadcastService,
		func() {
			defer recovery.Recover()
			lobbyService.HandleGameEnd()
		},
	)

	lobbyService.SetGameService(gameService)
	wsHandler := websocket.NewWebsocketHandler(lobbyService)

	http.HandleFunc("/ws", wsHandler.InitWebsocket)

	fmt.Println("Server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
