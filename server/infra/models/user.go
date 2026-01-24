package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"      json:"id,omitempty"`
	Name      string             `bson:"name"               json:"name"               validate:"required"`
	UserID    string             `bson:"userId"             json:"userId"             validate:"required,min=8,max=8"` // Your 8-character string ID
	Team      int                `bson:"team"               json:"team"               validate:"required"`             // Ordre
	Group     int                `bson:"group"              json:"group"              validate:"required"`             // Unité scout
	Username  *string            `bson:"username,omitempty" json:"username,omitempty"`
	CreatedAt time.Time          `bson:"createdAt"          json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"          json:"updatedAt"`
}

// CollectionName returns the MongoDB collection name
func (u *User) CollectionName() string {
	return "users"
}
