package main

import (
	"fmt"
	"net/http"
	"random-chat-app/handlers"
)

func main() {
	// WebSocket 핸들러
	http.HandleFunc("/ws", handlers.HandleWebSocket)

	// 정적 파일 서빙
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// 서버 시작
	fmt.Println("Server is running on http://0.0.0.0:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
