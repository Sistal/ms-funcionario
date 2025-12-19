package router

import (
	"github.com/Sistal/ms-funcionario/internal/interfaces/http/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(funcionarioHandler *handler.FuncionarioHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/health", funcionarioHandler.HealthCheck)

	v1 := router.Group("/api/v1")
	{
		funcionarios := v1.Group("/funcionarios")
		{
			funcionarios.POST("", funcionarioHandler.CreateFuncionario)
			funcionarios.GET("", funcionarioHandler.GetAllFuncionarios)
			funcionarios.GET("/:id", funcionarioHandler.GetFuncionario)
			funcionarios.PUT("/:id", funcionarioHandler.UpdateFuncionario)
			funcionarios.DELETE("/:id", funcionarioHandler.DeleteFuncionario)
		}
	}

	return router
}
