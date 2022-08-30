package controllers

import (
	"carvescoAPI/configs"
	"carvescoAPI/models"
	"carvescoAPI/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var emailCollection *mongo.Collection = configs.GetCollection(configs.DB, "emails")

var validate = validator.New()

func CreateEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var email models.Email

		defer cancel()

		//validate the request body
		if err := c.BindJSON(&email); err != nil {
			c.JSON(http.StatusBadRequest, responses.EmailResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}

		//validatea the required fields
		if validationErr := validate.Struct(&email); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.EmailResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newEmail := models.Email{
			Id:          primitive.NewObjectID(),
			FromAddress: email.FromAddress,
			Name:        email.Name,
			ToAddress:   email.ToAddress,
			Message:     email.Message,
		}

		result, err := emailCollection.InsertOne(ctx, newEmail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.EmailResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.EmailResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})

	}
}
