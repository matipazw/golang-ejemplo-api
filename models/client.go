package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	Id primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
}
