package controller

import (
	"encoding/json"
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

	userID := r.URL.Query().Get("user_id")
	answerID := r.URL.Query().Get("answer_id")
	if userID == "" || answerID == "" {
		http.Error(w, "Failed to get query param", http.StatusBadRequest)
		return
	}

	user, err := c.db.FindUser(userID)
	if err != nil {
		http.Error(w, "Failed to find user", http.StatusInternalServerError)
		return
	}

	for _, claimedCache := range user.ClaimedCaches {
		if claimedCache.CacheID == answerID {
			http.Error(w, "Vous avez déjà eu cette cache", http.StatusInternalServerError)
			return
		}
	}

	cache, err := c.db.ClaimCaches(userID, answerID)
	if err != nil {
		http.Error(w, "Failed to claim cache", http.StatusInternalServerError)
		return
	}

	mulFactor := 1
	if cache.Answer_count <= 5 {
		mulFactor = 2
	} else if cache.Answer_count <= 10 {
		mulFactor = 3
	}
	userAddedPoint := cache.Weight * mulFactor

	user, err = c.db.UpdateWeightToUser(userID, cache.Answers, cache.Answer_count, userAddedPoint)
	if err != nil {
		http.Error(w, "Couldnt update player weight", http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Could not marshal the response", http.StatusForbidden)
		return
	}

	writeResponseJson(w, userJson)
}
