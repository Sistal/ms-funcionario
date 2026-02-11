package middleware

import (
	"net/http"

	"github.com/Sistal/ms-funcionario/internal/interfaces/dto"
	"github.com/gin-gonic/gin"
)

// RequireAdmin verifica que el usuario tenga rol de administrador
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el rol del contexto (establecido por el middleware de autenticación)
		roleInterface, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, dto.NewErrorResponse("Sin permisos suficientes"))
			c.Abort()
			return
		}

		role, ok := roleInterface.(string)
		if !ok {
			c.JSON(http.StatusForbidden, dto.NewErrorResponse("Sin permisos suficientes"))
			c.Abort()
			return
		}

		// Verificar si el rol es administrador
		// Aceptar variaciones: "Administrador", "Admin", "admin", "administrador"
		if role != "Administrador" && role != "Admin" && role != "admin" && role != "administrador" {
			c.JSON(http.StatusForbidden, dto.NewErrorResponse("Sin permisos suficientes"))
			c.Abort()
			return
		}

		c.Next()
	}
}
