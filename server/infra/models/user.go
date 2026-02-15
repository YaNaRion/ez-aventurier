package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClaimedCache struct {
	CacheID  string `bson:"cache_id" json:"cache_id"`
	Position int    `bson:"position" json:"position"`
}

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name          string             `bson:"name"          json:"name"         validate:"required"`
	UserID        string             `bson:"userId"        json:"userId"       validate:"required,min=8,max=8"` // Your 8-character string ID
	Ordre         string             `bson:"order"         json:"order"        validate:"required"`             // Ordre
	Unity         string             `bson:"unity"         json:"unity"        validate:"required"`             // Unité scout
	CreatedAt     time.Time          `bson:"createdAt"     json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt"     json:"updatedAt"`
	Score         int                `bson:"score"         json:"score"`
	ClaimedCaches []ClaimedCache     `bson:"claimedCache"  json:"claimedCache"`
}

// CollectionName returns the MongoDB collection name
func (u *User) CollectionName() string {
	return "users"
}
func NewUser(name, userID, unity, order string) *User {
	return &User{
		Name:          name,
		UserID:        userID,
		Unity:         unity,
		Ordre:         order,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		ClaimedCaches: []ClaimedCache{},
	}
}
