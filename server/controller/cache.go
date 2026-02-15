package controller

import (
	"encoding/json"
	"io"
	"log"
	"main/infra"
	"main/infra/models"
	"net/http"
	"strconv"
	"time"
)

func (c *Controller) postCache(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")

	_, _, err := c.isSessionValid(sessionID, r.URL.Host)
	if err != nil {
		log.Println("Session not valid")
		http.Error(w, "Session not valid", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read the body: %s", err)
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	var cache models.Cache
	err = json.Unmarshal(bodyBytes, &cache)
	if err != nil {
		log.Println("Failed to unmashal the body", err)
		http.Error(w, "Failed to unmarshal body", http.StatusInternalServerError)
		return
	}

	cache.Answers, err = infra.CustomID(8, infra.AlphaNumeric)
	cache.Answer_count = 0
	cache.Weight = 14
	cache.CreatedAt = time.Now()
	if cache.Name == "" {
		log.Println("Cache name is empty")
		http.Error(w, "Cache name is empty", http.StatusBadRequest)
		return
	}

	if cache.Description == "" {
		log.Println("Cache description is empty")
		http.Error(w, "Cache description is empty", http.StatusBadRequest)
		return
	}

	err = c.db.AddCache(cache)
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
