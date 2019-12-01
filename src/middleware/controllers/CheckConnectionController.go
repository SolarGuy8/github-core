package controllers

import (
	"appstud.com/github-core/src/github-api/controllers"
	"github.com/gin-gonic/gin"
)

// Handle for operation "check connection"
func handleCheckConnection(c *gin.Context) {

	controllers.GitHubApiCheckConnection(c)
}

// Check connection to the api
func CheckConnectionController(engine *gin.Engine) {
	engine.GET("/api/healthcheck", handleCheckConnection)
}
