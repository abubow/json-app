package controllers

import (
	"net/http"

	"github.com/a/json-app/initial"
	"github.com/a/json-app/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// type User struct {
// 	ID           uint   `json:"id" gorm:"primary_key"`
// 	Username     string `json:"username" gorm:"unique;not null"`
// 	Email        string `json:"email" gorm:"unique;not null"`
// 	Password     string `json:"password" gorm:"not null"`
// 	ProfileImage string `json:"profile_image"`
// 	Followers    []User `json:"followers" gorm:"many2many:followers"`
// 	Followings   []User `json:"followings" gorm:"many2many:followings"`
// }

func SignUpUser(c *gin.Context) {
	// get data from the request body
	var json models.User
	c.Bind(&json)
	// validate the input
	var errors []string
	if json.Username == "" {
		errors = append(errors, "Username is required")
	}
	if json.Email == "" {
		errors = append(errors, "Email is required")
	}
	if json.Password == "" {
		errors = append(errors, "Password is required")
	}
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errors,
		})
		return
	}
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	// create a user object
	user := models.User{
		Username: json.Username,
		Email:    json.Email,
		Password: string(hash),
	}
	// save the user in the database
	result := initial.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": result.Error,
		})
		return
	}
	// return the user
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
