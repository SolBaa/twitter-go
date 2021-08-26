package db

import (
	"context"
	"log"
	"time"

	"github.com/SolBaa/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(ID string, page int64) ([]*models.GetTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCT.Database("twittergo")
	col := db.Collection("tweet")

	var result []*models.GetTweets

	condition := bson.M{"userID": ID}

	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var register models.GetTweets
		err := cursor.Decode(&register)
		if err != nil {
			return nil, false
		}

		result = append(result, &register)
	}
	return result, true

}
