package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunioshi/hashed-notes/api/services"
)

func EmailConfirmHandler(c *gin.RouterGroup) {
	c.GET("/:email/:confirmation", confirmEmail)
}

func confirmEmail(c *gin.Context) {
	p := c.Params
	r, err := services.ConfirmEmail(p.ByName("email"), p.ByName("confirmation"))
	if printError(err, c) {
		return
	}

	if !r {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't confirm the email"})
	}

	c.JSON(http.StatusOK, nil)
}
