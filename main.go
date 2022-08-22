package main

import (
	"carvescoAPI/configs"
	"os"

	"carvescoAPI/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	//run database
	configs.ConnectDB()

	//routes
	routes.EmailRoute(router)

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
