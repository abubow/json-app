package main

import (
	"github.com/a/json-app/initial"
	"github.com/a/json-app/models"
)

func init() {
	initial.LoadEnv()
	initial.ConnectToDB()
}

func main() {
	// migrate
	initial.DB.AutoMigrate(&models.Post{})
}
