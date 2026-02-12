package main

import (
	"log"
	"main/router"
	"net/http"
)

type AppConfig struct {
	Port        int
	DatabaseURL string
}

type Config struct {
	Router *router.Router
}

func NewConf(
	router *router.Router,
) *Config {
	return &Config{Router: router}
}

type Server struct {
	Conf       *Config
	mux        *http.ServeMux
	httpServer *http.Server
}

func NewServer(
	config *Config,
	mux *http.ServeMux,
	httpServer *http.Server,
) *Server {
	return &Server{
		Conf:       config,
		mux:        mux,
		httpServer: httpServer,
	}
}

func Setup() *Server {
	mux := http.NewServeMux()
	// Setup HTTP request
	log.Println("Setup Web router")
	router := router.Setup(mux)

	configServer := NewConf(router)
	// Start server
	server := &http.Server{
		Addr:    ":6969",
		Handler: mux, // Add this line!
	}

	return NewServer(configServer, mux, server)
}

func main() {
	// Setup DB connection
	server := Setup()
	err := server.httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
