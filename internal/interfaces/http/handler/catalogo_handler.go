package handler

import (
	"net/http"

	"github.com/Sistal/ms-funcionario/internal/application/services"
	"github.com/Sistal/ms-funcionario/internal/interfaces/dto"
	"github.com/gin-gonic/gin"
)

type CatalogoHandler struct {
	service *services.CatalogoService
}

func NewCatalogoHandler(service *services.CatalogoService) *CatalogoHandler {
	return &CatalogoHandler{
		service: service,
	}
}

// ListCargos lista todos los cargos disponibles
// @Summary Listar cargos
// @Description Obtiene la lista de todos los cargos disponibles
// @Tags catálogos
// @Produce json
// @Success 200 {object} dto.SuccessResponse
// @Router /api/v1/cargos [get]
func (h *CatalogoHandler) ListCargos(c *gin.Context) {
	cargos, err := h.service.ListCargos(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	// Convertir a DTOs
	cargoDTOs := make([]dto.CargoDTO, len(cargos))
	for i, cargo := range cargos {
		cargoDTOs[i] = dto.CargoDTO{
			IDCargo:     cargo.IDCargo,
			NombreCargo: cargo.NombreCargo,
		}
	}

	response := gin.H{
		"success": true,
		"data":    cargoDTOs,
		"meta": gin.H{
			"total": len(cargoDTOs),
		},
	}

	c.JSON(http.StatusOK, response)
}

// ListGeneros lista todos los géneros disponibles
// @Summary Listar géneros
// @Description Obtiene la lista de todos los géneros disponibles
// @Tags catálogos
// @Produce json
// @Success 200 {object} dto.SuccessResponse
// @Router /api/v1/generos [get]
func (h *CatalogoHandler) ListGeneros(c *gin.Context) {
	generos, err := h.service.ListGeneros(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	// Convertir a DTOs
	generoDTOs := make([]dto.GeneroDTO, len(generos))
	for i, genero := range generos {
		generoDTOs[i] = dto.GeneroDTO{
			IDGenero:     genero.IDGenero,
			NombreGenero: genero.NombreGenero,
		}
	}

	response := gin.H{
		"success": true,
		"data":    generoDTOs,
		"meta": gin.H{
			"total": len(generoDTOs),
		},
	}

	c.JSON(http.StatusOK, response)
}
