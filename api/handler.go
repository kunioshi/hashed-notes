package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kunioshi/hashed-notes/api/controllers"
)

// Delegate the API handlers to the corresponding endpoint to its features' Controller
func Handler(r *gin.RouterGroup) {
	controllers.UserHandler(r.Group("/users"))
}
