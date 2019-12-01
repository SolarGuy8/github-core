package controllers

import (
	"appstud.com/github-core/src/github-api/models/responses"
	"appstud.com/github-core/src/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// handle for operation "register user"
func handleRegisterUser(c *gin.Context, username string, password string) {

	log.Printf("Register user api: Username: %s\n, password: %s\n", username, password)

	c.JSON(200, responses.RegisterUser{
		Username: username,
		Token:    utils.TokenGenerator(),
	})
}

// handle for operation "get users"
func handleGetUsers(c *gin.Context) {

	log.Printf("Get users api")

	c.JSON(200, responses.Users{Array: []responses.User{
		{Username: "Daymond"},
		{Username: "mkdir"},
		{Username: "James78"},
		{Username: "Kasnop1"},
		{Username: "qwerty"},
	}})
}

// handle for operation "login user"
func handleLoginUser(c *gin.Context, username string, password string) {

	log.Printf("Login user api: Username: %s\n, password: %s\n", username, password)

	c.JSON(200, responses.RegisterUser{
		Username: username,
		Token:    utils.TokenGenerator(),
	})
}

// register user
func GitHubRegisterUser(c *gin.Context) {

	handleRegisterUser(c, c.Param("username"), c.Param("password"))
}

// get users
func GitHubGetUsers(c *gin.Context) {

	handleGetUsers(c)
}

// login user
func GitHubLoginUser(c *gin.Context) {

	handleLoginUser(c, c.Param("username"), c.Param("password"))
}
