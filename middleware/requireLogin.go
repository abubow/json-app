package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/a/json-app/initial"
	"github.com/a/json-app/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	// get token from cookies
	tokenString, err := c.Cookie("token")
	if err != nil {
		fmt.Println("error getting token from cookies")
		// expand and print c
		fmt.Printf("%+v \n", c)
		// get it from body
		var json struct {
			Token string
		}
		c.Bind(&json)
		fmt.Println("json: ", json)
		tokenString = json.Token
	}
	fmt.Println("tokenString: ", tokenString)
	// print full request body
	// fmt.Println(c.Request.Body)
	// Decode/Validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		fmt.Println("error decoding token")
		fmt.Println("error decoding token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Check if token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check exp
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Get the user from the database based on the information in the token
		var user models.User
		initial.DB.First(&user, claims["sub"])
		// if the user does not exist, return unauthorized
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			// otherwise attach the user to the context and call the next handler
			c.Set("user", user)
			// expand and print user
			fmt.Printf("%+v \n", user)
			c.Next()
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()

}
