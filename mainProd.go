package main

import (
	"embed"
	"fmt"
	"gamedevRooms/internal/Lobby"
	"gamedevRooms/internal/websocket"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:frontend
var frontendFs embed.FS

func main() {
	lobby := Lobby.NewLobby()

	wsHandler := websocket.NewWebsocketHandler(lobby)
	http.HandleFunc("/ws", wsHandler.InitWebsocket)

	staticFiles, err := fs.Sub(frontendFs, "frontend")
	if err != nil {
		log.Fatal("Failed to load frontend files:", err)
	}

	http.Handle("/", http.FileServer(http.FS(staticFiles)))

	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
