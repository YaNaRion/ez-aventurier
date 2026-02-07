package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	UserID    string    `bson:"userID"    json:"userID"`
	SessionID string    `bson:"sessionID" json:"sessionID"`
	CreatedON time.Time `bson:"createdOn" json:"createdOn"`
}

func NewSession(userID string) Session {
	return Session{
		UserID:    userID,
		SessionID: uuid.New().String(),
		CreatedON: time.Now().UTC(),
	}
}
