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
// @Description Crea un nuevo funcionario en el sistema (requiere rol admin)
// @Tags funcionarios
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param funcionario body dto.CreateFuncionarioRequest true "Datos del funcionario"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Router /api/v1/funcionarios [post]
func (h *FuncionarioHandler) CreateFuncionario(c *gin.Context) {
	var req dto.CreateFuncionarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Error en la validación", dto.FieldError{
			Field:   "request",
			Message: err.Error(),
		}))
		return
	}

	funcionario := req.ToFuncionario()
	if err := h.service.CreateFuncionario(c.Request.Context(), funcionario); err != nil {
		if errors.Is(err, services.ErrFuncionarioAlreadyExists) {
			// Determinar si es RUT o email duplicado
			fieldName := "rut_funcionario"
			if errors.Is(err, services.ErrFuncionarioAlreadyExists) && req.Email != "" {
				fieldName = "email"
			}
			c.JSON(http.StatusConflict, dto.NewErrorResponse("Error en la validación", dto.FieldError{
				Field:   fieldName,
				Message: err.Error(),
			}))
			return
		}
		if errors.Is(err, services.ErrInvalidInput) {
			c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Error en la validación", dto.FieldError{
				Field:   "request",
				Message: err.Error(),
			}))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusCreated, dto.NewSuccessResponseWithMessage(
		dto.ToFuncionarioResponse(funcionario),
		"Funcionario creado exitosamente",
	))
}

// GetFuncionario obtiene un funcionario por ID
// @Summary Obtener funcionario
// @Description Obtiene un funcionario por su ID
// @Tags funcionarios
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID del funcionario"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/funcionarios/{id} [get]
func (h *FuncionarioHandler) GetFuncionario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Formato de ID inválido"))
		return
	}

	funcionario, err := h.service.GetFuncionario(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("Funcionario no encontrado"))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(dto.ToFuncionarioResponse(funcionario)))
}

// GetAllFuncionarios obtiene todos los funcionarios con paginación y filtros
// @Summary Listar funcionarios
// @Description Obtiene la lista de funcionarios con paginación y filtros
// @Tags funcionarios
// @Security BearerAuth
// @Produce json
// @Param page query int false "Página" default(1)
// @Param limit query int false "Límite por página" default(20)
// @Param id_sucursal query int false "ID de la sucursal"
// @Param id_estado query int false "ID del estado"
// @Param id_genero query int false "ID del género"
// @Param id_cargo query int false "ID del cargo"
// @Param tallas_registradas query bool false "Tallas registradas"
// @Param search query string false "Buscar en nombres, apellidos, RUT, email"
// @Param sort_by query string false "Campo de ordenamiento" default("fecha_creación")
// @Param order query string false "Orden (asc|desc)" default("desc")
// @Success 200 {object} dto.PaginatedResponse
// @Router /api/v1/funcionarios [get]
func (h *FuncionarioHandler) GetAllFuncionarios(c *gin.Context) {
	var filterReq dto.FuncionarioFilterRequest
	if err := c.ShouldBindQuery(&filterReq); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Parámetros inválidos"))
		return
	}

	// Valores por defecto
	if filterReq.Page <= 0 {
		filterReq.Page = 1
	}
	if filterReq.Limit <= 0 {
		filterReq.Limit = 20
	}
	if filterReq.Limit > 100 {
		filterReq.Limit = 100 // Máximo 100 según contrato
	}
	if filterReq.SortBy == "" {
		filterReq.SortBy = "fecha_creación"
	}
	if filterReq.Order == "" {
		filterReq.Order = "desc"
	}

	// Calcular offset desde page
	filterReq.Offset = (filterReq.Page - 1) * filterReq.Limit

	filter := filterReq.ToFilter()

	funcionarios, err := h.service.GetFuncionariosByFilter(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	total, err := h.service.CountFuncionarios(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewPaginatedResponse(
		dto.ToFuncionarioResponseList(funcionarios),
		filterReq.Page,
		filterReq.Limit,
		total,
	))
}

