// Package middleware contains the application middleware.
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// ZerologMiddleware logs all HTTP requests in structured JSON format.
func ZerologMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("latency", time.Since(start)).
			Str("client_ip", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent()).
			Msg("HTTP Request Logged")
	}
}
