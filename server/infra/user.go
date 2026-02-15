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
	foundUser := db.Cache.Users[userID]
	if foundUser != nil {
		return foundUser, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("users")
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

	db.Cache.Users[userID] = &user

	return &user, nil
}

func (db *DB) AddUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("users")
	if collection == nil {
		return fmt.Errorf("failed to get collection")
	}
	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return fmt.Errorf("failed to insert the session in db: %w", err)
	}
	return nil
}

func (db *DB) AddUsers(users []models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var docs []interface{}
	for _, user := range users {
		docs = append(docs, user)
	}

	collection := db.Client.Database("dev1").Collection("users")
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

	collection := db.Client.Database("dev1").Collection("users")
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

	db.Cache.LeaderBoard = nil
	db.Cache.Users[updatedUser.UserID] = &updatedUser

	return &updatedUser, nil
}

func (db *DB) GetAllUserOrderByScoreDes() ([]models.User, error) {
	if len(db.Cache.LeaderBoard) > 0 {
		return db.Cache.LeaderBoard, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("users")
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

	// Slice to store the results
	var users []models.User

	// Iterate through the cursor
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
	db.Cache.LeaderBoard = users

	return users, nil
}
