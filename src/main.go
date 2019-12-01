package main

import (
	"appstud.com/github-core/src/middleware/configurations"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	configurations.InitControllers(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
