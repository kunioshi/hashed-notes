package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Delegate User's CRUD endpoints
func UserHandler(r *gin.RouterGroup) {
	r.POST("", createUser)
	r.GET("", getUsers)
	r.GET("/:id", getUser)
	r.PUT("/:id", updateUser)
	r.DELETE("/:id", deleteUser)
}

func createUser(c *gin.Context) {
	// TODO: Implement Endpoint
	c.JSON(http.StatusServiceUnavailable, "Users/createUser not yet implemented!")
}

func getUser(c *gin.Context) {
	// TODO: Implement Endpoint
	c.JSON(http.StatusServiceUnavailable, "Users/getUser not yet implemented!")
}

func getUsers(c *gin.Context) {
	// TODO: Implement Endpoint
	c.JSON(http.StatusServiceUnavailable, "Users/getUsers not yet implemented!")
}

func updateUser(c *gin.Context) {
	// TODO: Implement Endpoint
	c.JSON(http.StatusServiceUnavailable, "Users/updateUser not yet implemented!")
}

func deleteUser(c *gin.Context) {
	// TODO: Implement Endpoint
	c.JSON(http.StatusServiceUnavailable, "Users/deleteUser not yet implemented!")
}
