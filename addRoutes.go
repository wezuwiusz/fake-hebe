package main

import (
	"fakeHebe/routes"
	"fakeHebe/utils"
	"github.com/gin-gonic/gin"
)

func addRoutes(router *gin.Engine) {
	wezuwiusz := router.Group("/powiatwezuwiusz")
	wezuwiusz.Use(utils.CheckHeaders())
	routes.AddAppStartRoutes(wezuwiusz.Group("/12345/api/mobile/appstart"))
	routes.AddLuckyNumberRoutes(wezuwiusz.Group("/12345/api/mobile/school/lucky"))
}
