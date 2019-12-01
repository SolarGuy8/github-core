package controllers

import (
	"appstud.com/github-core/src/github-api/controllers"
	"github.com/gin-gonic/gin"
)

// Handle for operation "create user"
func handleCreateUser(c *gin.Context) {

	controllers.GitHubRegisterUser(c)
}

// Handle for operation "get all users"
func handleGetUsers(c *gin.Context) {

	controllers.GitHubGetUsers(c)
}

// create user with specified name and password (Get request)
func CreateUserController(engine *gin.Engine) {
	engine.GET("/api/users/register/:username/:password", handleCreateUser)
}

// Get all user
func GetUsersController(engine *gin.Engine) {
	engine.GET("/api/users", handleGetUsers)
}
