package broadcast

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}

func BenchmarkBroadcastRealWS(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(wsHandler))
	defer server.Close()

	wsURL := "ws" + server.URL[4:]

	service := NewBroadcastService()
	conns := make([]*websocket.Conn, 500)

	for i := 0; i < 1; i++ {
		conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			b.Fatalf("Failed to dial: %v", err)
		}
		conns[i] = conn

		id := fmt.Sprintf("player_%d", i)
		service.AddConnection(id, conn)
	}

	msg := map[string]string{"data": "test"}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			playerId := fmt.Sprintf("player_%d", i%50)
			service.BroadcastToPlayer(playerId, msg)
			i++
		}
	})
	for _, conn := range conns {
		conn.Close()
	}
}
