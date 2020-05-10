package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type To struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email"`
	Name  string             `json:"name" bson:"name"`
}
