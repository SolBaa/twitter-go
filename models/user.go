package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User es el modelo de user de la db
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id" `
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	LastName  string             `bson:"last_name,omitempty" json:"last_name,omitempty"`
	BirthDate time.Time          `bson:"birth_date,omitempty" json:"birth_date,omitempty" `
	Email     string             `bson:"email,omitempty" json:"email"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	Avatar    string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
	Banner    string             `bson:"banner,omitempty" json:"banner,omitempty"`
	Biography string             `bson:"bio,omitempty" json:"bio,omitempty"`
	Location  string             `bson:"location,omitempty" json:"location,omitempty"`
	WebSite   string             `bson:"web_site,omitempty" json:"web_site,omitempty"`
}
