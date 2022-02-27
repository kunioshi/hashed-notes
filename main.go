package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kunioshi/hashed-notes/api"
	"github.com/kunioshi/hashed-notes/config"
	"github.com/kunioshi/hashed-notes/frontend"
)

func main() {
	router := gin.Default()

	// Delegate the Main Handlers
	frontend.Handler(router.Group("/"))
	api.Handler(router.Group("/api"))

	router.Run(":" + config.GetEnvItem("SV_PORT"))
}
