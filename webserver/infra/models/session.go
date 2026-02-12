package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	UserID    string    `bson:"userID"    json:"userID"`
	Host      string    `bson:"host"      json:"host"`
	SessionID string    `bson:"sessionID" json:"sessionID"`
	CreatedON time.Time `bson:"createdOn" json:"createdOn"`
}

func NewSession(userID string, host string) Session {
	return Session{
		UserID:    userID,
		Host:      host,
		SessionID: uuid.New().String(),
		CreatedON: time.Now().UTC(),
	}
}
