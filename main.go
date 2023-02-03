package main

import (
	"github.com/gin-gonic/gin"
	"github.com/linhnphe163269/GO-JSON-CRUD-API/controllers"
	"github.com/linhnphe163269/GO-JSON-CRUD-API/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
