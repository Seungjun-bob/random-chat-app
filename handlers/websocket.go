package handlers

import (
	"fmt"
	"net/http"
)

// WebSocket 연결 처리 함수
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket connection requested")
}
