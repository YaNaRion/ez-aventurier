package main

import (
	"log"
	"main/controller"
	"main/infra"
	"main/router"
	"net/http"

	"github.com/rs/cors"
)

const (
	httpPort = ":3000"
)

type Config struct {
	DB         *infra.DB
	Controller *controller.Controller
	Router     *router.Router
}

func NewConf(
	db *infra.DB,
	controller *controller.Controller,
	router *router.Router,
) *Config {
	return &Config{DB: db, Controller: controller, Router: router}
}

type Server struct {
	Conf       *Config
	mux        *http.ServeMux
	cors       *cors.Cors
	httpServer *http.Server
}

func NewServer(
	config *Config,
	mux *http.ServeMux,
	cors *cors.Cors,
	httpServer *http.Server,
) *Server {
	return &Server{
		Conf:       config,
		mux:        mux,
		cors:       cors,
		httpServer: httpServer,
	}
}

func Setup() *Server {
	log.Println("Setup DB connection")
	var db *infra.DB

	db, _ = infra.Setup()
	if db == nil {
		log.Println("DB NOT CONNECTED")
	} else {
		log.Println("DB CONNECTED")
	}

	// Setup des routes de l'API
	log.Println("Setup Http controller")
	mux := http.NewServeMux()
	control := controller.SetUpController(mux, db)

	// Setup HTTP request
	log.Println("Setup Web router")
	router := router.Setup(mux)

	configServer := NewConf(db, control, router)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})

	// Start server
	server := &http.Server{
		Addr:    httpPort,
		Handler: corsHandler.Handler(mux),
	}

	return NewServer(configServer, mux, corsHandler, server)
}

func main() {
	// Setup DB connection
	server := Setup()

	log.Println("Listen on localhost:3000")
	err := server.httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
