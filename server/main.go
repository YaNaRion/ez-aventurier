package main

import (
	"github.com/rs/cors"
	"log"
	"main/controller"
	"main/gateway"
	"main/infra"
	"main/router"
	"net/http"
)

const (
	httpPort = ":3000"
)

type Config struct {
	Router     *router.Router
	DB         *infra.DB
	Controller *controller.Controller
	Gateway    *gateway.WebSocketHandler
}

func NewConf(
	db *infra.DB,
	router *router.Router,
	controller *controller.Controller,
	gateway *gateway.WebSocketHandler,
) *Config {
	return &Config{DB: db, Router: router}
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

	// Setup Websocket
	wsHandler := gateway.NewWebSocketHandler(db)
	mux.Handle("/ws", wsHandler)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(
			[]byte(`{"status":"ok","clients":` + string(wsHandler.GetClientCount()) + `}`),
		)
		if err != nil {
			log.Println("health check bug")
		}
	})

	// Setup HTTP request
	log.Println("Setup Web router")
	router := router.Setup(mux)

	configServer := NewConf(db, router, control, wsHandler)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
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
