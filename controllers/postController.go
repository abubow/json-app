package controllers

import (
	"log"

	"github.com/a/json-app/initial"
	"github.com/a/json-app/models"
	"github.com/gin-gonic/gin"
)

//	type Post struct {
//		ID        uint   `json:"id" gorm:"primary_key"`
//		Title     string `json:"title"`
//		Body      string `json:"body"`
//		Author    string `json:"author"`
//		Published bool   `json:"published"`
//	}
func CreatePost(c *gin.Context) {
	// get data from the request body
	var json models.Post
	c.Bind(&json)
	log.Println(json)
	// validate the input
	var errors []string
	if json.Title == "" {
		errors = append(errors, "Title is required\n")
	}
	if json.Body == "" {
		errors = append(errors, "Body is required\n")
	}
	if json.Author == "" {
		errors = append(errors, "Author is required\n")
	}
	if len(errors) > 0 {
		c.JSON(400, gin.H{
			"errors": errors,
		})
		return
	}
	// create a post object
	post := models.Post{
		Title:     json.Title,
		Body:      json.Body,
		Author:    json.Author,
		Published: json.Published,
	}
	// save to the database
	result := initial.DB.Create(&post)
	if result.Error != nil {
		log.Fatal(result.Error)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	// return a response
	log.Println("Post created successfully")
	c.JSON(201, gin.H{
		"message": "Post created successfully",
		"post":    post,
	})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	result := initial.DB.Find(&posts)
	if result.Error != nil {
		log.Fatal(result.Error)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Posts fetched successfully",
		"posts":   posts,
	})
}
