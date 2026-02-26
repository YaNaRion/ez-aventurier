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

	collection := db.Client.Database(db.CurrentDB).Collection("caches")
	if collection == nil {
		return fmt.Errorf("failed to get collection")
	}

	// Set creation time
	cache.CreatedAt = time.Now()

	// Get current count for CacheNumber
	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Printf("Error while counting the number of caches in DB: %v", err)
		return fmt.Errorf("error while counting the number of caches in DB: %w", err)
	}
	cache.CacheNumber = count + 1

	// Insert the new cache
	_, err = collection.InsertOne(ctx, cache)
	if err != nil {
		log.Printf("Failed to insert cache in db: %v", err)
		return fmt.Errorf("failed to insert cache in db: %w", err)
	}

	// Invalidate the cache to force refresh on next read
	db.Cache.CacheStore.AllCaches = []models.Cache{}
	db.Cache.CacheStore.VisibleCaches = []models.Cache{}
	db.Cache.CacheStore.LastUpdated = time.Time{}

	log.Printf("Cache '%s' (ID: %d) created successfully with release time: %v",
		cache.Name, cache.CacheNumber, cache.ReleaseTime)

	return nil
}

func (db *DB) GetCache(cacheNumber int) (*models.Cache, error) {
	now := time.Now().UTC()

	if !db.Cache.CacheStore.LastUpdated.IsZero() &&
		now.Sub(db.Cache.CacheStore.LastUpdated) <= time.Minute {

		for _, cache := range db.Cache.CacheStore.VisibleCaches {
			if cache.CacheNumber == int64(cacheNumber) {
				return &cache, nil
			}
		}

		for _, cache := range db.Cache.CacheStore.AllCaches {
			if cache.CacheNumber == int64(cacheNumber) {
				log.Printf("Cache %d exists but is not visible (hidden in cache)", cacheNumber)
				return nil, fmt.Errorf("cache '%d' is not yet available", cacheNumber)
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database(db.CurrentDB).Collection("caches")
	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}

	filter := bson.M{
		"cacheNumber": cacheNumber,
		"$or": []bson.M{
			{"releaseTime": nil},
			{"releaseTime": time.Time{}},
			{"releaseTime": bson.M{"$lte": now}},
		},
	}

	var cache models.Cache
	err := collection.FindOne(ctx, filter).Decode(&cache)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Check if cache exists but is hidden
			var hiddenCache models.Cache
			errHidden := collection.FindOne(ctx, bson.M{"cacheNumber": cacheNumber}).
				Decode(&hiddenCache)
			if errHidden == nil {
				// Cache exists but is hidden
				log.Printf("Cache %d exists in DB but is hidden (ReleaseTime: %v)",
					cacheNumber, hiddenCache.ReleaseTime)
				return nil, fmt.Errorf("cache '%d' is not yet available", cacheNumber)
			}
			return nil, fmt.Errorf("cache '%d' not found", cacheNumber)
		}
		return nil, fmt.Errorf("failed to find cache: %w", err)
	}

	go db.refreshCacheIfNeeded()

	return &cache, nil
}

func filterVisibleCaches(caches []models.Cache, now time.Time) []models.Cache {
	visible := make([]models.Cache, 0)
	for _, cache := range caches {
		// Same visibility rule:
		// - ReleaseTime is zero (not set) → visible
		// - ReleaseTime is in the past or exactly now → visible
		// - ReleaseTime is in the future → hidden
		if cache.ReleaseTime.IsZero() || !cache.ReleaseTime.After(now) {
			visible = append(visible, cache)
		}
	}
	return visible
}

// Helper function to refresh the cache
func (db *DB) refreshCacheIfNeeded() {
	now := time.Now().UTC()

	// Check if cache needs refresh
	if db.Cache.CacheStore.LastUpdated.IsZero() ||
		now.Sub(db.Cache.CacheStore.LastUpdated) > time.Minute {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		collection := db.Client.Database(db.CurrentDB).Collection("caches")
		if collection == nil {
			log.Println("Failed to refresh cache: collection is nil")
			return
		}

		// Get all caches
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Printf("Failed to refresh cache: %v", err)
			return
		}
		defer cursor.Close(ctx)

		var allCaches []models.Cache
		if err = cursor.All(ctx, &allCaches); err != nil {
			log.Printf("Failed to decode caches for refresh: %v", err)
			return
		}

		// Update cache
		db.Cache.CacheStore.AllCaches = allCaches
		db.Cache.CacheStore.LastUpdated = now

		// Filter visible caches
		db.Cache.CacheStore.VisibleCaches = filterVisibleCaches(allCaches, now)
	}
}

func (db *DB) GetVisibleCaches() ([]models.Cache, error) {
	now := time.Now().UTC()

	// Check cache first
	if !db.Cache.CacheStore.LastUpdated.IsZero() &&
		now.Sub(db.Cache.CacheStore.LastUpdated) <= time.Minute {
		return db.Cache.CacheStore.VisibleCaches, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database(db.CurrentDB).Collection("caches")
	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}

	// MongoDB query for visible caches
	filter := bson.M{
		"$or": []bson.M{
			{"releaseTime": nil},
			{"releaseTime": time.Time{}},
			{"releaseTime": bson.M{"$lte": now}},
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get visible caches: %w", err)
	}
	defer cursor.Close(ctx)

	var caches []models.Cache
	if err = cursor.All(ctx, &caches); err != nil {
		return nil, fmt.Errorf("failed to decode caches: %w", err)
	}

	// Update cache
	db.Cache.CacheStore.VisibleCaches = caches
	db.Cache.CacheStore.LastUpdated = now

	return caches, nil
}

// ajouter logique des de deadline
func (db *DB) ClaimCaches(user_id, answer_id string) (*models.Cache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database(db.CurrentDB).Collection("caches")

	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}

	var cache models.Cache
	oneWeekAgo := time.Now().UTC().AddDate(0, 0, -7)
	err := collection.FindOneAndUpdate(
		ctx,
		bson.M{
			"answer": answer_id,
			"releaseTime": bson.M{
				"$gte": oneWeekAgo, // releaseTime should be greater than or equal to one week ago
			},
		},
		bson.M{
			"$inc": bson.M{"answer_count": 1},
			"$push": bson.M{"claimed_by": models.Claim{
				UserID:      user_id,
				ClaimedTime: time.Now().UTC(),
			}},
		},
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
