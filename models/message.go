package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Client  Client             `json:"client" bson:"client"`
	Sender  Sender             `json:"sender"  bson:"sender"`
	Title   string             `json:"title" bson:"title"`
	Text    string             `json:"text"  bson:"text"`
	Created time.Time          `json:"created"  bson:"created"`
	To      []To               `json:"to"  bson:"to"`
}
