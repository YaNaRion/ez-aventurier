package gateway

import (
	"encoding/json"
	"log"
	"main/infra"
	"time"
)

// NewRegisterEventHandlers returns all registered event handlers
func NewRegisterEventHandlers(db *infra.DB) EventHandlers {
	eventHandlers := NewEventHandlers(db)

	// Register built-in handlers
	eventHandlers.Handlers[EventTypeMessage] = eventHandlers.handleMessage
	eventHandlers.Handlers[EventTypeJoin] = eventHandlers.handleJoin
	eventHandlers.Handlers[EventTypeLeave] = eventHandlers.handleLeave

	// Register your custom event eventHandlers.Handlers here
	eventHandlers.Handlers[EventTypeChat] = eventHandlers.handleChat
	eventHandlers.Handlers[EventTypeNotification] = eventHandlers.handleNotification

	return eventHandlers
}

// Built-in event handlers
func (eh *EventHandlers) handleMessage(client *Client, data json.RawMessage) error {
	var messageData map[string]string
	if err := json.Unmarshal(data, &messageData); err != nil {
		return err
	}

	// Process message
	log.Printf("Message from %s: %v", client.ID, messageData)

	// Broadcast or send response
	response := Event{
		Type: EventTypeMessage,
		Data: map[string]string{
			"from":    client.ID,
			"message": messageData["content"],
		},
	}

	client.Pool.Broadcast <- response
	eh.DB.HelloWorld()
	return nil
}

func (eh *EventHandlers) handleJoin(client *Client, data json.RawMessage) error {
	log.Printf("Client %s joined", client.ID)

	// Broadcast join notification
	joinEvent := Event{
		Type: EventTypeJoin,
		Data: map[string]string{
			"clientId": client.ID,
		},
	}

	client.Pool.Broadcast <- joinEvent
	return nil
}

func (eh *EventHandlers) handleLeave(client *Client, data json.RawMessage) error {
	log.Printf("Client %s left", client.ID)
	return nil
}

// Custom event handlers - ADD NEW EVENTS HERE
func (eh *EventHandlers) handleChat(client *Client, data json.RawMessage) error {
	var chatData struct {
		Message string `json:"message"`
		Room    string `json:"room"`
	}

	if err := json.Unmarshal(data, &chatData); err != nil {
		return err
	}

	log.Printf("Chat message from %s in room %s: %s",
		client.ID, chatData.Room, chatData.Message)

	// Example: Broadcast to room
	response := Event{
		Type: EventTypeChat,
		Data: map[string]interface{}{
			"from":      client.ID,
			"room":      chatData.Room,
			"message":   chatData.Message,
			"timestamp": time.Now().Unix(),
		},
	}

	// You could implement room-based broadcasting here
	client.Pool.Broadcast <- response
	return nil
}

func (eh *EventHandlers) handleNotification(client *Client, data json.RawMessage) error {
	var notificationData struct {
		Title string `json:"title"`
		Body  string `json:"body"`
		To    string `json:"to"` // Specific client ID or "all"
	}

	if err := json.Unmarshal(data, &notificationData); err != nil {
		return err
	}

	log.Printf("Notification: %s - %s", notificationData.Title, notificationData.Body)

	// Implement targeted notification logic
	response := Event{
		Type: EventTypeNotification,
		Data: notificationData,
	}

	// Broadcast to all or specific clients
	client.Pool.Broadcast <- response
	return nil
}

// Helper function to send error to client
func sendError(client *Client, errorMessage string) {
	errorEvent := Event{
		Type: EventTypeError,
		Data: map[string]string{
			"message": errorMessage,
		},
	}

	client.Send <- errorEvent
}

// func (eh *EventHandlers)handleYourNewEvent(client *Client, data json.RawMessage) error {
// 	// Parse your event data
// 	var eventData any
// 	if err := json.Unmarshal(data, &eventData); err != nil {
// 		return err
// 	}
//
// 	// Process the event
// 	log.Printf("Processing new event from %s: %v", client.ID, eventData)
//
// 	// Send response if needed
// 	response := Event{
// 		Type: EventTypeYourNewEvent,
// 		Data: map[string]interface{}{
// 			"status": "processed",
// 			"data":   eventData,
// 		},
// 	} // Send to client, broadcast, or send to specific clients
// 	client.Send <- response
// 	// or: client.Pool.Broadcast <- response
//
// 	return nil
// }
