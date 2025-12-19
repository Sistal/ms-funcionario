package services

import (
	"context"
	"errors"

	"github.com/Sistal/ms-funcionario/internal/domain/funcionario"
	"github.com/google/uuid"
)

var (
	ErrFuncionarioNotFound     = errors.New("funcionario not found")
	ErrFuncionarioAlreadyExists = errors.New("funcionario already exists")
	ErrInvalidInput            = errors.New("invalid input")
)

type FuncionarioService struct {
	repo funcionario.Repository
}

func NewFuncionarioService(repo funcionario.Repository) *FuncionarioService {
	return &FuncionarioService{
		repo: repo,
	}
}

func (s *FuncionarioService) CreateFuncionario(ctx context.Context, f *funcionario.Funcionario) error {
	if f.Nombre == "" || f.Apellido == "" || f.Email == "" {
		return ErrInvalidInput
	}

	existing, _ := s.repo.GetByEmail(ctx, f.Email)
	if existing != nil {
		return ErrFuncionarioAlreadyExists
	}

	f.ID = uuid.New()
	f.Activo = true

	return s.repo.Create(ctx, f)
}

func (s *FuncionarioService) GetFuncionario(ctx context.Context, id uuid.UUID) (*funcionario.Funcionario, error) {
	f, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrFuncionarioNotFound
	}
	return f, nil
}

func (s *FuncionarioService) GetAllFuncionarios(ctx context.Context) ([]*funcionario.Funcionario, error) {
	return s.repo.GetAll(ctx)
}

func (s *FuncionarioService) UpdateFuncionario(ctx context.Context, f *funcionario.Funcionario) error {
	if f.ID == uuid.Nil {
		return ErrInvalidInput
	}

	existing, err := s.repo.GetByID(ctx, f.ID)
	if err != nil || existing == nil {
		return ErrFuncionarioNotFound
	}

	return s.repo.Update(ctx, f)
}

func (s *FuncionarioService) DeleteFuncionario(ctx context.Context, id uuid.UUID) error {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil || existing == nil {
		return ErrFuncionarioNotFound
	}

	return s.repo.Delete(ctx, id)
}
