package controller

import (
	"main/infra"
	"net/http"
)

const ContentTypeJSON = "application/json"

type Controller struct {
	db    *infra.DB
	mux   *http.ServeMux
	tasks []Task
}

func newController(db *infra.DB, mux *http.ServeMux) *Controller {
	return &Controller{
		db:    db,
		mux:   mux,
		tasks: GenerateTemplateTask(),
	}
}

func SetUpController(mux *http.ServeMux, db *infra.DB) *Controller {
	controller := newController(db, mux)
	// Test et template
	mux.HandleFunc("GET /tasks/", controller.getTasks)
	mux.HandleFunc("GET /api/test", controller.getTest)

	// Verification de la connection
	mux.HandleFunc("GET /api/isSessionValid", controller.isSessionValid)

	// Demande de connection
	mux.HandleFunc("GET /api/login", controller.connection)

	mux.HandleFunc("GET /api/user", controller.getUser)

	mux.HandleFunc("POST /api/cache", controller.getUser)

	return controller
}

func writeResponseJson(w http.ResponseWriter, data []byte) {
	w.Header().Set("Contend-Type", ContentTypeJSON)
	_, err := w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
