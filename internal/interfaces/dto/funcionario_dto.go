package dto

import (
	"time"

	"github.com/Sistal/ms-funcionario/internal/domain/funcionario"
	"github.com/google/uuid"
)

type CreateFuncionarioRequest struct {
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Cargo    string `json:"cargo"`
}

type UpdateFuncionarioRequest struct {
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Cargo    string `json:"cargo"`
	Activo   bool   `json:"activo"`
}

type FuncionarioResponse struct {
	ID        uuid.UUID `json:"id"`
	Nombre    string    `json:"nombre"`
	Apellido  string    `json:"apellido"`
	Email     string    `json:"email"`
	Cargo     string    `json:"cargo"`
	Activo    bool      `json:"activo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToFuncionarioResponse(f *funcionario.Funcionario) *FuncionarioResponse {
	return &FuncionarioResponse{
		ID:        f.ID,
		Nombre:    f.Nombre,
		Apellido:  f.Apellido,
		Email:     f.Email,
		Cargo:     f.Cargo,
		Activo:    f.Activo,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

func ToFuncionarioResponseList(funcionarios []*funcionario.Funcionario) []*FuncionarioResponse {
	responses := make([]*FuncionarioResponse, len(funcionarios))
	for i, f := range funcionarios {
		responses[i] = ToFuncionarioResponse(f)
	}
	return responses
}

func (req *CreateFuncionarioRequest) ToFuncionario() *funcionario.Funcionario {
	return &funcionario.Funcionario{
		Nombre:   req.Nombre,
		Apellido: req.Apellido,
		Email:    req.Email,
		Cargo:    req.Cargo,
	}
}

func (req *UpdateFuncionarioRequest) ToFuncionario(id uuid.UUID) *funcionario.Funcionario {
	return &funcionario.Funcionario{
		ID:       id,
		Nombre:   req.Nombre,
		Apellido: req.Apellido,
		Email:    req.Email,
		Cargo:    req.Cargo,
		Activo:   req.Activo,
	}
}