// GetFuncionariosByFilter obtiene funcionarios con filtros
// @Summary Filtrar funcionarios
// @Description Obtiene funcionarios según criterios de filtrado con paginación
// @Tags funcionarios
// @Security BearerAuth
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
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Parámetros inválidos"))
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
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	total, err := h.service.CountFuncionarios(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
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
// @Description Actualiza los datos de un funcionario existente (requiere rol admin)
// @Tags funcionarios
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID del funcionario"
// @Param funcionario body dto.UpdateFuncionarioRequest true "Datos actualizados"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/funcionarios/{id} [put]
func (h *FuncionarioHandler) UpdateFuncionario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Formato de ID inválido"))
		return
	}

	var req dto.UpdateFuncionarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Error en la validación", dto.FieldError{
			Field:   "request",
			Message: err.Error(),
		}))
		return
	}

	funcionario := req.ToFuncionario(id)
	if err := h.service.UpdateFuncionario(c.Request.Context(), funcionario); err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("Funcionario no encontrado"))
			return
		}
		if errors.Is(err, services.ErrInvalidInput) {
			c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Error en la validación", dto.FieldError{
				Field:   "request",
				Message: err.Error(),
			}))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponseWithMessage(
		dto.ToFuncionarioResponse(funcionario),
		"Funcionario actualizado exitosamente",
	))
}

// DeleteFuncionario elimina un funcionario (soft delete)
// @Summary Eliminar funcionario
// @Description Desactiva un funcionario del sistema - eliminación lógica (requiere rol admin)
// @Tags funcionarios
// @Security BearerAuth
// @Param id path int true "ID del funcionario"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/funcionarios/{id} [delete]
func (h *FuncionarioHandler) DeleteFuncionario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Formato de ID inválido"))
		return
	}

	// Usar DeactivateFuncionario en lugar de Delete para soft delete
	if err := h.service.DeactivateFuncionario(c.Request.Context(), id); err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("Funcionario no encontrado"))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponseWithMessage(nil, "Funcionario desactivado exitosamente"))
}

// GetFuncionarioByRut obtiene un funcionario por RUT
// @Summary Obtener funcionario por RUT
// @Description Obtiene un funcionario por su RUT
// @Tags funcionarios
// @Security BearerAuth
// @Produce json
// @Param rut path string true "RUT del funcionario"
// @Success 200 {object} dto.SuccessResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/funcionarios/buscar/rut/{rut} [get]
func (h *FuncionarioHandler) GetFuncionarioByRut(c *gin.Context) {
	rut := c.Param("rut")

	funcionario, err := h.service.GetFuncionarioByRut(c.Request.Context(), rut)
	if err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("Funcionario no encontrado"))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(dto.ToFuncionarioResponse(funcionario)))
}

