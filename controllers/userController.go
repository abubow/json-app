package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/a/json-app/initial"
	"github.com/a/json-app/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func SignInUser(c *gin.Context) {
	// get data from the request body
	var json models.User
	c.Bind(&json)
	// validate the input
	var errors []string
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
	// get the user
	var user models.User
	result := initial.DB.Where("email = ?", json.Email).First(&user)
	// the user is not found by email, find by username
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		result = initial.DB.Where("username = ?", json.Email).First(&user)
	}
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": "User not found",
		})
		return
	}
	// compare the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": "Invalid password",
		})
		return
	}
	// create a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret (hmacSampleSecret)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": "Something went wrong",
			"token":  tokenString,
		})
		return
	}

	// make a cookie with the token string, this cookie will be stored in the browser client side, this cookie will be
	// sent back to the server each time the user requests a protected endpoint, it is called a JWT
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", tokenString, 60*60*24, "", "", false, true)
	// return the token
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User logged in",
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   user,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User logged out",
	})
}
