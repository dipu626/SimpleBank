package routes

import (
	"golang-simple-bank/controllers"

	"github.com/gin-gonic/gin"
)

func AccountRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/accounts", controllers.GetAccounts())
	incomingRoutes.GET("/accounts/:account_id", controllers.GetAccount())
	incomingRoutes.POST("/accounts/signup", controllers.SingUp())
	incomingRoutes.POST("/accounts/login", controllers.Login())
}
