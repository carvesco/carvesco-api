package routes

import (
	"carvescoAPI/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//POST
	router.POST("/user", controllers.CreateUser())

	//get
	router.GET("/user/:userId", controllers.GetAUser())
}
