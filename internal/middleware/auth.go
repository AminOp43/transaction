package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	header := c.GetHeader("Authorization")
	parts := strings.SplitN(header, "", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format is incorrect"})
		c.Abort()
		return
	}
	tokenStr := parts[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unsupported method")
		}
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			return nil, errors.New("JWT_SECRET environment variable is not set")
		}
		return []byte(secret), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		c.Abort()
		return
	}
	if token == nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
		c.Abort()
		return
	}
	userId, ok := claims["user_id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
		c.Abort()
		return
	}
	c.Set("user_id", int64(userId))
}
