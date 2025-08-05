package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manikantadon/todo-backend/routes"
)

func main() {
	// Create a Gin router
	router := gin.Default()

	// Register app routes
	routes.RegisterTodoRoutes(router)

	// Test route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Start the server
	router.Run(":8080") // default port 8080
}
