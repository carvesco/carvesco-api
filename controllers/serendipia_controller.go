package controllers

import (
	"carvescoAPI/configs"
	"carvescoAPI/models"
	"carvescoAPI/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
func GetASerendipia() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//var serendipia models.Serendipia
		defer cancel()

		sample := bson.D{{"$sample", bson.D{{"size", 1}}}}

		cursor, err := serendipiaCollection.Aggregate(ctx, mongo.Pipeline{sample})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SerendipiaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var result []bson.M
		if err = cursor.All(ctx, &result); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, responses.SerendipiaResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})

	}
}

func EditSerendipia() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		serendipiaId := c.Param("serendipiaId")
		var serendipia models.Serendipia
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(serendipiaId)

		//validate the request body
		if err := c.BindJSON(&serendipia); err != nil {
			c.JSON(http.StatusBadRequest, responses.SerendipiaResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&serendipia); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.SerendipiaResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"details": serendipia.Details, "image": serendipia.Image, "title": serendipia.Title, "type": serendipia.Type, "review": serendipia.Review, "link": serendipia.Link}

		result, err := serendipiaCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SerendipiaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated user details
		var updatedSerendipia models.Serendipia
		if result.MatchedCount == 1 {
			err := serendipiaCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedSerendipia)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.SerendipiaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}
		c.JSON(http.StatusOK, responses.SerendipiaResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedSerendipia}})

	}
}
func GetAllSerendipias() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var serendipias []models.Serendipia
		defer cancel()

		results, err := serendipiaCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SerendipiaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleSrenedipia models.Serendipia
			if err = results.Decode(&singleSrenedipia); err != nil {
				c.JSON(http.StatusInternalServerError, responses.SerendipiaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			serendipias = append(serendipias, singleSrenedipia)
		}

		c.JSON(http.StatusOK,
			responses.SerendipiaResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": serendipias}},
		)
	}
}
