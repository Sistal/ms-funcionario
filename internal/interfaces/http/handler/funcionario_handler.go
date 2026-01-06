package handler

import (
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/Sistal/ms-funcionario/internal/application/services"
	"github.com/Sistal/ms-funcionario/internal/interfaces/dto"
	"github.com/gin-gonic/gin"
)

type FuncionarioHandler struct {
	service *services.FuncionarioService
}

func NewFuncionarioHandler(service *services.FuncionarioService) *FuncionarioHandler {
	return &FuncionarioHandler{
		service: service,
	}
}

// CreateFuncionario crea un nuevo funcionario
// @Summary Crear funcionario
// @Description Crea un nuevo funcionario en el sistema
// @Tags funcionarios
// @Accept json
// @Produce json
// @Param funcionario body dto.CreateFuncionarioRequest true "Datos del funcionario"
// @Success 201 {object} dto.FuncionarioResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /api/v1/funcionarios [post]
func (h *FuncionarioHandler) CreateFuncionario(c *gin.Context) {
	var req dto.CreateFuncionarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	funcionario := req.ToFuncionario()
	if err := h.service.CreateFuncionario(c.Request.Context(), funcionario); err != nil {
		if errors.Is(err, services.ErrFuncionarioAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, services.ErrInvalidInput) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, dto.ToFuncionarioResponse(funcionario))
}

// GetFuncionario obtiene un funcionario por ID
// @Summary Obtener funcionario
// @Description Obtiene un funcionario por su ID
// @Tags funcionarios
// @Produce json
// @Param id path int true "ID del funcionario"
// @Success 200 {object} dto.FuncionarioResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/funcionarios/{id} [get]
func (h *FuncionarioHandler) GetFuncionario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	funcionario, err := h.service.GetFuncionario(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "funcionario not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.ToFuncionarioResponse(funcionario))
}

// GetAllFuncionarios obtiene todos los funcionarios
// @Summary Listar funcionarios
// @Description Obtiene la lista de todos los funcionarios
// @Tags funcionarios
// @Produce json
// @Success 200 {array} dto.FuncionarioResponse
// @Router /api/v1/funcionarios [get]
func (h *FuncionarioHandler) GetAllFuncionarios(c *gin.Context) {
	funcionarios, err := h.service.GetAllFuncionarios(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.ToFuncionarioResponseList(funcionarios))
}

// GetFuncionariosByFilter obtiene funcionarios con filtros
// @Summary Filtrar funcionarios
// @Description Obtiene funcionarios según criterios de filtrado con paginación
// @Tags funcionarios
// @Produce json
// @Param rut_funcionario query string false "RUT del funcionario"
// @Param email query string false "Email del funcionario"
// @Param id_empresa_cliente query int false "ID de la empresa cliente"
// @Param id_sucursal query int false "ID de la sucursal"
// @Param id_segmento query int false "ID del segmento"
// @Param id_estado query int false "ID del estado"
// @Param id_cargo query int false "ID del cargo"
// @Param tallas_registradas query bool false "Tallas registradas"
// @Param limit query int false "Límite de resultados" default(20)
// @Param offset query int false "Offset para paginación" default(0)
// @Success 200 {object} dto.PaginatedFuncionariosResponse
// @Router /api/v1/funcionarios/filter [get]
func (h *FuncionarioHandler) GetFuncionariosByFilter(c *gin.Context) {
	var filterReq dto.FuncionarioFilterRequest
	if err := c.ShouldBindQuery(&filterReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Establecer valores por defecto
	if filterReq.Limit <= 0 {
		filterReq.Limit = 20
	}
	if filterReq.Offset < 0 {
		filterReq.Offset = 0
	}

	filter := filterReq.ToFilter()
	
	funcionarios, err := h.service.GetFuncionariosByFilter(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	total, err := h.service.CountFuncionarios(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(filterReq.Limit)))

	response := dto.PaginatedFuncionariosResponse{
		Data:       dto.ToFuncionarioResponseList(funcionarios),
		Total:      total,
		Limit:      filterReq.Limit,
		Offset:     filterReq.Offset,
		TotalPages: totalPages,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateFuncionario actualiza un funcionario
// @Summary Actualizar funcionario
// @Description Actualiza los datos de un funcionario existente
// @Tags funcionarios
// @Accept json
// @Produce json
// @Param id path int true "ID del funcionario"
// @Param funcionario body dto.UpdateFuncionarioRequest true "Datos actualizados"
// @Success 200 {object} dto.FuncionarioResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/funcionarios/{id} [put]
func (h *FuncionarioHandler) UpdateFuncionario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	var req dto.UpdateFuncionarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	funcionario := req.ToFuncionario(id)
	if err := h.service.UpdateFuncionario(c.Request.Context(), funcionario); err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "funcionario not found"})
			return
		}
		if errors.Is(err, services.ErrInvalidInput) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.ToFuncionarioResponse(funcionario))
}

