package configurations

import (
	"appstud.com/github-core/src/middleware/controllers"
	"github.com/gin-gonic/gin"
)

// Init all controllers
func InitControllers(r *gin.Engine) {
	controllers.CheckConnectionController(r) // Check connection to the api
	controllers.TinyEasterEggController(r)   // easter egg
	controllers.CreateUserController(r)      // create user
	controllers.GetUsersController(r)        // get all users
	controllers.UserLoginController(r)       // login user
}
