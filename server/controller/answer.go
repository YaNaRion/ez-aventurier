package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func (c *Controller) claimCache(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	if sessionID == "" {
		http.Error(w, "Failed to find session", http.StatusBadRequest)
		return
	}

	isSessionValid, _, err := c.isSessionValid(sessionID, r.URL.Host)
	if err != nil && !isSessionValid {
		http.Error(w, "Session not valid", http.StatusBadRequest)
		return
	}

	answerID := r.URL.Query().Get("answer_id")
	if answerID == "" {
		log.Println("Failed to find answer_id in query param")
		http.Error(w, "Failed to find query param", http.StatusBadRequest)
		return
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		log.Println("Failed to find user_id in query param")
		http.Error(w, "Failed to find query param", http.StatusBadRequest)
		return
	}

	weight, err := c.db.ModifyAnswer(answerID, userID)
	if err != nil {
		log.Println("Failed to find answer")
		http.Error(w, "Failed to find answer", http.StatusBadRequest)
		return
	}

	updatedUser, err := c.db.UpdateWeightToUser(userID, weight)
	if err != nil {
		log.Println("Failed to find update user score")
		http.Error(w, "Failed to find update user score", http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(updatedUser)
	if err != nil {
		http.Error(w, "Could not marshal the user", http.StatusInternalServerError)
		return
	}

	writeResponseJson(w, userJson)
}
