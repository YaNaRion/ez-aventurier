package controller

import (
	"encoding/json"
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

func (c *Controller) isSessionValid(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incomming session validation from: %s", r.Host)
	log.Printf("Incomming session validation from: %s", r.URL)

	session_id := r.URL.Query().Get("session_id")
	user_id := r.URL.Query().Get("user_id")

	log.Println(user_id)
	log.Println(session_id)
	if session_id == "" || user_id == "" {
		http.Error(w, "Missing SessionID or UserID", http.StatusBadRequest)
		return
	}

	session, err := c.db.FindSession(session_id)
	if err != nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	if IsSessionExpired(session) {
		http.Error(w, "Session not valid", http.StatusForbidden)
		return
	}
	log.Println("APRES SESSION EXPIRE")

	response := IsSessionValidResponse{
		Session: *session,
		IsValid: true,
	}

	sessionJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Could not marshal the response", http.StatusForbidden)
		return
	}

	log.Println("APRES JSON MARSHAL")
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
		log.Printf("Connection try from: %s, but was mission user_id in request", r.Host)
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

	user_json, err := json.Marshal(session)
	if err != nil {
		log.Println("Error while marchaling the user account")
		http.Error(w, "Ce code secret correspond a personne", http.StatusInternalServerError)
	}
	writeResponseJson(w, user_json)
}
