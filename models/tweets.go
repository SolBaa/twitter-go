package models

type Tweets struct {
	Message string `bson:"message" json:"message"`
}
