package main

import (
	"fmt"
	"gamedevRooms/internal/game"
	"gamedevRooms/internal/websocket"
	"log"
	"net/http"
)

// TODO: постепенно перевести на обмен данных между клиентом и сервером на сырые байты.
// TODO: нужна обратная мапа для подключений для упрощения поиска и мониторинга
// TODO: коллизия для пуль и их непосредственная обработка

func main() {
	gameState := game.NewGameState()
	gameLoop := game.NewGameLoop(gameState)
	go gameLoop.Run()

	wsHandler := websocket.NewWebsocketHandler(gameState)
	http.HandleFunc("/ws", wsHandler.InitWebsocket)

	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
