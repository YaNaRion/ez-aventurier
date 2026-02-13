package main

import (
	"flag"
	"fmt"
	"log"
	"main/controller"
	"main/infra"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type AppConfig struct {
	Port        int
	DatabaseURL string
}

type Config struct {
	DB         *infra.DB
	Controller *controller.Controller
}

func NewConf(
	db *infra.DB,
	controller *controller.Controller,
) *Config {
	return &Config{DB: db, Controller: controller}
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

func LoadEnvVariable() AppConfig {
	env := flag.String("env", "prod", "environment to run in")

	flag.Parse()

	if *env == "dev" {
		godotenv.Load()
	}

	return AppConfig{
		DatabaseURL: getEnv("DB_CONNECTION_STRING", "None"),
		Port:        getEnvAsInt("PORT", 3000),
	}
}

func Setup() *Server {
	log.Println("Setup DB connection")
	var db *infra.DB

	appConfig := LoadEnvVariable()
	log.Println(appConfig)
	db, _ = infra.Setup(appConfig.DatabaseURL)
	if db == nil {
		log.Println("DB NOT CONNECTED")
	} else {
		log.Println("DB CONNECTED")
	}

	// Setup des routes de l'API
	log.Println("Setup Http controller")
	mux := http.NewServeMux()
	control := controller.SetUpController(mux, db)

	configServer := NewConf(db, control)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{
			"*",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: false,
		Debug:            true,
	})

	// Start server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", appConfig.Port),
		Handler: corsHandler.Handler(mux),
	}

	return NewServer(configServer, mux, corsHandler, server)
}

// To reset db
// infra.Add_jeune_to_DB(server.Conf.DB)
func main() {

	// Setup DB connection
	server := Setup()

	log.Println("Listen on localhost:3000")
	err := server.httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
