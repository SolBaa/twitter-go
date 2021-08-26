package db

import (
	"context"
	"time"

	"github.com/SolBaa/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCT.Database("twittergo")
	col := db.Collection("relation")

	condicion := bson.M{
		"userID":    t.UserID,
		"userRelID": t.UserRelationID,
	}

	var result models.Relation
	err := col.FindOne(ctx, condicion).Decode(&result)
	if err != nil {
		return false, err
	}
	return true, nil
}
