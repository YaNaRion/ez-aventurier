package infra

import (
	"context"
	"fmt"
	"main/infra/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *DB) FindUser(userID string) (*models.User, error) {
	// Check cache with expiration
	db.Cache.Mu.RLock()
	cached, exists := db.Cache.Users[userID]
	db.Cache.Mu.RUnlock()

	if exists && time.Since(cached.timestamp) < 10*time.Minute {
		return cached.user, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database(db.CurrentDB).Collection("users")
	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}

	var user models.User
	err := collection.FindOne(ctx, bson.M{"userId": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user '%s' not found", userID)
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	db.Cache.Mu.Lock()
	db.Cache.Users[userID] = CachedUser{
		user:      &user,
		timestamp: time.Now(),
	}
	db.Cache.Mu.Unlock()

	return &user, nil
}

func (db *DB) AddUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database(db.CurrentDB).Collection("users")
	if collection == nil {
		return fmt.Errorf("failed to get collection")
	}
	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return fmt.Errorf("failed to insert the session in db: %w", err)
	}
	return nil
}

// Fonction pour ajouter tous les users dans la db
func (db *DB) AddUsers(users []models.User, dbstr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var docs []interface{}
	for _, user := range users {
		docs = append(docs, user)
	}

	collection := db.Client.Database(dbstr).Collection("users")
	if collection == nil {
		return fmt.Errorf("failed to get collection")
	}
	_, err := collection.InsertMany(ctx, docs)

	if err != nil {
		return fmt.Errorf("failed to insert the session in db: %w", err)
	}
	return nil
}

func (db *DB) UpdateWeightToUser(
	userID, cacheID string,
	position, addedWeight int,
) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database(db.CurrentDB).Collection("users")
	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}

	filter := bson.M{"userId": userID}

	newClaimedCache := models.ClaimedCache{
		CacheID:  cacheID,  // Your cache ID
		Position: position, // Your position value
	}

	update := bson.M{
		"$inc": bson.M{
			"score": addedWeight,
		},
		"$set": bson.M{
			"updatedAt": time.Now(),
		},
		"$push": bson.M{
			"claimedCache": newClaimedCache,
		},
	}

	var updatedUser models.User
	err := collection.FindOneAndUpdate(ctx, filter, update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user '%s' not found", userID)
		}
		return nil, fmt.Errorf("failed to update user weight: %w", err)
	}

	db.Cache.LeaderBoard.user = make([]models.User, 0)
	db.Cache.Mu.Lock()
	db.Cache.Users[userID] = CachedUser{
		user:      &updatedUser,
		timestamp: time.Now(),
	}
	db.Cache.Mu.Unlock()

	return &updatedUser, nil
}

func (db *DB) GetAllUserOrderByScoreDes() ([]models.User, error) {
	if len(db.Cache.LeaderBoard.user) > 0 &&
		time.Since(db.Cache.LeaderBoard.timestamp) < 1*time.Minute {
		return db.Cache.LeaderBoard.user, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database(db.CurrentDB).Collection("users")
	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}
	findOptions := options.Find()
	findOptions.SetSort(
		bson.D{{Key: "score", Value: -1}},
	)

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)

	if err != nil {
		return nil, fmt.Errorf("error finding users: %v", err)
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, fmt.Errorf("error decoding user: %v", err)
		}
		users = append(users, user)
	}

	// Check for cursor errors
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	db.Cache.Mu.Lock()
	db.Cache.LeaderBoard = CacheLeaderBoard{
		user:      users,
		timestamp: time.Now(),
	}
	db.Cache.Mu.Unlock()

	return users, nil
}
