package helper

import (
	"go-project/config"
	"go-project/dto_global"
	"go-project/models"
	// "net/http"
	"github.com/gin-gonic/gin"
)


func GetUserOptions(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)

	var result []dto_global.UserOptions
	for _, u := range users {
		result = append(result, dto_global.UserOptions{
			ID:   u.ID,
			Name: u.Name,
		})
	}
	c.JSON(200, result)
}
