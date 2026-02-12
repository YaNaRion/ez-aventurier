package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
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

	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	var cache_text = string(bodyBytes)
	log.Println(cache_text)
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

func (c *Controller) getCaches(w http.ResponseWriter, r *http.Request) {
	log.Println("DANS GES CACHES")
	caches, err := c.db.GetCaches()
	if err != nil {
		http.Error(w, "Failed to get the caches to db", http.StatusBadRequest)
		return
	}

	cachesJson, err := json.Marshal(caches)
	if err != nil {
		http.Error(w, "Could not marshal the response", http.StatusForbidden)
		return
	}

	writeResponseJson(w, cachesJson)
}

func (c *Controller) getCache(w http.ResponseWriter, r *http.Request) {
	cacheNumber := r.URL.Query().Get("cache_number")
	log.Println(cacheNumber)

	cacheNumberInt, err := strconv.Atoi(cacheNumber)
	if err != nil {
		http.Error(w, "param is not an int", http.StatusBadRequest)
		return
	}

	caches, err := c.db.GetCache(cacheNumberInt)
	if err != nil {
		http.Error(w, "Failed to get the caches to db", http.StatusBadRequest)
		return
	}

	cachesJson, err := json.Marshal(caches)
	if err != nil {
		http.Error(w, "Could not marshal the response", http.StatusForbidden)
		return
	}

	writeResponseJson(w, cachesJson)
}
