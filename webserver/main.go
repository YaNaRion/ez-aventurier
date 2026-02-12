package main

import (
	"log"
	"main/controller"
	"main/infra"
	"main/router"
	"net/http"
	"os"
	"strconv"
)

type AppConfig struct {
	Port        int
	DatabaseURL string
}

type Config struct {
	DB         *infra.DB
	Controller *controller.Controller
	Router     *router.Router
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

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
func getEnvAsInt(key string, defaultValue int) int {
	strValue := getEnv(key, "")
	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}
	return defaultValue
}

func Setup() *Server {
	mux := http.NewServeMux()
	// Setup HTTP request
	log.Println("Setup Web router")
	router := router.Setup(mux)

	configServer := NewConf(router)
	// Start server
	server := &http.Server{
		Addr: ":3000",
	}

	return NewServer(configServer, mux, server)
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
