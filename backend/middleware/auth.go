package middleware

import (
	"net/http"
	"strings"

	"backend/database"
	"backend/models"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates Bearer token and loads current user into context.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		const prefix = "Bearer "

		if !strings.HasPrefix(authHeader, prefix) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing or invalid"})
			c.Abort()
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		var user models.User
		if err := database.DB.Where("token = ?", token).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

// GetCurrentUser retrieves the user stored in Gin context by AuthMiddleware.
func GetCurrentUser(c *gin.Context) (models.User, bool) {
	val, exists := c.Get("currentUser")
	if !exists {
		return models.User{}, false
	}
	user, ok := val.(models.User)
	return user, ok
}
