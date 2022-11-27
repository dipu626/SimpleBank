package routes

import (
	"golang-simple-bank/controllers"

	"github.com/gin-gonic/gin"
)

func TransfersRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/transfers", controllers.Transfer())
}
