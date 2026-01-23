package gateway

import (
	"encoding/json"
	"log"
	"main/infra"
)

// NewRegisterEventHandlers returns all registered event handlers
func NewRegisterEventHandlers(db *infra.DB) EventHandlers {
	eventHandlers := NewEventHandlers(db)

	// Real Event
	eventHandlers.Handlers[EventTypeLoginRequest] = eventHandlers.handleLoginRequest

	return eventHandlers
}

func (eh *EventHandlers) handleLoginRequest(client *Client, data json.RawMessage) error {
	var loginRequestPayload struct {
		UniqueID string `json:"uniqueID"`
	}

	if err := json.Unmarshal(data, &loginRequestPayload); err != nil {
		log.Printf("Failed to unmarshal payload: %v\n", err)
		return err
	}

	log.Printf("Login Payload: UniqueID=%s\n", loginRequestPayload.UniqueID)

	// Process login...
	return nil
}

// TODO: Make it work with the new Event type
// Helper function to send error to client
func sendError(client *Client, errorMessage string) {
	log.Println("An error has occured")
	// errorEvent := Event{
	// 	Type: EventTypeError,
	// 	Payload: map[string]string{
	// 		"message": errorMessage,
	// 	},
	// }
	//
	// client.Send <- errorEvent
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
