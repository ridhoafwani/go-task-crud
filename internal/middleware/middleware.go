package middleware

import (
	"errors"
	"net/http"
	"strings"
	"task-crud/internal/config"
	"task-crud/utils/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := config.AppConfig.JWTSecretKey
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing or invalid token prefix"))
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		userID, username, err := jwt.ValidateToken(token, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}

func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := config.AppConfig.JWTSecretKey
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing or invalid token prefix"))
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		userID, username, err := jwt.ValidateTokenWithoutExpiry(token, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}
