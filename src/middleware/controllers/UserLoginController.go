package controllers

import (
	"appstud.com/github-core/src/github-api/controllers"
	"github.com/gin-gonic/gin"
)

// Handle for operation "login user"
func handleUserLogin(c *gin.Context) {

	controllers.GitHubLoginUser(c)
}

// Login user
func UserLoginController(engine *gin.Engine) {
	engine.GET("/api/users/login/:username/:password", handleUserLogin)
}
