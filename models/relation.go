package models

type Relation struct {
	UserID         string `bson:"userID" json:"USerID"`
	UserRelationID string `bson:"userRelID,omitempty" json:"userRelID"`
}

type RelationResponse struct {
	Status bool `json:"status"`
}
