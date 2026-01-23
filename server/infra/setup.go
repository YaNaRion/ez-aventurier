package infra

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const urlConnection = "mongodb+srv://YaNaRion:kolia1@dev.ddwky9s.mongodb.net/?appName=dev"

type DB struct {
	Ctx    context.Context
	Client mongo.Client
}

func Setup() (*DB, error) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(urlConnection))
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := DB{
		Client: *client,
		Ctx:    ctx,
	}

	return &db, nil
}

func (db *DB) HelloWorld() {
	log.Println("hello from db")
}
