package util

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Add proper CheckOrigin function for production to ensure security
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	// Maintain a map of clients with the user ID as the key for targeted messaging
	clients = make(map[uint][]*websocket.Conn)
	broadcast = make(chan Message)
	lock sync.RWMutex
)

type Message struct {
	Sender    uint   `json:"sender"`
	Receiver  uint   `json:"receiver"`
	Content   string `json:"content"`
}

// HandleChatConnections upgrades the HTTP connection to a WebSocket.
func HandleChatConnections(w http.ResponseWriter, r *http.Request, userID uint) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade error: %v", err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	// Register new client
	lock.Lock()
	clients[userID] = append(clients[userID], conn)
	lock.Unlock()

	// Listen for messages
	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("error on ws message: %v", err)
			removeClient(conn, userID)
			break
		}
		broadcast <- msg
	}
}

// StartMessageHandler listens and routes messages.
func StartMessageHandler() {
	for msg := range broadcast {
		lock.RLock()
		receivers, exists := clients[msg.Receiver]
		if exists {
			for _, client := range receivers {
				if err := client.WriteJSON(msg); err != nil {
					log.Printf("error broadcasting to client: %v", err)
					client.Close()
					removeClient(client, msg.Receiver)
				}
			}
		}
		lock.RUnlock()
	}
}

// removeClient safely removes a WebSocket connection from the clients map.
func removeClient(conn *websocket.Conn, userID uint) {
	lock.Lock()
	defer lock.Unlock()
	for i, client := range clients[userID] {
		if client == conn {
			clients[userID] = append(clients[userID][:i], clients[userID][i+1:]...)
			break
		}
	}
	if len(clients[userID]) == 0 {
		delete(clients, userID)
	}
}
