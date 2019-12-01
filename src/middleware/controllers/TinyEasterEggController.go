package controllers

import (
	"appstud.com/github-core/src/github-api/models/responses"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"time"
)

// Array of easter eggs
type TinyEasterEggs struct {
	Array []responses.TinyEasterEgg
}

// Handle for operation "easter egg"
func handleTinyEasterEgg(c *gin.Context) {

	log.Printf("USER FOUNDS TINY EGG!")

	c.JSON(200, TinyEasterEggs{Array: []responses.TinyEasterEgg{
		{
			Name:    "My mom is in love with me",
			Version: "1.0",
			Time:    time.Unix(0, -446723100*int64(time.Millisecond)).Format("Mon, 02 Jan 2006 15:04:05 MST"),
		},
		{
			Name:    "I go to the future and my mom end up with the wrong guy",
			Version: "2.0",
			Time:    time.Unix(0, -1445470140*int64(time.Millisecond)).Format("Mon, 02 Jan 2006 15:04:05 MST"),
		},
		{
			Name:    "I go to the past and you will not believe what happens next",
			Version: "3.0",
			Time:    time.Unix(0, math.MaxInt16*int64(time.Millisecond)).Format("Mon, 02 Jan 2006 15:04:05 MST"),
		},
	}})

}

// Easter egg
func TinyEasterEggController(engine *gin.Engine) {
	engine.GET("/api/timemachine/logs/mcfly", handleTinyEasterEgg)
}
