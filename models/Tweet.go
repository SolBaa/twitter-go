package models

import "time"

type Tweet struct {
	UserID  string    `bson:"userID" json:"userID,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
