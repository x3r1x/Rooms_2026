package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}()
	fmt.Println("New connection established.")
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Message Received:", string(p))

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// scp -i "C:\Users\maxim\.ssh\id_ed25519.pub" server yc-user@84.201.159.214:/home/yc-user/
func main() {
	http.HandleFunc("/ws", handleWebsocket)
	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
