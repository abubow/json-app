package main

import (
	"time"

	"github.com/a/json-app/controllers"
	"github.com/a/json-app/initial"
	"github.com/a/json-app/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initial.LoadEnv()
	initial.ConnectToDB()
}
func main() {
	r := gin.Default()
	// setup cors to allow requests from all origins
	r.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
	)
	// post routes
	r.GET("/ping", controllers.PingApp)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PATCH("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	// user routes
	r.POST("/users", controllers.SignUpUser)
	r.POST("/login", controllers.SignInUser)
	r.POST("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/users/:id", controllers.GetUser)
	r.GET("/users", controllers.GetAllUsers)
	r.Run()
}
