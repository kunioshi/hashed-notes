package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunioshi/hashed-notes/api/services"
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
	u, err := services.CreateUser(c)
	if printError(err, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": u})
}

func getUsers(c *gin.Context) {
	us, err := services.GetUsers(c)
	if printError(err, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": us})
}

func getUser(c *gin.Context) {
	u, err := services.GetUser(c)
	if printError(err, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": u})
}

func updateUser(c *gin.Context) {
	u, err := services.UpdateUser(c)
	if printError(err, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": u})
}

func deleteUser(c *gin.Context) {
	// TODO: Implement Endpoint
	c.JSON(http.StatusServiceUnavailable, "Users/deleteUser not yet implemented!")
}

// Checks whether `err` is nil and if it ISN'T, prints the JSON error
func printError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return true
	}

	return false
}
