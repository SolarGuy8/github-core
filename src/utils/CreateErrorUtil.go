package utils

import (
	error2 "appstud.com/github-core/src/github-api/models/responses/error"
	"github.com/gin-gonic/gin"
	"time"
)

// Create error response with specified code and message. Also set current time
func CreateError(c *gin.Context, code int, message string) {

	c.JSON(code, error2.ErrorResponse{
		Message: message,
		Time:    time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"),
	})
}
