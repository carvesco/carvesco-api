package routes

import (
	"carvescoAPI/controllers"

	"github.com/gin-gonic/gin"
)

func SerendipiaRoute(router *gin.Engine) {
	router.POST("/serendipia", controllers.CreateSerendipia())
}
