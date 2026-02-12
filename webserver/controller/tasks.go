package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func (c *Controller) getTasks(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET TASKS REQUEST FROM: %s", r.RemoteAddr)
	tasks := Tasks{
		Tasks: c.tasks,
	}
	log.Println(tasks)
	tasksJson, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeResponseJson(w, tasksJson)
}

func (c *Controller) getTest(w http.ResponseWriter, r *http.Request) {
	log.Println("RECEIVE TASKS REQUEST FROM RUST CLIENT")
	writeResponseJson(w, []byte("OUI"))
}
