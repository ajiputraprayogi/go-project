package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(UserID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte("SECRET_KEY"))
}