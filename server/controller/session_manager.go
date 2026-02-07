package controller

import (
	"encoding/json"
	"log"
	"main/infra/models"
	"net/http"
	"time"
)

func IsSessionExpired(session *models.Session) bool {
	return time.Since(session.CreatedON) > 10*time.Minute
}

func (c *Controller) isSessionValid(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("sessionID")
	if sessionID == "" {
		http.Error(w, "Missing SessionID", http.StatusBadRequest)
		return
	}

	session, err := c.db.FindSession(sessionID)

	if err != nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	if !IsSessionExpired(session) {
		http.Error(w, "Session not valid", http.StatusForbidden)
		return
	}

	sessionJson, err := json.Marshal(*session)
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
	log.Println("DANS CONNECTION REQUEST")
	// vérification des informations de connection dans le server

	// le login est bon, je fais les cookies
	token := "Voici le session ID"
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		// Secure:   true, // Use with HTTPS a rajouter quand le https sera
		SameSite: http.SameSiteStrictMode,
		MaxAge:   10 * 60, // 10 minutes
	})
}
