package main

import (
	"fmt"
	"gamedevRooms/internal/game"
	"gamedevRooms/internal/websocket"
	"log"
	"net/http"
)

////go:embed all:frontend
//var frontendFs embed.FS

func main() {
	//newMap := mapModel.NewMap(7)
	//fmt.Println(newMap)

	gameLoop := game.NewGameLoop(game.NewGameState())
	go gameLoop.Run()

	wsHandler := websocket.NewWebsocketHandler(gameLoop)
	http.HandleFunc("/ws", wsHandler.InitWebsocket)

	//staticFiles, err := fs.Sub(frontendFs, "frontend")
	//if err != nil {
	//	log.Fatal("Failed to load frontend files:", err)
	//}

	//http.Handle("/", http.FileServer(http.FS(staticFiles)))

	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
