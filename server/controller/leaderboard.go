package controller

import (
	"encoding/json"
	"net/http"
)

func (c *Controller) getLeaderboard(w http.ResponseWriter, r *http.Request) {
	leaderboard, err := c.db.GetAllUserOrderByScoreDes()
	if err != nil {
		LOG_ERROR_TO_CONSOLE("Failed to get leaderboard", r)
		http.Error(w, ERROR_CONTACT_TECH_SUPPORT, http.StatusInternalServerError)
		return
	}

	leaderBoard, err := json.Marshal(leaderboard)
	if err != nil {
		LOG_ERROR_TO_CONSOLE("Failed to get marshal the leaderboard", r)
		http.Error(w, ERROR_CONTACT_TECH_SUPPORT, http.StatusInternalServerError)
		return
	}

	writeResponseJson(w, leaderBoard)
}
