package handlers

import (
	"fmt"
	"net/http"
)

// WebSocket 연결 처리 함수
func HandleCjat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("chat connection requested")
}
