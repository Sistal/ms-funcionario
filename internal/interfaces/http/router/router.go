package router

import (
	"github.com/Sistal/ms-funcionario/internal/interfaces/http/handler"
	"github.com/gin-gonic/gin"
	
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(funcionarioHandler *handler.FuncionarioHandler) *gin.Engine {
	router := gin.Default()

	// Swagger UI endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint
	router.GET("/health", funcionarioHandler.HealthCheck)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Rutas de funcionarios
		funcionarios := v1.Group("/funcionarios")
		{
			// CRUD básico
			funcionarios.POST("", funcionarioHandler.CreateFuncionario)
			funcionarios.GET("", funcionarioHandler.GetAllFuncionarios)
			funcionarios.GET("/:id", funcionarioHandler.GetFuncionario)
			funcionarios.PUT("/:id", funcionarioHandler.UpdateFuncionario)
			funcionarios.DELETE("/:id", funcionarioHandler.DeleteFuncionario)

			// Búsquedas con filtros
			funcionarios.GET("/filter", funcionarioHandler.GetFuncionariosByFilter)
			funcionarios.GET("/rut/:rut", funcionarioHandler.GetFuncionarioByRut)
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
		}
	}

	return router
}