// GetFuncionariosByEmpresa obtiene funcionarios por empresa
// @Summary Obtener funcionarios por empresa
// @Description Obtiene todos los funcionarios de una empresa
// @Tags funcionarios
// @Security BearerAuth
// @Produce json
// @Param id_empresa path int true "ID de la empresa"
// @Success 200 {object} dto.SuccessResponse
// @Router /api/v1/funcionarios/empresa/{id_empresa} [get]
func (h *FuncionarioHandler) GetFuncionariosByEmpresa(c *gin.Context) {
	idEmpresa, err := strconv.Atoi(c.Param("id_empresa"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Formato de ID de empresa inválido"))
		return
	}

	funcionarios, err := h.service.GetFuncionariosByEmpresa(c.Request.Context(), idEmpresa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(dto.ToFuncionarioResponseList(funcionarios)))
}

// GetFuncionariosBySucursal obtiene funcionarios por sucursal
// @Summary Obtener funcionarios por sucursal
// @Description Obtiene todos los funcionarios de una sucursal
// @Tags funcionarios
// @Security BearerAuth
// @Produce json
// @Param id_sucursal path int true "ID de la sucursal"
// @Success 200 {object} dto.SuccessResponse
// @Router /api/v1/funcionarios/sucursal/{id_sucursal} [get]
func (h *FuncionarioHandler) GetFuncionariosBySucursal(c *gin.Context) {
	idSucursal, err := strconv.Atoi(c.Param("id_sucursal"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Formato de ID de sucursal inválido"))
		return
	}

	funcionarios, err := h.service.GetFuncionariosBySucursal(c.Request.Context(), idSucursal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(dto.ToFuncionarioResponseList(funcionarios)))
}

// GetFuncionariosBySegmento obtiene funcionarios por segmento
// @Summary Obtener funcionarios por segmento
// @Description Obtiene todos los funcionarios de un segmento
// @Tags funcionarios
// @Security BearerAuth
// @Produce json
// @Param id_segmento path int true "ID del segmento"
// @Success 200 {object} dto.SuccessResponse
// @Router /api/v1/funcionarios/segmento/{id_segmento} [get]
func (h *FuncionarioHandler) GetFuncionariosBySegmento(c *gin.Context) {
	idSegmento, err := strconv.Atoi(c.Param("id_segmento"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Formato de ID de segmento inválido"))
		return
	}

	funcionarios, err := h.service.GetFuncionariosBySegmento(c.Request.Context(), idSegmento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(dto.ToFuncionarioResponseList(funcionarios)))
}

// ActivateFuncionario activa un funcionario
// @Summary Activar funcionario
// @Description Activa un funcionario previamente desactivado (requiere rol admin)
// @Tags funcionarios
// @Security BearerAuth
// @Param id path int true "ID del funcionario"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/funcionarios/{id}/activate [patch]
func (h *FuncionarioHandler) ActivateFuncionario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Formato de ID inválido"))
		return
	}

	if err := h.service.ActivateFuncionario(c.Request.Context(), id); err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("Funcionario no encontrado"))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponseWithMessage(nil, "Funcionario activado exitosamente"))
}

// DeactivateFuncionario desactiva un funcionario
// @Summary Desactivar funcionario
// @Description Desactiva un funcionario sin eliminarlo (requiere rol admin)
// @Tags funcionarios
// @Security BearerAuth
// @Param id path int true "ID del funcionario"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/funcionarios/{id}/deactivate [patch]
func (h *FuncionarioHandler) DeactivateFuncionario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Formato de ID inválido"))
		return
	}

	if err := h.service.DeactivateFuncionario(c.Request.Context(), id); err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("Funcionario no encontrado"))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponseWithMessage(nil, "Funcionario desactivado exitosamente"))
}

// CreateMedidas crea medidas para un funcionario
// @Summary Crear medidas
// @Description Registra las medidas corporales de un funcionario
// @Tags medidas
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID del funcionario"
// @Param medidas body dto.CreateMedidasRequest true "Datos de las medidas"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/funcionarios/{id}/medidas [post]
func (h *FuncionarioHandler) CreateMedidas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Formato de ID inválido"))
		return
	}

	var req dto.CreateMedidasRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Error en la validación", dto.FieldError{
			Field:   "request",
			Message: err.Error(),
		}))
		return
	}

	medidas := req.ToMedidas()
	if err := h.service.CreateMedidas(c.Request.Context(), id, medidas); err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("Funcionario no encontrado"))
			return
		}
		if errors.Is(err, services.ErrInvalidInput) {
			c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Error en la validación", dto.FieldError{
				Field:   "medidas",
				Message: err.Error(),
			}))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusCreated, dto.NewSuccessResponseWithMessage(
		dto.ToMedidasResponse(medidas, id),
		"Medidas registradas exitosamente",
	))
}

// GetMedidasActivas obtiene las medidas activas de un funcionario
// @Summary Obtener medidas activas
// @Description Obtiene las medidas corporales actuales de un funcionario
// @Tags medidas
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID del funcionario"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/funcionarios/{id}/medidas [get]
func (h *FuncionarioHandler) GetMedidasActivas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Formato de ID inválido"))
		return
	}

	medidas, err := h.service.GetMedidasActivas(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("Funcionario no encontrado"))
			return
		}
		if errors.Is(err, services.ErrMedidasNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("El funcionario no tiene medidas registradas"))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(dto.ToMedidasResponse(medidas, id)))
}

