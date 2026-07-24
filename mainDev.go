package main

import (
	"fmt"
	"gamedevRooms/internal/adapters/websocket"
	"gamedevRooms/internal/lobby"
	"log"
	"net/http"
)

func main() {
	lobby := lobby.NewLobby()

	wsHandler := websocket.NewWebsocketHandler(lobby)
	http.HandleFunc("/ws", wsHandler.InitWebsocket)

	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
