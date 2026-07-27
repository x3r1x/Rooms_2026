package main

import (
	"fmt"
	"gamedevRooms/internal/adapters/broadcast"
	"gamedevRooms/internal/adapters/collision"
	"gamedevRooms/internal/adapters/map"
	"gamedevRooms/internal/adapters/websocket"
	"gamedevRooms/internal/application/game"
	"gamedevRooms/internal/application/lobby"
	"log"
	"net/http"
)

func main() {

	broadcastService := broadcast.NewBroadcastService()
	gameState := game.NewGameState()
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
		mapManager,
		broadcastService,
		func() {
			lobbyService.HandleGameEnd()
		},
	)

	lobbyService.SetGameService(gameService)
	wsHandler := websocket.NewWebsocketHandler(lobbyService)

	http.HandleFunc("/ws", wsHandler.InitWebsocket)

	fmt.Println("Server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
