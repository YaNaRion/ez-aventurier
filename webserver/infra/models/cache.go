package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cache struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Text        string             `bson:"text"          json:"text"`
	CacheNumber int                `bson:"cacheNumber"   json:"cacheNumber"`
	CreatedAt   time.Time          `bson:"createdAt"     json:"createdAt"`
}

// CollectionName returns the MongoDB collection name
func (u *Cache) CollectionName() string {
	return "caches"
}

func NewCache(text string, cacheNumber int) *Cache {
	return &Cache{
		Text:        text,
		CacheNumber: cacheNumber,
		CreatedAt:   time.Now(),
	}
}
