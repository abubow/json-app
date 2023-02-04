package main

import (
	"github.com/a/json-app/controllers"
	"github.com/a/json-app/initial"
	"github.com/gin-gonic/gin"
)

func init() {
	initial.LoadEnv()
	initial.ConnectToDB()
}
func main() {
	r := gin.Default()
	// post routes
	r.GET("/ping", controllers.PingApp)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PATCH("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	// user routes
	r.POST("/users", controllers.SignUpUser)
	r.Run()
}
