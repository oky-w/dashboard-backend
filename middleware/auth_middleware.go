package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/okyws/dashboard-backend/config"
	"github.com/okyws/dashboard-backend/utils"
)

// AuthMiddleware untuk validasi JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Authorization header not found")
			c.Abort()

			return
		}

		// Remove "Bearer " prefix
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := config.ParseToken(tokenString)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
			c.Abort()

			return
		}

		// Set the claims in the context
		c.Set("id", claims.ID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
