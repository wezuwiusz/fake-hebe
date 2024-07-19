package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())
	addRoutes(r)
	r.LoadHTMLGlob("public/*")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "index.html", nil)
	})
	r.Run()
}
