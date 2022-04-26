package routes

import (
	controller "github.com/i-am-yuvi/go-restaurant-management"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUsers())
	incomingRoutes.GET("/users/signup", controller.SignUp())
	incomingRoutes.GET("/users/login", controller.Login())
}
