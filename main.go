package main

import (
	"fmt"
	"net/http"
	"random-chat-app/handlers"
)

func main() {
	// WebSocket 라우트 설정
	http.HandleFunc("/ws", handlers.HandleWebSocket)

	// 정적 파일 서빙
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// 서버 실행
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
