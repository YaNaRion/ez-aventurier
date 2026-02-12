package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageSender struct {
	UserID string `bson:"userId" json:"userId" validate:"required"`
	Name   string `bson:"name"   json:"name"`
	Team   int    `bson:"team"   json:"team"`
	Group  int    `bson:"group"  json:"group"`
}

type Reaction struct {
	UserID    string    `bson:"userId"    json:"userId"`
	Emoji     string    `bson:"emoji"     json:"emoji"`
	ReactedAt time.Time `bson:"reactedAt" json:"reactedAt"`
}

type Message struct {
	ID             primitive.ObjectID  `bson:"_id,omitempty"     json:"id,omitempty"`
	ConversationID primitive.ObjectID  `bson:"conversationId"    json:"conversationId"    validate:"required"`
	Sender         MessageSender       `bson:"sender"            json:"sender"            validate:"required"`
	Text           *string             `bson:"text,omitempty"    json:"text,omitempty"`
	MessageType    string              `bson:"messageType"       json:"messageType"       validate:"oneof=text image file system"`
	Reactions      []Reaction          `bson:"reactions"         json:"reactions"`
	Edited         bool                `bson:"edited"            json:"edited"`
	Deleted        bool                `bson:"deleted"           json:"deleted"`
	DeletedFor     []string            `bson:"deletedFor"        json:"deletedFor"`
	ReplyTo        *primitive.ObjectID `bson:"replyTo,omitempty" json:"replyTo,omitempty"`
	CreatedAt      time.Time           `bson:"createdAt"         json:"createdAt"`
	UpdatedAt      time.Time           `bson:"updatedAt"         json:"updatedAt"`
}

// CollectionName returns the MongoDB collection name
func (m *Message) CollectionName() string {
	return "messages"
}

// BeforeCreate hook to set timestamps
func (m *Message) BeforeCreate() {
	now := time.Now()
	m.CreatedAt = now
	m.UpdatedAt = now

	if m.MessageType == "" {
		m.MessageType = "text"
	}
}

// BeforeUpdate hook to update timestamp
func (m *Message) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
