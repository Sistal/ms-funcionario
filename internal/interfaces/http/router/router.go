package router

import (
	"github.com/Sistal/ms-funcionario/internal/infrastructure/middleware"
	"github.com/Sistal/ms-funcionario/internal/interfaces/http/handler"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(funcionarioHandler *handler.FuncionarioHandler, profileHandler *handler.ProfileHandler, catalogoHandler *handler.CatalogoHandler, allowedOrigins string) *gin.Engine {
	router := gin.Default()

	// CORS Middleware
	router.Use(middleware.CORSMiddleware(allowedOrigins))

	// Swagger UI endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint
	router.GET("/health", funcionarioHandler.HealthCheck)

	// API v1 routes (sin autenticación requerida)
	v1 := router.Group("/api/v1")
	{
		// Contract Routes (/employees) - BFF endpoints
		employees := v1.Group("/employees")
		{
			employees.GET("/:userId/profile", profileHandler.GetProfile)
			employees.PUT("/:userId/contact", profileHandler.UpdateContact)
			employees.GET("/:userId/measurements", profileHandler.GetMeasurements)
			employees.PUT("/:userId/measurements", profileHandler.UpdateMeasurements)
		}

		// Statistics endpoint
		v1.GET("/funcionarios/me/stats", profileHandler.GetMyStats)

		// Catálogos (públicos para usuarios autenticados)
		v1.GET("/cargos", catalogoHandler.ListCargos)
		v1.GET("/generos", catalogoHandler.ListGeneros)

		// Operaciones Globales (BFF legacy)
		v1.GET("/branches", funcionarioHandler.ListBranches)
		v1.POST("/transfers", funcionarioHandler.RequestTransfer)

		// Rutas de funcionarios
		funcionarios := v1.Group("/funcionarios")
		{
			// CRUD básico
			funcionarios.POST("", funcionarioHandler.CreateFuncionario)
			funcionarios.GET("", funcionarioHandler.GetAllFuncionarios)
			funcionarios.GET("/:id", funcionarioHandler.GetFuncionario)
			funcionarios.PUT("/:id", funcionarioHandler.UpdateFuncionario)
			funcionarios.DELETE("/:id", funcionarioHandler.DeleteFuncionario)

			// Búsquedas específicas
			funcionarios.GET("/by-usuario/:userId", funcionarioHandler.GetFuncionarioByUserID)

			// Búsquedas - IMPORTANTE: Estas deben ir ANTES de /:id para evitar conflictos
			funcionarios.GET("/filter", funcionarioHandler.GetFuncionariosByFilter)
			funcionarios.GET("/buscar/rut/:rut", funcionarioHandler.GetFuncionarioByRut) // Ruta correcta según contrato
			funcionarios.GET("/empresa/:id_empresa", funcionarioHandler.GetFuncionariosByEmpresa)
			funcionarios.GET("/sucursal/:id_sucursal", funcionarioHandler.GetFuncionariosBySucursal)
			funcionarios.GET("/segmento/:id_segmento", funcionarioHandler.GetFuncionariosBySegmento)

			// Activación/Desactivación
			funcionarios.PATCH("/:id/activate", funcionarioHandler.ActivateFuncionario)
			funcionarios.PATCH("/:id/deactivate", funcionarioHandler.DeactivateFuncionario)

			// Rutas de medidas
			funcionarios.POST("/:id/medidas", funcionarioHandler.CreateMedidas)
			funcionarios.GET("/:id/medidas", funcionarioHandler.GetMedidasActivas)
			funcionarios.PUT("/:id/medidas", funcionarioHandler.UpdateMedidas)
			funcionarios.GET("/:id/medidas/historial", funcionarioHandler.GetHistorialMedidas)

			// Registrar funcionario (flujo alternativo)
			funcionarios.POST("/register", funcionarioHandler.RegisterFuncionario)
		}
	}

	return router
}
