package main

import (
	"fmt"
	"gamedevRooms/internal/game"
	"gamedevRooms/internal/websocket"
	"log"
	"net/http"
)

// TODO: Сделать обработку смертей, а конкретно понять, что делать с игроком, когда его забили на смерть

func main() {
	gameState := game.NewGameState()
	gameLoop := game.NewGameLoop(gameState)
	go gameLoop.Run()

	wsHandler := websocket.NewWebsocketHandler(gameLoop)
	http.HandleFunc("/ws", wsHandler.InitWebsocket)

	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
