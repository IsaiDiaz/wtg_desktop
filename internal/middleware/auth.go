package middleware

import (
	"net/http"
	token "wtg_desktop/internal/utils/jwt"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenStr, err := token.ExtractTokenFromHeader(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		claims, err := token.ParseToken(tokenStr, false)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("roles", claims.Roles)
		c.Next()
	}
}

func RequireRoles(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		raw, exists := c.Get("roles")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "roles not found"})
			return
		}

		userRoles := raw.([]string)
		roleMap := make(map[string]bool)
		for _, role := range userRoles {
			roleMap[role] = true
		}

		for _, required := range requiredRoles {
			if roleMap[required] {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
	}
}
