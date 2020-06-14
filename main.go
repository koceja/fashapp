package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/lib/pq"

)

func insertImages(db *sql.DB, personId string) {
	ins := "INSERT INTO test (personId, images) VALUES ($1, $2)"

	// "tags" is the list of tags, as a string slice
	images := []string{}

	// the pq.Array function is the secret sauce
	if _, err := db.Exec(ins, personId, pq.Array(images)); err != nil {
		log.Fatal(err)
	}
}

func getImages(db *sql.DB, personId string) (images []string) {
    // the select query, returning 1 column of array type
    sel := "SELECT images FROM test WHERE personId=$1"

    // wrap the output parameter in pq.Array for receiving into it
    if err := db.QueryRow(sel, personId).Scan(pq.Array(&images)); err != nil {
        log.Fatal(err)
    }

    return
}

func updateImages(db *sql.DB, personId string, imageUrl string) {
	upd := "UPDATE test SET images = array_append(images, $1) WHERE personId = $2"

	if _, err := db.Exec(upd, imageUrl, personId); err != nil {
		log.Fatal(err)
	}
}


func getImageHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS test (personId text, images text[])"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q", err))
			return
		}
		personId := c.PostForm("personId")
		imageUrl := c.PostForm("image")

		insertImages(db, personId)
		updateImages(db, personId, imageUrl)
	}
}

func profileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS test (personId text, images text[])"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q", err))
			return
		}
		id := c.Param("personId")

		images := getImages(db, id)

		var temp string

		if len(images) > 0 {
			temp = images[0]
		} else {
			temp = "empty"
		}

		c.HTML(http.StatusOK, "profile.tmpl.html", gin.H{"personId": id, "images": temp})
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

	router.GET("/profile/:personId", profileHandler(db))

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
