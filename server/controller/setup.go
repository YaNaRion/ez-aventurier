package controller

import (
	"main/infra"
	"net/http"
)

const ContentTypeJSON = "application/json"

type Controller struct {
	db  *infra.DB
	mux *http.ServeMux
}

func newController(db *infra.DB, mux *http.ServeMux) *Controller {
	return &Controller{
		db:  db,
		mux: mux,
	}
}

func SetUpController(mux *http.ServeMux, db *infra.DB) *Controller {
	controller := newController(db, mux)
	controller.setUpRouter(mux)
	return controller
}

func (C *Controller) setUpRouter(mux *http.ServeMux) {
	// Demande de connection
	mux.HandleFunc("GET /login", C.connection)
	mux.HandleFunc("GET /user", C.getUser)

	mux.HandleFunc("GET /caches", C.getCaches)

	mux.HandleFunc("GET /cache", C.getCache)

	mux.HandleFunc("POST /cache", C.postCache)

	mux.HandleFunc("PUT /claimCache", C.claimCache)
}

func writeResponseJson(w http.ResponseWriter, data []byte) {
	w.Header().Set("Contend-Type", ContentTypeJSON)
	_, err := w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
