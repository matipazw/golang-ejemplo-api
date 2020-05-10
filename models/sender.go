package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sender struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"username"  bson:"username"`
	Name     string             `json:"name"  bson:"name"`
}
