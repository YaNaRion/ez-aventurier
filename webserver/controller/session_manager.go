package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"main/infra/models"
	"net/http"
	"time"
	// "time"
)

func IsSessionExpired(session *models.Session) bool {
	return time.Since(session.CreatedON) > 10*time.Minute
}

type IsSessionValidResponse struct {
	Session models.Session `json:"session"`
	IsValid bool           `json:"isValid"`
}

func (c *Controller) isSessionValid(
	session *models.Session,
	urlHost string,
) (bool, error) {
	if session.Host != urlHost {
		return false, fmt.Errorf("failed to connect session")
	}
	return !IsSessionExpired(session), nil
}

func (c *Controller) isSessionValidMiddle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incomming session validation from: %s", r.Host)

	session_id := r.URL.Query().Get("session_id")
	if session_id == "" {
		http.Error(w, "Missing SessionID or UserID", http.StatusBadRequest)
		return
	}

	session, err := c.db.FindSession(session_id)
	if err != nil {
		http.Error(w, "Failed to find session", http.StatusBadRequest)
		return
	}

	isSessionValid, err := c.isSessionValid(session, r.URL.Host)

	if err != nil {
		http.Error(w, "Missing SessionID or UserID", http.StatusBadRequest)
		return
	}

	response := IsSessionValidResponse{
		Session: *session,
		IsValid: isSessionValid,
	}

	sessionJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Could not marshal the response", http.StatusForbidden)
		return
	}

	writeResponseJson(w, sessionJson)
	// response := "true"
	// w.Write([]byte(response))
	// http.Error(w, "", http.StatusBadRequest)
}

// Si la connection est bonne, il faut retourner au client sa sessionID
func (c *Controller) connection(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("user_id")

	if user_id == "" {
		log.Printf("Connection try from: %s, but was mission user_id in request", r.Host)
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}

	user, err := c.db.FindUser(user_id)
	if err != nil {
		log.Printf("Error while request user to db: %s", err)
		http.Error(w, "Ce code secret correspond a personne", http.StatusBadRequest)
		return
	}
	// FACILITE LA CONNECTION POUR TEST
	// log.Println("Un client a réussi a se connecter")
	// response := "true"
	// w.Write([]byte(response))
	// return
	// vérification des informations de connection dans le server

	session, err := c.db.AddSession(user.UserID, r.URL.Host)
	if err != nil {
		log.Println("Error while creating session")
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	session_json, err := json.Marshal(session)
	if err != nil {
		log.Println("Error while marchaling the session account")
		http.Error(w, "Ce code secret correspond a personne", http.StatusInternalServerError)
	}
	writeResponseJson(w, session_json)
}

type UserResponse struct {
	Name   string `bson:"name"   json:"name"   validate:"required"`
	UserID string `bson:"userId" json:"userId" validate:"required,min=8,max=8"`
	Ordre  string `bson:"order"  json:"order"  validate:"required"`
	Unity  string `bson:"unity"  json:"unity"  validate:"required"`
}

func (c *Controller) getUser(w http.ResponseWriter, r *http.Request) {
	log.Println("GET USER")
	session_id := r.URL.Query().Get("session_id")
	user_id := r.URL.Query().Get("user_id")

	session, err := c.db.FindSession(session_id)
	if err != nil {
		http.Error(w, "Failed to find session", http.StatusBadRequest)
		return
	}

	isSessionValid, err := c.isSessionValid(session, r.URL.Host)
	if err != nil && !isSessionValid {
		http.Error(w, "Session not valid", http.StatusBadRequest)
		return
	}

	user, err := c.db.FindUser(user_id)
	response_user := UserResponse{
		UserID: user_id,
		Name:   user.Name,
		Unity:  user.Unity,
		Ordre:  user.Ordre,
	}

	user_json, err := json.Marshal(response_user)
	if err != nil {
		log.Println("Error while marchaling the user account")
		http.Error(w, "Ce code secret correspond a personne", http.StatusInternalServerError)
	}
	writeResponseJson(w, user_json)
}
