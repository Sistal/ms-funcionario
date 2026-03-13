package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApiKeyMiddleware valida la solicitud mediante el header x-api-key
func ApiKeyMiddleware(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("x-api-key")
		if key == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing x-api-key header"})
			c.Abort()
			return
		}

		if key != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid api key"})
			c.Abort()
			return
		}

		c.Next()
	}
}
