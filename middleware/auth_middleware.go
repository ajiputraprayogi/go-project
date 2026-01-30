package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(401, gin.H{
				"message": "Unauthorized: token not found",
			})
			c.Abort()
			return
		}

		token, _ := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte("SECRET_KEY"), nil
		})

		if !token.Valid {
			c.JSON(401, gin.H{
				"message": "Unauthorized: invalid token",
			})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("user_id", claims["user_id"])

		c.Next()
	}
}
