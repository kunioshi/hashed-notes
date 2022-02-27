package frontend

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Handles the Front-end/Website access
func Handler(r *gin.RouterGroup) {
	r.GET("/", showHome)
}

func showHome(c *gin.Context) {
	// TODO: Implement Home Page
	fmt.Fprintln(c.Writer, "Welcome to Hashed Notes by Kunioshi!")
}
