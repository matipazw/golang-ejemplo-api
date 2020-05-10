package messagesCollection

import (
	"api-messages/config"
	"api-messages/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessagesCollection struct {
	Collection *mongo.Collection
	Context    context.Context
}

func New() MessagesCollection {

	var database, collection string = "dbMessages", "messages"
	var uri = config.Provider().Database.Uri

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return MessagesCollection{client.Database(database).Collection(collection), ctx}
}

func Find(c MessagesCollection, filter bson.M, sort bson.D) (messages []*models.Message) {

	findOptions := options.Find()
	findOptions.SetSort(sort)

	cursor, err := c.Collection.Find(c.Context, filter, findOptions)
	defer cursor.Close(c.Context)

	if err != nil && err.Error() == "mongo: no documents in result" {
		return nil
	}

	for cursor.Next(c.Context) {
		var message models.Message

		err = cursor.Decode(&message)

		if err == nil {
			messages = append(messages, &message)
		}
	}

	return
}

func Save(c MessagesCollection, message models.Message) (interface{}, error) {

	message.Id = primitive.NewObjectID()
	insertRet, err := c.Collection.InsertOne(c.Context, message)

	return insertRet.InsertedID, err
}

func FindOne(c MessagesCollection, filter bson.M) *models.Message {

	var m models.Message

	err := c.Collection.FindOne(c.Context, filter).Decode(&m)

	if err != nil && err.Error() == "mongo: no documents in result" {

		return nil
	}

	return &m
}