// GetHistorialMedidas obtiene el historial de medidas de un funcionario
// @Summary Obtener historial de medidas
// @Description Obtiene todo el historial de medidas corporales de un funcionario
// @Tags medidas
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID del funcionario"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/funcionarios/{id}/medidas/historial [get]
func (h *FuncionarioHandler) GetHistorialMedidas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Formato de ID inválido"))
		return
	}

	medidas, err := h.service.GetHistorialMedidas(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("Funcionario no encontrado"))
			return
		}
		if errors.Is(err, services.ErrMedidasNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("No se encontraron medidas"))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(dto.ToMedidasResponseList(medidas)))
}

// UpdateMedidas actualiza las medidas activas de un funcionario
// @Summary Actualizar medidas
// @Description Actualiza las medidas corporales actuales de un funcionario
// @Tags medidas
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID del funcionario"
// @Param medidas body dto.UpdateMedidasRequest true "Datos actualizados"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/funcionarios/{id}/medidas [put]
func (h *FuncionarioHandler) UpdateMedidas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Formato de ID inválido"))
		return
	}

	var req dto.UpdateMedidasRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Error en la validación", dto.FieldError{
			Field:   "request",
			Message: err.Error(),
		}))
		return
	}

	// Obtener medidas actuales para actualizar
	currentMedidas, err := h.service.GetMedidasActivas(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrFuncionarioNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("Funcionario no encontrado"))
			return
		}
		if errors.Is(err, services.ErrMedidasNotFound) {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse("No se encontraron medidas"))
			return
		}
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	// Actualizar solo los campos proporcionados
	if req.EstaturaM != nil {
		currentMedidas.EstaturaM = req.EstaturaM
	}
	if req.PechoCm != nil {
		currentMedidas.PechoCm = req.PechoCm
	}
	if req.CinturaCm != nil {
		currentMedidas.CinturaCm = req.CinturaCm
	}
	if req.CaderaCm != nil {
		currentMedidas.CaderaCm = req.CaderaCm
	}
	if req.MangaCm != nil {
		currentMedidas.MangaCm = req.MangaCm
	}
	if req.FechaFin != nil {
		currentMedidas.FechaFin = req.FechaFin
	}

	if err := h.service.UpdateMedidas(c.Request.Context(), id, currentMedidas); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(dto.ToMedidasResponse(currentMedidas, id)))
}

// HealthCheck verifica el estado del servicio
// @Summary Health Check
// @Description Verifica el estado del servicio
// @Tags health
// @Produce json
// @Success 200 {object} dto.SuccessResponse
// @Router /health [get]
func (h *FuncionarioHandler) HealthCheck(c *gin.Context) {
	healthData := map[string]string{
		"status":  "ok",
		"service": "ms-funcionario",
	}
	c.JSON(http.StatusOK, dto.NewSuccessResponse(healthData))
}

// ListBranches obtiene la lista de sucursales
// @Summary Listar sucursales
// @Description Obtiene todas las sucursales disponibles
// @Tags sucursales
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.SuccessResponse
// @Router /api/v1/branches [get]
func (h *FuncionarioHandler) ListBranches(c *gin.Context) {
	branches, err := h.service.ListBranches(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	// Map to DTO
	dtos := make([]dto.SucursalDTO, len(branches))
	for i, b := range branches {
		dtos[i] = dto.SucursalDTO{
			IDSucursal:     b.IDSucursal,
			NombreSucursal: b.NombreSucursal,
			Direccion:      b.Direccion,
		}
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(dtos))
}

// RequestTransfer solicita un traslado
// @Summary Solicitar traslado
// @Description Solicita un traslado de sucursal para el funcionario autenticado
// @Tags traslados
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body dto.TransferRequest true "Datos traslado"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /api/v1/transfers [post]
func (h *FuncionarioHandler) RequestTransfer(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("Usuario no autenticado"))
		return
	}

	var req dto.TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Error en la validación", dto.FieldError{
			Field:   "request",
			Message: err.Error(),
		}))
		return
	}

	if err := h.service.RequestTransfer(c.Request.Context(), userID.(int), req.TargetBranchID, req.Reason); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Error interno del servidor"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponseWithMessage(nil, "Traslado solicitado exitosamente"))
}
