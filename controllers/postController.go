package controllers

import (
	"log"
	"time"

	"github.com/a/json-app/initial"
	"github.com/a/json-app/models"
	"github.com/gin-gonic/gin"
)

// type Post struct {
// 	ID        uint   `json:"id" gorm:"primary_key"`
// 	Title     string `json:"title"`
// 	Body      string `json:"body"`
// 	Published bool   `json:"published"`
// 	CreatedAt string `json:"created_at"`
// 	UpdatedAt string `json:"updated_at"`
// 	Author    User   `json:"author" gorm:"foreignkey:AuthorID"`
// 	AuthorID  uint   `json:"author_id"`
// }

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
	if json.AuthorID == 0 {
		errors = append(errors, "Author is required\n")
	}
	if len(errors) > 0 {
		c.JSON(400, gin.H{
			"errors": errors,
		})
		return
	}
	// get the author
	var author models.User
	resultU := initial.DB.First(&author, json.AuthorID)
	if resultU.Error != nil {
		c.JSON(404, gin.H{
			"message": "Author not found",
		})
		return
	}
	// set the author
	json.Author = author
	// create a post object
	post := models.Post{
		Title:     json.Title,
		Body:      json.Body,
		Author:    json.Author,
		Published: json.Published,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
		AuthorID:  json.AuthorID,
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

func UpdatePost(c *gin.Context) {
	// get the post id from the url params
	id := c.Param("id")
	// get the post with that id
	var post models.Post
	result := initial.DB.First(&post, id)
	if result.Error != nil {
		log.Fatal(result.Error)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
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
	if json.AuthorID == 0 {
		errors = append(errors, "Author is required\n")
	}
	if len(errors) > 0 {
		c.JSON(400, gin.H{
			"errors": errors,
		})
		return
	}
	// get the author
	var author models.User
	resultU := initial.DB.First(&author, json.AuthorID)
	if resultU.Error != nil {
		c.JSON(404, gin.H{
			"message": "Author not found",
		})
		return
	}
	// set the author
	json.Author = author
	// update the post object
	post.Title = json.Title
	post.Body = json.Body
	post.Author = json.Author
	post.Published = json.Published
	post.AuthorID = json.AuthorID
	post.UpdatedAt = time.Now().Format(time.RFC3339)
	// save to the database
	result = initial.DB.Save(&post)
	if result.Error != nil {
		log.Fatal(result.Error)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	// return a response
	log.Println("Post updated successfully")
	c.JSON(201, gin.H{
		"message": "Post updated successfully",
		"post":    post,
	})
}

func DeletePost(c *gin.Context) {
	// get the post id from the url params
	id := c.Param("id")
	// get the post with that id
	var post models.Post
	result := initial.DB.First(&post, id)
	if result.Error != nil {
		log.Fatal(result.Error)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	// delete the post
	result = initial.DB.Delete(&post)
	if result.Error != nil {
		log.Fatal(result.Error)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	// return a response
	log.Println("Post deleted successfully")
	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
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

func GetPost(c *gin.Context) {
	// get the post id from the url params
	id := c.Param("id")
	// get the post with that id
	var post models.Post
	result := initial.DB.First(&post, id)
	if result.Error != nil {
		log.Fatal(result.Error)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Post fetched successfully",
		"post":    post,
	})
}
