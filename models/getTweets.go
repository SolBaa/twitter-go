package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	UserID  string             `bson:"user_id" json:"user_id"`
	Message string             `bson:"message" json:"message"`
	Date    time.Time          `bson:"date" json:"date"`
}
