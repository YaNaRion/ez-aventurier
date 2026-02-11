package infra

import (
	"context"
	"fmt"
	"log"
	"main/infra/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) AddCache(cache_text string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("caches")
	cache_number, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Println("Error while counting the number of cache in DB")
		return fmt.Errorf("Error while counting the number of cache in DB: %e", err)
	}

	cache := models.NewCache(cache_text, int(cache_number+1))
	_, err = collection.InsertOne(ctx, cache)
	if err != nil {
		return fmt.Errorf("failed to insert the session in db: %w", err)
	}
	return nil
}

func (db *DB) GetCaches() ([]models.Cache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("caches")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all caches from db: %w", err)
	}
	defer cursor.Close(ctx)

	var caches []models.Cache
	if err = cursor.All(ctx, &caches); err != nil {
		return nil, fmt.Errorf("failed to decode caches: %w", err)
	}

	return caches, nil
}
