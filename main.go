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
	r.GET("/ping", controllers.PingApp)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetAllPosts)
	r.Run()
}