// DeleteFuncionario elimina un funcionario
// @Summary Eliminar funcionario
// @Description Elimina un funcionario del sistema
// @Tags funcionarios
// @Param id path int true "ID del funcionario"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/funcionarios/{id} [delete]
func (h *FuncionarioHandler) DeleteFuncionario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	if err := h.service.DeleteFuncionario(c.Request.Context(), id); err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "funcionario not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetFuncionarioByRut obtiene un funcionario por RUT
// @Summary Obtener funcionario por RUT
// @Description Obtiene un funcionario por su RUT
// @Tags funcionarios
// @Produce json
// @Param rut path string true "RUT del funcionario"
// @Success 200 {object} dto.FuncionarioResponse
// @Failure 404 {object} map[string]string
// @Router /api/v1/funcionarios/rut/{rut} [get]
func (h *FuncionarioHandler) GetFuncionarioByRut(c *gin.Context) {
	rut := c.Param("rut")
	
	funcionario, err := h.service.GetFuncionarioByRut(c.Request.Context(), rut)
	if err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "funcionario not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.ToFuncionarioResponse(funcionario))
}

// GetFuncionariosByEmpresa obtiene funcionarios por empresa
// @Summary Obtener funcionarios por empresa
// @Description Obtiene todos los funcionarios de una empresa
// @Tags funcionarios
// @Produce json
// @Param id_empresa path int true "ID de la empresa"
// @Success 200 {array} dto.FuncionarioResponse
// @Router /api/v1/funcionarios/empresa/{id_empresa} [get]
func (h *FuncionarioHandler) GetFuncionariosByEmpresa(c *gin.Context) {
	idEmpresa, err := strconv.Atoi(c.Param("id_empresa"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid empresa id"})
		return
	}

	funcionarios, err := h.service.GetFuncionariosByEmpresa(c.Request.Context(), idEmpresa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.ToFuncionarioResponseList(funcionarios))
}

// GetFuncionariosBySucursal obtiene funcionarios por sucursal
// @Summary Obtener funcionarios por sucursal
// @Description Obtiene todos los funcionarios de una sucursal
// @Tags funcionarios
// @Produce json
// @Param id_sucursal path int true "ID de la sucursal"
// @Success 200 {array} dto.FuncionarioResponse
// @Router /api/v1/funcionarios/sucursal/{id_sucursal} [get]
func (h *FuncionarioHandler) GetFuncionariosBySucursal(c *gin.Context) {
	idSucursal, err := strconv.Atoi(c.Param("id_sucursal"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sucursal id"})
		return
	}

	funcionarios, err := h.service.GetFuncionariosBySucursal(c.Request.Context(), idSucursal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.ToFuncionarioResponseList(funcionarios))
}

// GetFuncionariosBySegmento obtiene funcionarios por segmento
// @Summary Obtener funcionarios por segmento
// @Description Obtiene todos los funcionarios de un segmento
// @Tags funcionarios
// @Produce json
// @Param id_segmento path int true "ID del segmento"
// @Success 200 {array} dto.FuncionarioResponse
// @Router /api/v1/funcionarios/segmento/{id_segmento} [get]
func (h *FuncionarioHandler) GetFuncionariosBySegmento(c *gin.Context) {
	idSegmento, err := strconv.Atoi(c.Param("id_segmento"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid segmento id"})
		return
	}

	funcionarios, err := h.service.GetFuncionariosBySegmento(c.Request.Context(), idSegmento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.ToFuncionarioResponseList(funcionarios))
}

// ActivateFuncionario activa un funcionario
// @Summary Activar funcionario
// @Description Activa un funcionario previamente desactivado
// @Tags funcionarios
// @Param id path int true "ID del funcionario"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/funcionarios/{id}/activate [patch]
func (h *FuncionarioHandler) ActivateFuncionario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	if err := h.service.ActivateFuncionario(c.Request.Context(), id); err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "funcionario not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "funcionario activated successfully"})
}

// DeactivateFuncionario desactiva un funcionario
// @Summary Desactivar funcionario
// @Description Desactiva un funcionario sin eliminarlo
// @Tags funcionarios
// @Param id path int true "ID del funcionario"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/funcionarios/{id}/deactivate [patch]
func (h *FuncionarioHandler) DeactivateFuncionario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	if err := h.service.DeactivateFuncionario(c.Request.Context(), id); err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "funcionario not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "funcionario deactivated successfully"})
}

// CreateMedidas crea medidas para un funcionario
// @Summary Crear medidas
// @Description Registra las medidas corporales de un funcionario
// @Tags medidas
// @Accept json
// @Produce json
// @Param id path int true "ID del funcionario"
// @Param medidas body dto.CreateMedidasRequest true "Datos de las medidas"
// @Success 201 {object} dto.MedidasResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/funcionarios/{id}/medidas [post]
func (h *FuncionarioHandler) CreateMedidas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	var req dto.CreateMedidasRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medidas := req.ToMedidas()
	if err := h.service.CreateMedidas(c.Request.Context(), id, medidas); err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "funcionario not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, dto.ToMedidasResponse(medidas))
}

