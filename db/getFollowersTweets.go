package db

import (
	"context"
	"time"

	"github.com/SolBaa/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFollowersTweets(ID string, age int) ([]models.FollowersTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCT.Database("twittergo")
	col := db.Collection("relation")

	skip := (age - 1) * 20

	condition := make([]bson.M, 0)
	condition = append(condition, bson.M{"$match": bson.M{"usuarioid": ID}})
	condition = append(condition, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userRelID",
			"foreignField": "userID",
			"as":           "tweet",
		}})
	condition = append(condition, bson.M{"$unwind": "$tweet"})
	condition = append(condition, bson.M{"$sort": bson.M{"tweet.date": -1}})
	condition = append(condition, bson.M{"$skip": skip})
	condition = append(condition, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condition)
	var result []models.FollowersTweets
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
