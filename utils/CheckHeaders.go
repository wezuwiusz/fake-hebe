package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement certificate checking
		if c.GetHeader("vapi") == "1" &&
			len(c.GetHeader("vcanonicalurl")) > 1 &&
			len(c.GetHeader("vcontext")) > 1 &&
			len(c.GetHeader("vdate")) > 1 &&
			(c.GetHeader("vos") == "Android" || c.GetHeader("vos") == "iOS") &&
			len(c.GetHeader("vversioncode")) >= 3 {
			c.Next()
		} else {
			c.JSON(http.StatusOK, ConstructApiResponse(nil, 102))
			c.Abort()
			return
		}
	}
}
