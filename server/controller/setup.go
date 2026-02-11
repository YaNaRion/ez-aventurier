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
	controller.setUpRouter(mux)
	return controller
}

func (C *Controller) setUpRouter(mux *http.ServeMux) {
	// Test et template
	mux.HandleFunc("GET /tasks/", C.getTasks)
	mux.HandleFunc("GET /api/test", C.getTest)

	// Verification de la connection
	mux.HandleFunc("GET /api/isSessionValid", C.isSessionValidMiddle)

	// Demande de connection
	mux.HandleFunc("GET /api/login", C.connection)
	mux.HandleFunc("GET /api/user", C.getUser)

	mux.HandleFunc("GET /api/caches", C.getCaches)

	mux.HandleFunc("POST /api/cache", C.postCache)
}

func writeResponseJson(w http.ResponseWriter, data []byte) {
	w.Header().Set("Contend-Type", ContentTypeJSON)
	_, err := w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
