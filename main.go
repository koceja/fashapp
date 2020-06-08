package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
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

	router.Run(":" + port)
}
