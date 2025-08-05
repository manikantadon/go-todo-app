package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manikantadon/todo-backend/controllers"
)

func RegisterTodoRoutes(router *gin.Engine) {
	todoRoutes := router.Group("/api/todos")
	{
		todoRoutes.GET("/", controllers.GetTodos)
		// Other routes like POST, PUT, DELETE will go here later
	}
}
