package routes

import (
	"fakeHebe/models"
	"fakeHebe/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func AddLuckyNumberRoutes(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		if c.Query("pupilId") != "12345" && c.Query("constituentId") == "10" {
			c.JSON(http.StatusOK, utils.ConstructApiResponse(nil, 101))
		} else {
			c.JSON(http.StatusOK, utils.ConstructApiResponse(models.LuckyNumberPayload{
				// as far as I know, the day parameter that the app sends is ignored by the API anyway.
				Day:    time.Now().Format("2006-01-02"),
				Number: randInt(0, 40),
			}, 0))
		}
	})
}
