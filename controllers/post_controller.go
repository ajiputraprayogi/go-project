package controllers

import (
	"go-project/config"
	"go-project/models"
	"go-project/dto"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&post)
	c.JSON(200, post)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post

	// Preload user
	config.DB.Preload("User").Find(&posts)

	// Mapping ke DTO
	var result []dto.PostResponse
	for _, p := range posts {
		result = append(result, dto.PostResponse{
			ID:      p.ID,
			Title:   p.Title,
			Content: p.Content,
			UserID:  p.UserID,
			User: dto.UserResponse{
				ID:    p.User.ID,
				Name:  p.User.Name,
				Email: p.User.Email,
			},
		})
	}

	c.JSON(200, result)
}

func GetPost(c *gin.Context) {
	id := c.Param(("id"))
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	c.ShouldBindJSON(&post)
	config.DB.Save(&post)
	c.JSON(200, post)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
	}

	c.ShouldBindJSON(&post)
	config.DB.Save(&post)
	c.JSON(200, post)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Post{}, id)
	c.JSON(200, gin.H{"message": "deleted"})
}