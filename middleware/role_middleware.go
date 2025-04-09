package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okyws/dashboard-backend/utils"
)

// CheckRoleMiddleware is a middleware that checks if the user has the required role.
func CheckRoleMiddleware(requiredRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			utils.ErrorResponse(c, http.StatusUnauthorized, "User role not found")
			c.Abort()

			return
		}

		if !contains(requiredRole, userRole.(string)) {
			utils.ErrorResponse(c, http.StatusUnauthorized, "User role not valid")
			c.Abort()

			return
		}

		c.Next()
	}
}

func contains(roles []string, currentRole string) bool {
	for _, role := range roles {
		if role == currentRole {
			return true
		}
	}

	return false
}
