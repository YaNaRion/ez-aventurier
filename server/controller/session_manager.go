package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"main/infra/models"
	"net"
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
	sessionID string,
	r *http.Request,
) (bool, *models.Session, error) {

	session, err := c.db.FindSession(sessionID)
	if err != nil {
		return false, nil, fmt.Errorf("failed to connect session")
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return false, nil, fmt.Errorf("wrong host")
	}
	if session.Host != ip {
		return false, nil, fmt.Errorf("failed to connect session")
	}
	return !IsSessionExpired(session), session, nil
}

// Si la connection est bonne, il faut retourner au client sa sessionID
func (c *Controller) connection(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("user_id")
	if user_id == "" {
		log.Printf("Connection try from: %s, but was mission user_id in request", r.Host)
		http.Error(w, "Ce code secret correspond a personne", http.StatusBadRequest)
		return
	}

	user, err := c.db.FindUser(user_id)
	if err != nil {
		log.Printf("Error while request user to db: %s", err)
		http.Error(w, "Ce code secret correspond a personne", http.StatusBadRequest)
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println("Error while creating session")
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	session, err := c.db.AddSession(user.UserID, ip)
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
	Score  int    `bson:"score"  json:"score"  validate:"required"`
}

func (c *Controller) getUser(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	if sessionID == "" {
		http.Error(w, "Failed to find session in query param", http.StatusBadRequest)
		return
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		log.Println("Failed to find user_id in query param")
		http.Error(w, "Failed to find query param", http.StatusBadRequest)
		return
	}

	isSessionValid, _, err := c.isSessionValid(sessionID, r)
	if err != nil && !isSessionValid {
		http.Error(w, "Session not valid", http.StatusBadRequest)
		return
	}

	user, err := c.db.FindUser(userID)
	if err != nil {
		http.Error(w, "User not found not valid", http.StatusBadRequest)
		return
	}

	response_user := UserResponse{
		UserID: user.UserID,
		Name:   user.Name,
		Unity:  user.Unity,
		Ordre:  user.Ordre,
		Score:  user.Score,
	}

	user_json, err := json.Marshal(response_user)
	if err != nil {
		log.Println("Error while marchaling the user account")
		http.Error(w, "Ce code secret correspond a personne", http.StatusInternalServerError)
	}
	writeResponseJson(w, user_json)
}
