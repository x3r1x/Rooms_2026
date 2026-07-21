package main

import (
	"embed"
	"fmt"
	"gamedevRooms/internal/game"
	mapModel "gamedevRooms/internal/gameMap"
	"gamedevRooms/internal/websocket"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:frontend
var frontendFs embed.FS

func main() {
	newMap := mapModel.NewMap(7)
	fmt.Println(newMap)

	gameState := game.NewGameState()
	gameLoop := game.NewGameLoop(gameState)
	go gameLoop.Run()

	wsHandler := websocket.NewWebsocketHandler(gameLoop)
	http.HandleFunc("/ws", wsHandler.InitWebsocket)

	staticFiles, err := fs.Sub(frontendFs, "frontend")
	if err != nil {
		log.Fatal("Failed to load frontend files:", err)
	}

	http.Handle("/", http.FileServer(http.FS(staticFiles)))

	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
