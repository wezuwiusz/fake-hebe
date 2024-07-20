package routes

import (
	"database/sql"
	"fakeHebe/models"
	"fakeHebe/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddTeacherRoutes(rg *gin.RouterGroup, db *sql.DB) {
	rg.GET("/", func(c *gin.Context) {
		if c.Query("pupilId") != "12345" || len(c.Query("periodId")) < 2 {
			c.JSON(http.StatusOK, utils.ConstructApiResponse(nil, 101))
		} else {
			var teachers []models.Teacher

			rows, err := db.Query("SELECT * FROM teachers")

			if err != nil {
				c.AbortWithStatus(500)
			}

			defer rows.Close()

			for rows.Next() {
				var teacher models.Teacher
				if err := rows.Scan(
					&teacher.Id,
					&teacher.Name,
					&teacher.Surname,
					&teacher.DisplayName,
					&teacher.Position,
					&teacher.Description); err != nil {
					c.AbortWithStatus(500)
				}
				teachers = append(teachers, teacher)
			}

			c.JSON(http.StatusOK, utils.ConstructApiResponse(teachers, 0))
		}
	})
}
