package routes

import (
	"fakeHebe/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAppStartRoutes(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		if c.Query("pupilId") != "12345" {
			c.JSON(http.StatusOK, utils.ConstructApiResponse(nil, 101))
		} else {
			c.JSON(http.StatusOK, utils.ConstructApiResponse(true, 0))
		}
	})
}
