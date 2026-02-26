package models

import (
	"main/class"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
	UserID      string    `bson:"user_id"      json:"user_id"`
	ClaimedTime time.Time `bson:"claimed_time" json:"claimed_time"`
}

type Cache struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Name         string             `bson:"name"          json:"name"`
	Description  string             `bson:"description"   json:"description"`
	CacheNumber  int64              `bson:"cacheNumber"   json:"cacheNumber"`
	CreatedAt    time.Time          `bson:"createdAt"     json:"createdAt"`
	Answers      string             `bson:"answer"        json:"answer"`
	Weight       int                `bson:"weight"        json:"weight"`
	Answer_count int                `bson:"answer_count"  json:"answer_count"`
	ReleaseTime  time.Time          `bson:"releaseTime"   json:"releaseTime"`
	ClaimBy      Claim              `bson:"claim_by"      json:"claim_by"`
}

// GetMontrealReleaseTime returns the release time in Montreal timezone
func (c *Cache) GetMontrealReleaseTime() time.Time {
	return class.FromUTC(c.ReleaseTime)
}

// GetFormattedReleaseTime returns the release time formatted for display
func (c *Cache) GetFormattedReleaseTime() string {
	return class.FormatMontrealTime(c.ReleaseTime)
}

// GetInputFormattedReleaseTime returns the release time formatted for datetime-local input
func (c *Cache) GetInputFormattedReleaseTime() string {
	return class.FormatMontrealTimeInput(c.ReleaseTime)
}
