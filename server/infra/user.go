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

func (db *DB) FindUser(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("users")
	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}

	var user models.User
	err := collection.FindOne(ctx, bson.M{"userId": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user '%s' not found", username)
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	user.Score = 10
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

func (db *DB) UpdateWeightToUser(userID string, addedWeight int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("users")
	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}

	filter := bson.M{"userId": userID}

	update := bson.M{
		"$inc": bson.M{
			"weight": addedWeight,
		},
		"$set": bson.M{
			"updatedAt": time.Now(),
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

	return &updatedUser, nil
}
