package handler

import (
	"net/http"
	"strconv"

	"github.com/Sistal/ms-funcionario/internal/application/services"
	"github.com/Sistal/ms-funcionario/internal/interfaces/dto"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	funcionarioService *services.FuncionarioService
}

func NewProfileHandler(funcionarioService *services.FuncionarioService) *ProfileHandler {
	return &ProfileHandler{
		funcionarioService: funcionarioService,
	}
}

// Helper para obtener UserID del path o del token
func (h *ProfileHandler) getUserID(c *gin.Context) (int, error) {
	userIdParam := c.Param("userId")

	// Si el parámetro está vacío o es explícitamente "me", usar token
	if userIdParam == "" || userIdParam == "me" {
		uid, exists := c.Get("user_id")
		if !exists {
			return 0, http.ErrNoCookie
		}
		return uid.(int), nil
	}

	// Intentar parsear el ID de la URL
	id, err := strconv.Atoi(userIdParam)
	if err != nil {
		return 0, strconv.ErrSyntax
	}
	return id, nil
}

// GetProfile obtiene el perfil de un funcionario
// @Summary Obtener perfil
// @Description Obtiene los datos del funcionario por User ID (o "me")
// @Tags employees
// @Security BearerAuth
// @Produce json
// @Param userId path string true "User ID o 'me'"
// @Success 200 {object} dto.ProfileResponse
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/employees/{userId}/profile [get]
func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		}
		return
	}

	// Buscar el funcionario por ID de usuario
	f, err := h.funcionarioService.GetFuncionarioByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "funcionario profile not found"})
		return
	}

	c.JSON(http.StatusOK, dto.ToProfileResponse(f))
}

// UpdateMyProfile (Legacy) - Alias a GetProfile (dado que era un PUT que no hacía update real en el código anterior,
// o si se requiere mantener el método, lo redirigimos)
func (h *ProfileHandler) UpdateMyProfile(c *gin.Context) {
	// El router legacy llama a esto como PUT.
	// Asumimos que es "GetProfile" o "NoOp" por ahora según el estado previo.
	h.GetProfile(c)
}

// UpdateContact actualiza datos de contacto
// @Summary Actualizar contacto
// @Tags employees
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param userId path string true "User ID o 'me'"
// @Param data body dto.ContactUpdateRequest true "Datos de contacto"
// @Success 200 {object} dto.ProfileResponse
// @Router /api/v1/employees/{userId}/contact [put]
func (h *ProfileHandler) UpdateContact(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		}
		return
	}

	var req dto.ContactUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.funcionarioService.UpdateContact(c.Request.Context(), userID, req.Email, req.Celular); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	f, _ := h.funcionarioService.GetFuncionarioByUserID(c.Request.Context(), userID)
	c.JSON(http.StatusOK, dto.ToProfileResponse(f))
}

// GetMeasurements obtiene las medidas
// @Summary Obtener medidas
// @Tags employees
// @Security BearerAuth
// @Produce json
// @Param userId path string true "User ID o 'me'"
// @Success 200 {object} dto.MedidasResponse
// @Router /api/v1/employees/{userId}/measurements [get]
func (h *ProfileHandler) GetMeasurements(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		}
		return
	}

	f, err := h.funcionarioService.GetFuncionarioByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "funcionario not found"})
		return
	}

	if f.Medidas != nil {
		c.JSON(http.StatusOK, dto.ToMedidasResponse(f.Medidas))
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "measurements not found"})
	}
}

// UpdateMeasurements actualiza las medidas
// @Summary Actualizar medidas
// @Tags employees
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param userId path string true "User ID o 'me'"
// @Param data body dto.CreateMedidasRequest true "Medidas"
// @Success 200 {object} dto.ProfileResponse
// @Router /api/v1/employees/{userId}/measurements [put]
func (h *ProfileHandler) UpdateMeasurements(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		}
		return
	}

	var req dto.CreateMedidasRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medidas := req.ToMedidas()

	if err := h.funcionarioService.ManageMeasurements(c.Request.Context(), userID, medidas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	f, _ := h.funcionarioService.GetFuncionarioByUserID(c.Request.Context(), userID)
	c.JSON(http.StatusOK, dto.ToProfileResponse(f))
}

// GetMyStats obtiene las estadísticas del dashboard
// @Summary Obtener estadísticas
// @Description Obtiene las estadísticas del funcionario autenticado
// @Tags funcionarios
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /api/v1/funcionarios/me/stats [get]
func (h *ProfileHandler) GetMyStats(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":                userID,
		"total_solicitudes":      5,
		"solicitudes_pendientes": 2,
		"entregas_proximas":      3,
	})
}
