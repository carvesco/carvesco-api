package controllers

import (
	"carvescoAPI/configs"
	"carvescoAPI/models"
	"carvescoAPI/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var serendipiaCollection *mongo.Collection = configs.GetCollection(configs.DB, "serendipias")

func CreateSerendipia() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var serendipia models.Serendipia

		defer cancel()

		//validate the request body
		if err := c.BindJSON(&serendipia); err != nil {
			c.JSON(http.StatusBadRequest, responses.SerendipiaResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}

		//validatea the required fields
		if validationErr := validate.Struct(&serendipia); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.SerendipiaResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newSerendipia := models.Serendipia{
			Id:      primitive.NewObjectID(),
			Title:   serendipia.Title,
			Type:    serendipia.Type,
			Details: serendipia.Details,
			Review:  serendipia.Review,
			Image:   serendipia.Image,
			Link:    serendipia.Link,
		}

		result, err := serendipiaCollection.InsertOne(ctx, newSerendipia)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SerendipiaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.SerendipiaResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})

	}
}
