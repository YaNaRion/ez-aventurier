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

func (eh *EventHandlers) SendResponse(client *Client, eventResponse Event) {
	client.Send <- eventResponse
}

func (eh *EventHandlers) handleLoginRequest(client *Client, data json.RawMessage) error {
	var loginRequestPayload EventLoginRequest

	if err := json.Unmarshal(data, &loginRequestPayload); err != nil {
		log.Printf("Failed to unmarshal payload: %v\n", err)
		return err
	}

	user, err := eh.DB.FindUser(loginRequestPayload.UniqueID)
	if err != nil {
		log.Printf("An error has occured while login Request: %v \n", err)
		return nil
	}

	// TODO envoyer une reponse au client que le serveur na pas trouvé d'utilisateur
	if user == nil {
		log.Printf("No user found with this ID: %s\n", loginRequestPayload.UniqueID)
		return nil
	}

	log.Printf("Login Payload: UniqueID=%s\n", loginRequestPayload.UniqueID)
	loginRequestPayload.UniqueID = user.Name
	payloadJson, err := json.Marshal(loginRequestPayload)
	if err != nil {
		log.Printf("Failed to marchar payload response: %v\n", err)
	}

	eh.SendResponse(client, Event{
		Type:    EventTypeLoginResponse,
		Payload: payloadJson,
	})

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
