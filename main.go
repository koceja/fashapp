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
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS images (personId string)"); err != nil {
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
		c.Redirect(302, "/login")
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl.html", nil)
	})


	router.GET("/home/:personId", func(c *gin.Context) {
		id := c.Param("personId")
		c.HTML(http.StatusOK, "home.tmpl.html", gin.H{"personId": id})
	})

	router.GET("/men/:personId", func(c *gin.Context) {
		id := c.Param("personId")
		c.HTML(http.StatusOK, "men.tmpl.html", gin.H{"personId": id})
	})

	router.GET("/women/:personId", func(c *gin.Context) {
		id := c.Param("personId")
		c.HTML(http.StatusOK, "women.tmpl.html", gin.H{"personId": id})
	})

	router.GET("/both/:personId", func(c *gin.Context) {
		id := c.Param("personId")
		c.HTML(http.StatusOK, "both.tmpl.html", gin.H{"personId": id})
	})

	router.GET("/profile/:personId", func(c *gin.Context) {
		id := c.Param("personId")
		c.HTML(http.StatusOK, "profile.tmpl.html", gin.H{"personId": id})
	})

	router.POST("/men/:personId", getImageHandler(db))

	router.POST("/women/:personId", getImageHandler(db))

	router.POST("/both/:personId", getImageHandler(db))

	router.POST("/login", func(c *gin.Context) {
		id := c.PostForm("personId")
		if (id == "") {
			c.Redirect(302, "/login")
		} else {
			c.Redirect(302, "/home/" + id)
		}
	})


	router.GET("/db/", getImageHandler(db))


	router.Run(":" + port)
}
