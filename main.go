package main

import (
	"github.com/a/json-app/initial"
	"github.com/gin-gonic/gin"
)

func init() {
	initial.LoadEnv()
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
