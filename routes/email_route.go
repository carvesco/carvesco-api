package routes

import (
	"carvescoAPI/controllers"

	"github.com/gin-gonic/gin"
)

func EmailRoute(router *gin.Engine) {
	router.POST("/email", controllers.CreateEmail())
}
