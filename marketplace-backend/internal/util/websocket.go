package util

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan Message)           // Broadcast channel
var upgrader = websocket.Upgrader{}
var lock sync.Mutex // To protect the clients map

type Message struct {
	Sender    uint   `json:"sender"`
	Receiver  uint   `json:"receiver"`
	Content   string `json:"content"`
}

// HandleChatConnections upgrades the HTTP connection to a WebSocket and handles incoming messages.
func HandleChatConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // Set to check the origin in production
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(w, "Upgrade: %v", err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			lock.Lock()
			delete(clients, conn)
			lock.Unlock()
			break
		}
		broadcast <- msg
	}
}

// StartMessageHandler listens for new messages and broadcasts them to the correct receiver.
func StartMessageHandler() {
	for {
		msg := <-broadcast
		lock.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		lock.Unlock()
	}
}
// Placeholder for storing active WebSocket connections
	var clients = make(map[*websocket.Conn]bool)

// Broadcast to all clients
	func BroadcastNotification(notification map[string]interface{}) {
	for client := range clients {
		err := client.WriteJSON(notification)
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}
	