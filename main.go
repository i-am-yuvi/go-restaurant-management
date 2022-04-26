package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/i-am-yuvi/go-restaurant-management/database"
	"github.com/i-am-yuvi/go-restaurant-management/middleware"
	"github.com/i-am-yuvi/go-restaurant-management/routes"
	"go.uber.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
