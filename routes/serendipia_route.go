package routes

import (
	"carvescoAPI/controllers"

	"github.com/gin-gonic/gin"
)

func SerendipiaRoute(router *gin.Engine) {
	router.POST("/serendipia", controllers.CreateSerendipia())
	router.GET("/serendipia", controllers.GetASerendipia())
	router.PUT("/serendipi/serendipiaId", controllers.EditSerendipia())
	router.GET("/serendipias", controllers.GetAllSerendipias())
}
