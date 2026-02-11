package controller

import (
	"log"
	"net/http"
)

func (c *Controller) postCache(w http.ResponseWriter, r *http.Request) {
	session_id := r.URL.Query().Get("session_id")

	session, err := c.db.FindSession(session_id)
	if err != nil {
		log.Println("Failed to find session")
		http.Error(w, "Failed to find session", http.StatusBadRequest)
		return
	}

	isSessionValid, err := c.isSessionValid(session, r.URL.Host)
	if err != nil && !isSessionValid {
		log.Println("Session not valid")
		http.Error(w, "Session not valid", http.StatusBadRequest)
		return
	}

	cache_text := r.URL.Query().Get("cache_txt")
	if cache_text == "" {
		log.Println("Cache_text is empty")
		http.Error(w, "Cache_text is empty", http.StatusBadRequest)
		return
	}

	err = c.db.AddCache(cache_text)
	if err != nil {
		http.Error(w, "Failed to add the cache to db", http.StatusBadRequest)
		return
	}

}
