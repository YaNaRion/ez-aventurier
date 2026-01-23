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

	loginRequestPayload.UniqueID = "Reponse du server"
	payloadJson, err := json.Marshal(loginRequestPayload)
	if err != nil {
		log.Println("Failed to marchar payload response: %v", err)
	}

	response := Event{
		Type:    EventTypeLoginResponse,
		Payload: payloadJson,
	}

	client.Send <- response
	return nil
}

// Helper function to send error to client
func sendError(client *Client, errorMessage string) {
	payloadMessage := map[string]string{
		"message": errorMessage,
	}

	payloadJSON, err := json.Marshal(payloadMessage)

	if err != nil {
		log.Println("Error while marshal error message")
	}

	timeStamp := time.Now().UnixMilli()
	errorID := "ERROR ID"
	errorEvent := Event{
		Type:          EventTypeError,
		Payload:       payloadJSON,
		Timestamp:     &timeStamp,
		EventUniqueID: &errorID,
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
