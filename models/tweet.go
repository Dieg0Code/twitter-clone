package models

// Tweet is the model for the tweet collection
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
