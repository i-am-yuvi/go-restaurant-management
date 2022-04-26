package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/i-am-yuvi/go-restaurant-management"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.OrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())
}
