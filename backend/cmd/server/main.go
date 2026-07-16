package main

import (
	"fmt"
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
// TODO: добавить коллизии

func main() {
	go GameLoop()
	http.HandleFunc("/ws", websocket.InitWebsocket)
	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
