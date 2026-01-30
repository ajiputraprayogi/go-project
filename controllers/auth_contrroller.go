package controllers

import (
	"go-project/config"
	"go-project/dto"
	"go-project/models"
	"go-project/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var dto dto.RegisterDTO
	c.ShouldBindJSON(&dto)

	hash, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)

	user := models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: string(hash),
	}

	config.DB.Create(&user)
	c.JSON(201, gin.H{"message": "registered"})
}

func Login(c *gin.Context) {
	var dto dto.LoginDTO
	c.ShouldBindJSON(&dto)

	var user models.User
	config.DB.Where("email = ?", dto.Email).First(&user)

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(dto.Password),
	); err != nil {
		c.JSON(401, gin.H{"error": "invalid"})
		return
	}

	token, _ := utils.GenerateToken(user.ID)

	c.SetCookie(
		"access_token",
		token,
		3600*24,
		"/",
		"",
		false,
		true,
	)

	c.JSON(200, gin.H{"message": "login success"})
}

func Logout(c *gin.Context) {
	c.SetCookie(
		"access_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	c.JSON(200, gin.H{"message": "Logout Berhasil"})
}