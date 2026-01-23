package gateway

import (
	"encoding/json"
	"fmt"
	"log"
	"main/infra"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Configure your CORS policy here
		return true
	},
}

// NewPool creates a new client pool
func NewPool() *Pool {
	return &Pool{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Event),
	}
}

// Start begins the pool's event loop
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.mu.Lock()
			pool.Clients[client] = true
			pool.mu.Unlock()
			log.Printf("Client registered. Total: %d", len(pool.Clients))

		case client := <-pool.Unregister:
			pool.mu.Lock()
			if _, ok := pool.Clients[client]; ok {
				delete(pool.Clients, client)
				close(client.Send)
			}
			pool.mu.Unlock()
			log.Printf("Client unregistered. Total: %d", len(pool.Clients))

		case event := <-pool.Broadcast:
			pool.mu.RLock()
			for client := range pool.Clients {
				select {
				case client.Send <- event:
				default:
					close(client.Send)
					delete(pool.Clients, client)
				}
			}
			pool.mu.RUnlock()
		}
	}
}

// readPump reads messages from the WebSocket connection
func (c *Client) readPump(handlers EventHandlers) {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Printf("Read error: %v", err)
			}
			break
		}

		// Parse the incoming event
		var event Event
		if err := json.Unmarshal(message, &event); err != nil {
			sendError(c, "Invalid event format")
			continue
		}

		// Find and execute the event handler
		if handler, exists := handlers.Handlers[event.Type]; exists {
			// Convert data to JSON for handler
			dataBytes, _ := json.Marshal(event.Payload)
			if err := handler(c, dataBytes); err != nil {
				sendError(c, fmt.Sprintf("Error processing event: %v", err))
			}
		} else {
			sendError(c, fmt.Sprintf("Unknown event type: %s", event.Type))
		}
	}
}

// writePump writes messages to the WebSocket connection
func (c *Client) writePump() {
	ticker := time.NewTicker(60 * time.Second) // Ping interval
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case event, ok := <-c.Send:
			if !ok {
				// Channel closed
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			// Write the event as JSON
			message, err := json.Marshal(event)
			if err != nil {
				log.Printf("Error marshaling event: %v", err)
				continue
			}

			c.mu.Lock()
			err = c.Conn.WriteMessage(websocket.TextMessage, message)
			c.mu.Unlock()

			if err != nil {
				log.Printf("Write error: %v", err)
				return
			}

		case <-ticker.C:
			// Send ping to keep connection alive
			c.mu.Lock()
			err := c.Conn.WriteMessage(websocket.PingMessage, nil)
			c.mu.Unlock()
			if err != nil {
				return
			}
		}
	}
}

// ServeWs handles WebSocket requests
func ServeWs(pool *Pool, db *infra.DB, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}

	// Generate client ID (you might want to use JWT or session ID)
	clientID := fmt.Sprintf("client-%d", time.Now().UnixNano())

	client := &Client{
		ID:   clientID,
		Conn: conn,
		Pool: pool,
		Send: make(chan Event, 256),
	}

	// Register event handlers with db dependencie
	handlers := NewRegisterEventHandlers(db)

	// Register client with pool
	pool.Register <- client

	// Start goroutines for reading and writing
	go client.writePump()
	go client.readPump(handlers)
}
