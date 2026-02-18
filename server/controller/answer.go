package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func LOG_ERROR_TO_CONSOLE(message string, r *http.Request) {
	log.Printf("Request from: %s failed because %s\n", r.URL, message)
}

const ERROR_CONTACT_TECH_SUPPORT = "Une erreur est survenue, contactez le camp.aventurier.229@gmail.com"

func (c *Controller) claimCache(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	if sessionID == "" {
		LOG_ERROR_TO_CONSOLE("Missing session_id in query param", r)
		http.Error(w, ERROR_CONTACT_TECH_SUPPORT, http.StatusBadRequest)
		return
	}

	isSessionValid, _, err := c.isSessionValid(sessionID, r)
	if err != nil && !isSessionValid {
		LOG_ERROR_TO_CONSOLE("Session is not valid anymore", r)
		http.Error(w, "Session not valid", http.StatusGatewayTimeout)
		return
	}

	userID := r.URL.Query().Get("user_id")
	answerID := r.URL.Query().Get("answer_id")
	if userID == "" || answerID == "" {
		LOG_ERROR_TO_CONSOLE("Missing user_id or answer_id in query param", r)
		http.Error(w, ERROR_CONTACT_TECH_SUPPORT, http.StatusBadRequest)
		return
	}

	user, err := c.db.FindUser(userID)
	if err != nil {
		LOG_ERROR_TO_CONSOLE(fmt.Sprintf("Could not find user: %s in DB", userID), r)
		http.Error(w, ERROR_CONTACT_TECH_SUPPORT, http.StatusInternalServerError)
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
		LOG_ERROR_TO_CONSOLE(fmt.Sprintf("Could not claim cache: %s ", err), r)
		http.Error(w, ERROR_CONTACT_TECH_SUPPORT, http.StatusInternalServerError)
		return
	}

	log.Println(cache.Answer_count)
	log.Println()
	mulFactor := 1
	if cache.Answer_count <= 5 {
		mulFactor = 2
	} else if cache.Answer_count <= 10 {
		mulFactor = 3
	}
	userAddedPoint := cache.Weight * mulFactor

	user, err = c.db.UpdateWeightToUser(userID, cache.Answers, cache.Answer_count, userAddedPoint)
	if err != nil {
		LOG_ERROR_TO_CONSOLE(fmt.Sprintf("an error happen when updating player score: %s", err), r)
		http.Error(w, "Couldnt not update player weight", http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		LOG_ERROR_TO_CONSOLE(fmt.Sprintf("could not marshal user: %s because: %s", userID, err), r)
		http.Error(
			w,
			"Votre cache a bien été enregistrer, mais une erreur est survenue, faite F5",
			http.StatusInternalServerError,
		)
		return
	}

	writeResponseJson(w, userJson)
}
