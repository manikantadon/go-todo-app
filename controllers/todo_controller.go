package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dummy handler for now
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetTodos works!",
	})
}
