package handlers

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// StartChatSession starts a chat session between two clients
func StartChatSession(client1, client2 *websocket.Conn) {
	defer client1.Close()
	defer client2.Close()

	// 매칭 완료 메시지 전송
	client1.WriteMessage(websocket.TextMessage, []byte("MATCHED"))
	client2.WriteMessage(websocket.TextMessage, []byte("MATCHED"))

	fmt.Println("Chat session started between two clients")

	// 메시지 교환
	for {
		messageType1, message1, err1 := client1.ReadMessage()
		if err1 != nil {
			fmt.Println("Client 1 disconnected")
			break
		}

		messageType2, message2, err2 := client2.ReadMessage()
		if err2 != nil {
			fmt.Println("Client 2 disconnected")
			break
		}

		client2.WriteMessage(messageType1, message1)
		client1.WriteMessage(messageType2, message2)
	}
}

// readMessages reads messages from a WebSocket client
func readMessages(client *websocket.Conn, ch chan<- string) {
	for {
		_, msg, err := client.ReadMessage()
		if err != nil {
			ch <- "disconnect"
			break
		}
		ch <- string(msg)
	}
}
