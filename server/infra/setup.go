package infra

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const urlConnection = "mongodb+srv://YaNaRion:kolia1@dev.ddwky9s.mongodb.net/?appName=dev"

type DB struct {
	Ctx    context.Context
	Client mongo.Client
}

func Setup() (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	opts := options.Client().
		ApplyURI("mongodb+srv://YaNaRion:<db_password>@dev.ddwky9s.mongodb.net/?appName=dev").
		SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return nil, nil
}

func (db *DB) HelloWorld() {
	log.Println("hello from db")
}
