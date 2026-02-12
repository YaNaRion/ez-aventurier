package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Participant struct {
	UserID   string    `bson:"userId"   json:"userId"   validate:"required"`
	Name     string    `bson:"name"     json:"name"`
	JoinedAt time.Time `bson:"joinedAt" json:"joinedAt"`
}

type LastMessage struct {
	Text      string    `bson:"text,omitempty"      json:"text,omitempty"`
	SenderID  string    `bson:"senderId,omitempty"  json:"senderId,omitempty"`
	Timestamp time.Time `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
}

type GroupInfo struct {
	Name        string   `bson:"name,omitempty"        json:"name,omitempty"`
	Description string   `bson:"description,omitempty" json:"description,omitempty"`
	Admin       string   `bson:"admin,omitempty"       json:"admin,omitempty"`
	Members     []string `bson:"members,omitempty"     json:"members,omitempty"`
}

type Conversation struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"              json:"id,omitempty"`
	Participants     []Participant      `bson:"participants"               json:"participants"               validate:"required,min=2"`
	ConversationType string             `bson:"conversationType"           json:"conversationType"           validate:"oneof=direct group"`
	ConversationName *string            `bson:"conversationName,omitempty" json:"conversationName,omitempty"`
	CreatedBy        string             `bson:"createdBy"                  json:"createdBy"                  validate:"required"`
	LastMessage      *LastMessage       `bson:"lastMessage,omitempty"      json:"lastMessage,omitempty"`
	UnreadCount      map[string]int     `bson:"unreadCount"                json:"unreadCount"`
	IsGroup          bool               `bson:"isGroup"                    json:"isGroup"`
	GroupInfo        *GroupInfo         `bson:"groupInfo,omitempty"        json:"groupInfo,omitempty"`
	CreatedAt        time.Time          `bson:"createdAt"                  json:"createdAt"`
	UpdatedAt        time.Time          `bson:"updatedAt"                  json:"updatedAt"`
}

// CollectionName returns the MongoDB collection name
func (c *Conversation) CollectionName() string {
	return "conversations"
}

// BeforeCreate hook to set timestamps and defaults
func (c *Conversation) BeforeCreate() {
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now

	// Set joinedAt for participants
	for i := range c.Participants {
		if c.Participants[i].JoinedAt.IsZero() {
			c.Participants[i].JoinedAt = now
		}
	}

	// Initialize unreadCount map
	if c.UnreadCount == nil {
		c.UnreadCount = make(map[string]int)
	}

	// Set conversation type based on participants count
	if len(c.Participants) > 2 || c.IsGroup {
		c.ConversationType = "group"
		c.IsGroup = true
	} else {
		c.ConversationType = "direct"
		c.IsGroup = false
	}
}

// BeforeUpdate hook to update timestamp
func (c *Conversation) BeforeUpdate() {
	c.UpdatedAt = time.Now()
}

// UpdateLastMessage updates the last message in the conversation
func (c *Conversation) UpdateLastMessage(text, senderID string) {
	now := time.Now()
	c.LastMessage = &LastMessage{
		Text:      text,
		SenderID:  senderID,
		Timestamp: now,
	}
	c.UpdatedAt = now
}

// IncrementUnreadCount increments unread count for all participants except sender
func (c *Conversation) IncrementUnreadCount(senderID string) {
	for _, participant := range c.Participants {
		if participant.UserID != senderID {
			c.UnreadCount[participant.UserID]++
		}
	}
}

// ResetUnreadCount resets unread count for a user
func (c *Conversation) ResetUnreadCount(userID string) {
	c.UnreadCount[userID] = 0
}
