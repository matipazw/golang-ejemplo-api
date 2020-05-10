package controllers

import (
	"api-messages/models"
	"api-messages/repositories"
	"api-messages/services"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

var GetMessages = func(c *gin.Context) {

	var client = GetClient(c)
	var order = c.Query("order")

	messages := repositories.GetMessages(client, order)

	c.JSON(200, messages)
}

var SaveMessage = func(c *gin.Context) {

	var m models.Message
	err := c.ShouldBindJSON(&m)

	if err == nil {
		m.Client = GetClient(c)
		m.Created = time.Now()
		message := repositories.SaveMessage(m)

		err := services.Notify(message)

		if err == nil {
			c.JSON(200, message)
			return
		}
	}

	c.JSON(500, "Internal Error Server")
}

var GetClientId = func(c *gin.Context) primitive.ObjectID {

	id, _ := c.Get("currentClientIdentity")
	client, _ := primitive.ObjectIDFromHex(id.(string))
	return client
}

var GetClient = func(c *gin.Context) models.Client {

	client := models.Client{}
	client.Id = GetClientId(c)
	return client
}
