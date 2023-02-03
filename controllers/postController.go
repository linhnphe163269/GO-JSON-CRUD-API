package controllers

import (
	"GO-JSON-CRUD-API/initializers"
	"GO-JSON-CRUD-API/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	c.BindJSON(&body)
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(500)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
func PostsShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}
func PostsUpdate(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	c.BindJSON(&body)
	post.Title = body.Title
	post.Body = body.Body
	initializers.DB.Save(&post)
	c.JSON(200, gin.H{
		"post": post,
	})
}
func PostsDelete(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	initializers.DB.Delete(&post)
	c.Status(200)
}
