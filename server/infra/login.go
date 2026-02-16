package infra

import (
	"context"
	"fmt"
	"main/infra/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DB) AddSession(userID string, urlHost string) (*models.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database(db.CurrentDB).Collection("session")
	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}

	newSession := models.NewSession(userID, urlHost)

	_, err := collection.InsertOne(ctx, newSession)

	if err != nil {
		return nil, fmt.Errorf("failed to insert the session in db: %w", err)
	}

	return &newSession, nil
}

func (db *DB) FindSession(sessionID string) (*models.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.Client.Database(db.CurrentDB).Collection("session")

	if collection == nil {
		return nil, fmt.Errorf("failed to get collection")
	}

	var session models.Session
	err := collection.FindOne(ctx, bson.M{"sessionID": sessionID}).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("session '%s' not found", sessionID)
		}
		return nil, fmt.Errorf("failed to find session: %w", err)
	}

	return &session, nil
}
