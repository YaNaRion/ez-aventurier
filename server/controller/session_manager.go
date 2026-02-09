package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"main/infra/models"
	"net/http"
	// "time"
)

func IsSessionExpired(session *models.Session) bool {
	// return time.Since(session.CreatedON) > 10*time.Minute
	return false
}

type IsSessionValidResponse struct {
	Session models.Session `json:"session"`
	IsValid bool           `json:"isValid"`
}

func (c *Controller) chechSession(session_id string) (*models.Session, error) {
	session, err := c.db.FindSession(session_id)
	if err != nil {
		return nil, fmt.Errorf("failed to find session: %w", err)
	}

	if IsSessionExpired(session) {
		return nil, fmt.Errorf("failed to find session: %w", err)
	}
	return session, nil
}

func (c *Controller) isSessionValid(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incomming session validation from: %s", r.Host)

	session_id := r.URL.Query().Get("session_id")
	user_id := r.URL.Query().Get("user_id")
	if session_id == "" || user_id == "" {
		http.Error(w, "Missing SessionID or UserID", http.StatusBadRequest)
		return
	}

	session, err := c.chechSession(session_id)

	if err != nil {
		http.Error(w, "Missing SessionID or UserID", http.StatusBadRequest)
		return
	}

	response := IsSessionValidResponse{
		Session: *session,
		IsValid: true,
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
	log.Println(user_id)

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

	session, err := c.db.AddSession(user.UserID)
	if err != nil {
		log.Println("Error while creating session")
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	// le login est bon, je fais les cookies
	// token := "Voici le session ID"
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "auth_token",
	// 	Value:    session.SessionID,
	// 	Path:     "/",
	// 	HttpOnly: true,
	// 	// Secure:   true, // Use with HTTPS a rajouter quand le https sera
	// 	SameSite: http.SameSiteStrictMode,
	// 	MaxAge:   10 * 60, // 10 minutes
	// })

	session_json, err := json.Marshal(session)
	if err != nil {
		log.Println("Error while marchaling the session account")
		http.Error(w, "Ce code secret correspond a personne", http.StatusInternalServerError)
	}
	writeResponseJson(w, session_json)
}

func (c *Controller) getUser(w http.ResponseWriter, r *http.Request) {
	log.Println("GET USER")
	session_id := r.URL.Query().Get("session_id")
	user_id := r.URL.Query().Get("user_id")

	_, err := c.chechSession(session_id)

	if err != nil {
		http.Error(w, "Missing SessionID or UserID", http.StatusBadRequest)
		return
	}

	user, err := c.db.FindUser(user_id)

	user_json, err := json.Marshal(user)
	if err != nil {
		log.Println("Error while marchaling the user account")
		http.Error(w, "Ce code secret correspond a personne", http.StatusInternalServerError)
	}
	writeResponseJson(w, user_json)
}
