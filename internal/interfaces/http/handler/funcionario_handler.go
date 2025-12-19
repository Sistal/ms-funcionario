package handler

import (
	"errors"
	"net/http"

	"github.com/Sistal/ms-funcionario/internal/application/services"
	"github.com/Sistal/ms-funcionario/internal/interfaces/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FuncionarioHandler struct {
	service *services.FuncionarioService
}

func NewFuncionarioHandler(service *services.FuncionarioService) *FuncionarioHandler {
	return &FuncionarioHandler{
		service: service,
	}
}

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

func (h *FuncionarioHandler) GetFuncionario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
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

func (h *FuncionarioHandler) GetAllFuncionarios(c *gin.Context) {
	funcionarios, err := h.service.GetAllFuncionarios(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.ToFuncionarioResponseList(funcionarios))
}

func (h *FuncionarioHandler) UpdateFuncionario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
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

func (h *FuncionarioHandler) DeleteFuncionario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
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

func (h *FuncionarioHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"service": "ms-funcionario",
	})
}
