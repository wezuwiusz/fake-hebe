package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Printf("Successfully loaded .env\n")
	}

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DBHOST"),
		DBName: os.Getenv("DBNAME"),
	}
	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully established connection to MySQL database server.\n")
	}

	r := gin.Default()
	r.Use(gin.Recovery())
	addRoutes(r, db)
	r.LoadHTMLGlob("public/*")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "index.html", nil)
	})
	r.Run()
}
