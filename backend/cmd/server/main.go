package main

import (
	"fmt"
	"gamedevRooms/internal/game"
	"gamedevRooms/internal/websocket"
	"log"
	"net/http"
)

// TODO: уменьшить объем обмена, чтобы уменьшить нагрузку
// TODO: постепенно перевести на обмен данных между клиентом и сервером на сырые байты.
// TODO: разбить на файлы, для нормального поддержания
// TODO: нужна обратная мапа для подключений для упрощения поиска и мониторинга
// TODO: коллизия для пуль и их непосредственная обработка
// TODO: создать отдельный цикл на обработку пуль и движения, наконец добиться нормального распределения

// scp -i "C:\Users\maxim\.ssh\id_ed25519.pub" server yc-user@84.201.159.214:/home/yc-user/
func main() {
	go game.StartGameLoop()
	http.HandleFunc("/ws", websocket.HandleWebsocket)
	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