// GetMedidasActivas obtiene las medidas activas de un funcionario
// @Summary Obtener medidas activas
// @Description Obtiene las medidas corporales actuales de un funcionario
// @Tags medidas
// @Produce json
// @Param id path int true "ID del funcionario"
// @Success 200 {object} dto.MedidasResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/funcionarios/{id}/medidas [get]
func (h *FuncionarioHandler) GetMedidasActivas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	medidas, err := h.service.GetMedidasActivas(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "funcionario not found"})
			return
		}
		if errors.Is(err, services.ErrMedidasNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "medidas not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.ToMedidasResponse(medidas))
}

// GetHistorialMedidas obtiene el historial de medidas de un funcionario
// @Summary Obtener historial de medidas
// @Description Obtiene todo el historial de medidas corporales de un funcionario
// @Tags medidas
// @Produce json
// @Param id path int true "ID del funcionario"
// @Success 200 {array} dto.MedidasResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/funcionarios/{id}/medidas/historial [get]
func (h *FuncionarioHandler) GetHistorialMedidas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	medidas, err := h.service.GetHistorialMedidas(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "funcionario not found"})
			return
		}
		if errors.Is(err, services.ErrMedidasNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "medidas not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.ToMedidasResponseList(medidas))
}

// UpdateMedidas actualiza las medidas activas de un funcionario
// @Summary Actualizar medidas
// @Description Actualiza las medidas corporales actuales de un funcionario
// @Tags medidas
// @Accept json
// @Produce json
// @Param id path int true "ID del funcionario"
// @Param medidas body dto.UpdateMedidasRequest true "Datos actualizados"
// @Success 200 {object} dto.MedidasResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/funcionarios/{id}/medidas [put]
func (h *FuncionarioHandler) UpdateMedidas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	var req dto.UpdateMedidasRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medidas := &req
	medidasDomain := &services.FuncionarioService{}
	_ = medidasDomain // Placeholder para evitar error de compilación
	
	// Obtener medidas actuales para actualizar
	currentMedidas, err := h.service.GetMedidasActivas(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "funcionario not found"})
			return
		}
		if errors.Is(err, services.ErrMedidasNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "medidas not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Actualizar solo los campos proporcionados
	if medidas.EstaturaM != nil {
		currentMedidas.EstaturaM = medidas.EstaturaM
	}
	if medidas.PechoCm != nil {
		currentMedidas.PechoCm = medidas.PechoCm
	}
	if medidas.CinturaCm != nil {
		currentMedidas.CinturaCm = medidas.CinturaCm
	}
	if medidas.CaderaCm != nil {
		currentMedidas.CaderaCm = medidas.CaderaCm
	}
	if medidas.MangaCm != nil {
		currentMedidas.MangaCm = medidas.MangaCm
	}
	if medidas.FechaFin != nil {
		currentMedidas.FechaFin = medidas.FechaFin
	}

	if err := h.service.UpdateMedidas(c.Request.Context(), id, currentMedidas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.ToMedidasResponse(currentMedidas))
}

// HealthCheck verifica el estado del servicio
// @Summary Health Check
// @Description Verifica el estado del servicio
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *FuncionarioHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "ms-funcionario",
	})
}
