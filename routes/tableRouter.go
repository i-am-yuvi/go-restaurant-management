package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/i-am-yuvi/go-restaurant-management"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:table_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controller.UpdateTable())
}
