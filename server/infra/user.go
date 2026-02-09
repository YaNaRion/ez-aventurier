package infra

import (
	"context"
	"fmt"
	"main/infra/models"
	"time"
)

func (db *DB) AddUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("users")
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
	_, err := collection.InsertMany(ctx, docs)

	if err != nil {
		return fmt.Errorf("failed to insert the session in db: %w", err)
	}
	return nil
}
