package db

import (
	"context"
	"time"

	"github.com/dieg0code/twitter-clone/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckIfUserExists checks if a user exists in the database by the email given
func CheckIfUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter-clone")
	col := db.Collection("users")

	condition := bson.M{"email": email}
	var result models.User
	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
