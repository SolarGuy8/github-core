package controllers

import (
	"appstud.com/github-core/src/github-api/models/responses"
	"appstud.com/github-core/src/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func handleCheckConnection(c *gin.Context, response *http.Response, err error) {

	if err != nil {
		log.Fatal(err)
		utils.CreateError(c, 500, "Server error (Send request)")
		return
	}

	// read body
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
		utils.CreateError(c, 500, "Server error (Read body)")
		return
	}

	if response.StatusCode != 200 {
		log.Fatal("Unexpected status code", response.StatusCode)
		utils.CreateError(c, 500, "Server error (Unexpected status code")
		return
	}

	log.Printf("Body: %s\n", body)

	//c.JSON(200, json.Unmarshal([]byte(body), &checkConnectionResponse))

	if response.StatusCode == 200 {
		c.JSON(200, responses.CheckConnectionResponse{
			Name:    "github-api",
			Version: "1.0",
			Time:    time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"),
		})
	}
}

// CheckConnectionController - Route controller
func GitHubApiCheckConnection(c *gin.Context) {

	response, err := http.Get("https://api.github.com/users/moficodes")

	handleCheckConnection(c, response, err)
}
