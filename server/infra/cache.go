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

func (db *DB) AddCache(cache models.Cache) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("caches")

	if collection == nil {
		return fmt.Errorf("failed to get collection")
	}

	var err error
	cache.CacheNumber, err = collection.CountDocuments(ctx, bson.D{})
	cache.CacheNumber++
	if err != nil {
		log.Println("Error while counting the number of cache in DB")
		return fmt.Errorf("Error while count:ing the number of cache in DB: %e", err)
	}

	if err != nil {
		log.Println("error while giving an answer to the new cache")
		return fmt.Errorf("error while giving an answer to the new cache")
	}

	_, err = collection.InsertOne(ctx, cache)
	if err != nil {
		return fmt.Errorf("failed to insert the session in db: %w", err)
	}
	return nil
}

func (db *DB) GetCache(cacheNumber int) (*models.Cache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("caches")

	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}

	var cache models.Cache
	err := collection.FindOne(ctx, bson.M{"cacheNumber": cacheNumber}).Decode(&cache)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("cache '%d' not found", cacheNumber)
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return &cache, nil
}

func (db *DB) GetCaches() ([]models.Cache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("caches")
	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}

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

// ajouter logique des de deadline
func (db *DB) ClaimCaches(user_id, answer_id string) (*models.Cache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("caches")

	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}

	var cache models.Cache
	err := collection.FindOneAndUpdate(
		ctx,
		bson.M{"answer": answer_id},
		bson.M{"$inc": bson.M{"answer_count": 1}},
		options.FindOneAndUpdate().SetReturnDocument(options.After), // Returns the updated document
	).Decode(&cache)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("cache '%d' not found", answer_id)
		}
		return nil, fmt.Errorf("failed to find cache: %w", err)
	}
	return &cache, nil
}
