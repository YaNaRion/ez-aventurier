package gateway

import (
	"main/infra"
	"net/http"
)

// WebSocketHandler wraps the pool for HTTP handling
type WebSocketHandler struct {
	Pool *Pool
	db   *infra.DB
}

// NewWebSocketHandler creates a new WebSocket handler
func NewWebSocketHandler(db *infra.DB) *WebSocketHandler {
	pool := NewPool()
	go pool.Start()

	return &WebSocketHandler{
		Pool: pool,
		db:   db,
	}
}

// ServeHTTP handles HTTP requests
func (h *WebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ServeWs(h.Pool, h.db, w, r)
}

// Helper function to send event to specific client
func (h *WebSocketHandler) SendToClient(clientID string, event Event) bool {
	h.Pool.mu.RLock()
	defer h.Pool.mu.RUnlock()

	for client := range h.Pool.Clients {
		if client.ID == clientID {
			client.Send <- event
			return true
		}
	}
	return false
}

// Helper function to broadcast to all clients
func (h *WebSocketHandler) Broadcast(event Event) {
	h.Pool.Broadcast <- event
}

// Helper to get connected clients count
func (h *WebSocketHandler) GetClientCount() int {
	h.Pool.mu.RLock()
	defer h.Pool.mu.RUnlock()
	return len(h.Pool.Clients)
}
