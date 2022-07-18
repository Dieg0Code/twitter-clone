package db

import (
	"context"
	"log"
	"time"

	"github.com/dieg0code/twitter-clone/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadTweets reads the tweets from a user by his ID
func ReadTweets(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter-clone")
	col := db.Collection("tweet")

	var result []*models.ReturnTweets

	condition := bson.M{"userid": ID}

	opt := options.Find()
	opt.SetLimit(20)
	opt.SetSort(bson.D{{Key: "date", Value: -1}}) // -1 para ordenar de forma descendente
	opt.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, opt)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var tweet models.ReturnTweets
		err := cursor.Decode(&tweet)
		if err != nil {
			return result, false
		}
		result = append(result, &tweet)
	}

	return result, true
}
