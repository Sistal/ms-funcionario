package middleware

import (
	"net/http"
	"strings"

	"github.com/Sistal/ms-funcionario/internal/infrastructure/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware valida el token JWT usando el servicio de autenticación
func AuthMiddleware(authClient *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extraer token del header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		// Validar formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		token := parts[1]

		// Validar token con el servicio de autenticación
		response, err := authClient.ValidateToken(token)
		if err != nil || !response.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// Guardar información del usuario en el contexto
		c.Set("user_id", response.UserID)
		c.Set("username", response.Username)
		c.Set("role", response.Role)

		c.Next()
	}
}
