package main

import (
	"test/initializers"
	"test/models"
)
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
