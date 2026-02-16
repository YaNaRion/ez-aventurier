package infra

import (
	"context"
	"fmt"
	"log"
	"main/infra/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CacheStore struct {
	AllCaches     []models.Cache
	VisibleCaches []models.Cache
	LastUpdated   time.Time
}

type CacheDB struct {
	LeaderBoard []models.User
	CacheStore  CacheStore
	Users       map[string]*models.User
}

type DB struct {
	Ctx         context.Context
	Client      *mongo.Client
	IsConnected bool
	Cache       CacheDB
	CurrentDB   string
}

func (db *DB) SetDB(isDev bool) {
	if isDev {
		db.CurrentDB = "dev1"
	} else {
		db.CurrentDB = "prod"
	}
}

func Setup(dbConnectionString string, isDev bool) (*DB, error) {
	// Create a context with timeout for connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConnectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Verify connection with fresh context
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()

	if err := client.Ping(pingCtx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	// List databases to verify connection and see what's available
	dbCtx, dbCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer dbCancel()

	databases, err := client.ListDatabaseNames(dbCtx, bson.M{})
	if err != nil {
		log.Printf("Warning: Could not list databases: %v", err)
	} else {
		log.Printf("Available databases: %v", databases)
	}

	db := &DB{
		Client: client,
		Ctx:    context.Background(), // Base context without timeout
		Cache: CacheDB{
			LeaderBoard: make([]models.User, 0),
			Users:       make(map[string]*models.User),
			CacheStore: CacheStore{
				AllCaches:     make([]models.Cache, 0),
				VisibleCaches: make([]models.Cache, 0),
				LastUpdated:   time.Time{}, // Zero time value
			},
		},
	}
	db.SetDB(isDev)

	log.Println("✅ MongoDB connection established successfully")
	return db, nil
}
