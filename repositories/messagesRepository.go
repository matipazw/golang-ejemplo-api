package repositories

import (
	"api-messages/models"
	messagesContext "api-messages/repositories/messagesContext"

	"go.mongodb.org/mongo-driver/bson"
)

var GetMessages = func(client models.Client, order string) []*models.Message {

	c := messagesContext.New()

	filter := bson.M{"client._id": client.Id}

	sort := bson.D{{"created", 1}}

	if order == "desc" {
		sort = bson.D{{"created", -1}}
	}

	return messagesContext.Find(c, filter, sort)
}

var SaveMessage = func(message models.Message) *models.Message {

	c := messagesContext.New()
	id, err := messagesContext.Save(c, message)

	if err == nil {

		filter := bson.M{"_id": id}
		return messagesContext.FindOne(c, filter)
	}

	return nil
}
