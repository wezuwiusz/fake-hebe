package main

import (
	"database/sql"
	"fakeHebe/routes"
	"fakeHebe/utils"
	"github.com/gin-gonic/gin"
	"os"
)

func addRoutes(router *gin.Engine, db *sql.DB) {
	wezuwiusz := router.Group("/powiatwezuwiusz")
	if os.Getenv("CHECK_HEADERS") == "true" {
		wezuwiusz.Use(utils.CheckHeaders())
	}
	routes.AddAppStartRoutes(wezuwiusz.Group("/12345/api/mobile/appstart"))
	routes.AddLuckyNumberRoutes(wezuwiusz.Group("/12345/api/mobile/school/lucky"))
	routes.AddTeacherRoutes(wezuwiusz.Group("/12345/api/mobile/teacher/byPeriod"), db)
}
