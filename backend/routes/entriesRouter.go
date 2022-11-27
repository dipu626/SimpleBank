package routes

import (
	"golang-simple-bank/controllers"

	"github.com/gin-gonic/gin"
)

func EntriesRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/entries/deposit", controllers.Deposit())
	incomingRoutes.POST("/entries/withdraw", controllers.Withdraw())
}
