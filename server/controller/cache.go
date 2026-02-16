package controller

import (
	"encoding/json"
	"log"
	"main/class"
	"main/infra"
	"main/infra/models"
	"net/http"
	"strconv"
	"time"
)

type CreateCacheRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ReleaseTime string `json:"release_time"` // From frontend in Montreal time
}

func (c *Controller) postCache(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")

	_, _, err := c.isSessionValid(sessionID, r.URL.Host)
	if err != nil {
		log.Println("Session not valid")
		http.Error(w, "Session not valid", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Decode request
	var req CreateCacheRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Failed to decode request body", err)
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	// Log received Montreal time
	log.Printf("Received Montreal release time: %s", req.ReleaseTime)

	// Parse the release time as Montreal time
	montrealTime, err := class.ParseMontrealTime(req.ReleaseTime)
	if err != nil {
		log.Printf("Failed to parse Montreal release time '%s': %v", req.ReleaseTime, err)
		http.Error(
			w,
			"Invalid release time format. Expected format: YYYY-MM-DDTHH:mm",
			http.StatusBadRequest,
		)
		return
	}

	// Log parsed Montreal time
	if !montrealTime.IsZero() {
		log.Printf("Parsed Montreal time: %s", montrealTime.Format("2006-01-02 15:04:05 MST"))

		// Convert to UTC for storage
		utcTime := class.ToUTC(montrealTime)
		log.Printf("Converted to UTC for storage: %s", utcTime.Format("2006-01-02 15:04:05 MST"))
	}

	// Generate answer
	answer, err := infra.CustomID(8, infra.AlphaNumeric)
	if err != nil {
		log.Println("Failed to generate answer", err)
		http.Error(w, "Failed to generate answer", http.StatusInternalServerError)
		return
	}

	// Create cache object with UTC time
	cache := models.Cache{
		Name:         req.Name,
		Description:  req.Description,
		ReleaseTime:  class.ToUTC(montrealTime), // Store UTC in MongoDB
		Answers:      answer,
		Answer_count: 0,
		Weight:       14,
		CreatedAt:    time.Now().UTC(), // Always store timestamps in UTC
	}

	// Validate
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

	// Add to database
	err = c.db.AddCache(cache)
	if err != nil {
		log.Printf("Failed to add the cache to db: %v", err)
		http.Error(w, "Failed to add the cache to db", http.StatusInternalServerError)
		return
	}

	// Success response with Montreal time for display
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Cache created successfully",
		"cache": map[string]interface{}{
			"name":         cache.Name,
			"description":  cache.Description,
			"cache_number": cache.CacheNumber,
			"release_time": cache.GetFormattedReleaseTime(),      // Formatted for display
			"input_time":   cache.GetInputFormattedReleaseTime(), // For datetime-local input
		},
	})
}

func (c *Controller) getCaches(w http.ResponseWriter, r *http.Request) {
	caches, err := c.db.GetVisibleCaches()
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
