package models

import "time"

// CreateTweet is the model for the given tweet
type CreateTweet struct {
	UserId    string    `bson:"userid" json:"userid,omitempty"`
	Message   string    `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time `bson:"createdat" json:"createdat,omitempty"`
}
