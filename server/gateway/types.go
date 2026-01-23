package gateway

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"main/infra"
	"sync"
)

// EventType represents different types of WebSocket events
type EventType string

const (
	// Built-in events
	EventTypeError   EventType = "error"
	EventTypeMessage EventType = "message"
	EventTypeJoin    EventType = "join"
	EventTypeLeave   EventType = "leave"

	// Add your custom events here
	EventTypeChat          EventType = "chat"
	EventTypeNotification  EventType = "notification"
	EventTypeLoginRequest  EventType = "login.request"
	EventTypeLoginResponse EventType = "login.response"

	// Template
	EventTypeYourNewEvent EventType = "place_new_event_here"
)

// Event represents a WebSocket event
type Event struct {
	Type          EventType       `json:"type"`
	Payload       json.RawMessage `json:"payload,omitempty"`
	Timestamp     *int64          `json:"timestamp,omitempty"`
	EventUniqueID *string         `json:"eventUniqueID,omitempty"`
}

// Client represents a WebSocket client
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
	Send chan Event
}

// Pool manages all connected clients
type Pool struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Event
	mu         sync.RWMutex
}

// EventHandler handles specific event types
type EventHandler func(client *Client, data json.RawMessage) error

// EventHandlers maps event types to their handlers
type EventHandlers struct {
	Handlers map[EventType]EventHandler
	DB       *infra.DB
}

func NewEventHandlers(db *infra.DB) EventHandlers {
	return EventHandlers{
		DB:       db,
		Handlers: make(map[EventType]EventHandler),
	}
}
