package handlers

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	mu    sync.Mutex                         // 동시성 제어
	queue = make(map[string]*websocket.Conn) // UUID와 WebSocket 연결 매핑
)

// AddToQueue adds a client (UUID and connection) to the matching queue
func AddToQueue(clientID string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	// 대기열에 클라이언트 추가
	queue[clientID] = conn
	fmt.Printf("Client added to queue: %s. Queue size: %d\n", clientID, len(queue))

	// 두 명 이상 대기 중일 때 매칭
	if len(queue) >= 2 {
		// 대기열에서 첫 두 클라이언트 선택
		var client1ID, client2ID string
		var client1, client2 *websocket.Conn
		i := 0
		for id, conn := range queue {
			if i == 0 {
				client1ID, client1 = id, conn
			} else if i == 1 {
				client2ID, client2 = id, conn
				break
			}
			i++
		}

		// 매칭된 두 클라이언트를 대기열에서 제거
		delete(queue, client1ID)
		delete(queue, client2ID)

		fmt.Printf("Starting chat session between %s and %s\n", client1ID, client2ID)

		// 채팅 세션 시작
		go StartChatSession(client1, client2)
	}
}
