package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweets struct {
	Message string `bson:"message" json:"message"`
}

type FollowersTweets struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userID" json:"userId,omitempty"`
	UserRelationID string             `bson:"userRelID" json:"userRelationId,omitempty"`
	Tweet          struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
