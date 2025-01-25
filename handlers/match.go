package handlers

import (
	"fmt"
	"net/http"
)

// WebSocket 연결 처리 함수
func HandleMatch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("match connection requested")
}
