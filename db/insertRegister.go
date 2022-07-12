package db

import (
	"context"
	"time"

	"github.com/dieg0code/twitter-clone/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertRegister insert a new user into the database
func InserRegister(usr models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter-clone")
	col := db.Collection("users")

	usr.Password, _ = EncryptPassword(usr.Password)

	result, err := col.InsertOne(ctx, usr)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
