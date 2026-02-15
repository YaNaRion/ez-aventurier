package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cache struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Name         string             `bson:"name"          json:"name"`
	Description  string             `bson:"description"   json:"description"`
	CacheNumber  int64              `bson:"cacheNumber"   json:"cacheNumber"`
	CreatedAt    time.Time          `bson:"createdAt"     json:"createdAt"`
	Answers      string             `bson:"answer"        json:"answer"`
	Weight       int                `bson:"weight"        json:"weight"`
	Answer_count int                `bson:"answer_count"  json:"answer_count"`
}

// CollectionName returns the MongoDB collection name
func (u *Cache) CollectionName() string {
	return "caches"
}
