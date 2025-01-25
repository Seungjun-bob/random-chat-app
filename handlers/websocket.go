package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid" // UUID 생성용 패키지
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 모든 도메인 허용
	},
}

// HandleWebSocket handles WebSocket connections
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade WebSocket", http.StatusInternalServerError)
		fmt.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	// 고유 ID 생성
	clientID := uuid.NewString()
	fmt.Printf("New client connected: %s\n", clientID)

	// 클라이언트에게 ID 전송
	conn.WriteMessage(websocket.TextMessage, []byte("Your ID: "+clientID))

	// 대기열에 클라이언트 추가
	AddToQueue(clientID, conn)

	// 메시지 처리 루프
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Connection closed by client (%s): %v\n", clientID, err)

			// 연결 종료 시 대기열에서 제거
			RemoveFromQueue(clientID)
			break
		}
		fmt.Printf("[%s]: %s\n", clientID, string(message))
	}
}

// RemoveFromQueue removes a client from the matching queue
func RemoveFromQueue(clientID string) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := queue[clientID]; exists {
		delete(queue, clientID)
		fmt.Printf("Client removed from queue: %s. Queue size: %d\n", clientID, len(queue))
	}
}
