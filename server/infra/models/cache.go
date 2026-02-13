package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Answer struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	CacheAnswerID string             `bson:"answer_id"     json:"answer_id"`
	IsAvailable   bool               `bson:"is_available"  json:"is_available"`
	CreatedAt     time.Time          `bson:"createdAt"     json:"createdAt"`
	Weight        int                `bson:"weight"        json:"weight"`
	AnwerBy       string             `bson:"answer_by"     json:"answer_by"`
}

type Cache struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Text        string             `bson:"text"          json:"text"`
	CacheNumber int                `bson:"cacheNumber"   json:"cacheNumber"`
	CreatedAt   time.Time          `bson:"createdAt"     json:"createdAt"`
	Answers     []string           `bson:"answer"        json:"answer"`
	Weight      int                `bson:"weight"        json:"weight"`
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
