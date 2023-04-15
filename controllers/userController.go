package controllers

import (
	"fmt"
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

func SignUpUser(c *gin.Context) {
	// get data from the request body
	var json struct {
		Username string
		Email    string
		Password string
	}
	err := c.Bind(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	fmt.Println("json: ", json)
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
	// print the content of the request body after stringifying it
	fmt.Println(c.Request)
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
	// expects user post request to have email and password in the following format
	// {
	// 	"email": "email",
	// 	"password": "password"
	// }
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
	// check if user is not a dummy user
	if user.DummyFlag == false {
		// compare the password
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"errors": "Invalid password",
			})
			return
		}
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
		"token":   tokenString,
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

//	type User struct {
//		ID           uint      `json:"id" gorm:"primary_key"`
//		Username     string    `json:"username" gorm:"unique;not null"`
//		Email        string    `json:"email" gorm:"unique;not null"`
//		Password     string    `json:"password" gorm:"not null"`
//		ProfileImage string    `json:"profile_image"`
//		CreatedAt    time.Time `json:"created_at"`
//		UpdatedAt    time.Time `json:"updated_at"`
//		Posts        []Post    `json:"posts" gorm:"foreignkey:UserID"`
//		DummyFlag    bool      `json:"dummy_flag"`
//	}
// everthing but password

func GetUser(c *gin.Context) {
	// get the user id from the request params
	id := c.Param("id")
	// get the user
	var user models.UserInfo
	result := initial.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": "User not found",
		})
		return
	}
	// return the user
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	initial.DB.Find(&users)
	// return all attributes except password
	var usersInfo []models.UserInfo
	for _, user := range users {
		usersInfo = append(usersInfo, models.UserInfo{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			ProfileImage: user.ProfileImage,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
			DummyFlag:    user.DummyFlag,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"users": usersInfo,
	})
}
