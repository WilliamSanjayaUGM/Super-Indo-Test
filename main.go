package main

import (
	"PROJECT_TEST_GO/controllers"
	"PROJECT_TEST_GO/middlewares"
	"PROJECT_TEST_GO/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	models.ConnectDatabase()
	models.DBMigrate()

	router.GET("/categories", controllers.GetListCategories)
	router.POST("/products", controllers.GetListProducts)
	router.GET("/products/:id", controllers.GetProductById)
	router.POST("/cart", middlewares.AuthMiddleware, controllers.AddToCart)

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.DELETE("/logout", controllers.Logout)

	log.Println("Server started!")
	router.Run("localhost:9090")
}
