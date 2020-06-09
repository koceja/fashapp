package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)


func getImageHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS images (personId img)"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q", err))
			return
		}
	}
}



func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"));


    if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/men/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "men.tmpl.html", nil)
	})

	router.GET("/women/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "women.tmpl.html", nil)
	})

	router.GET("/both/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "both.tmpl.html", nil)
	})

	router.POST("/men/", getImageHandler(db))

	router.POST("/women/", getImageHandler(db))

	router.POST("/both/", getImageHandler(db))

	router.GET("/db/", getImageHandler(db))


	router.Run(":" + port)
}
