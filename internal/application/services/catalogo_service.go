package services

import (
	"context"

	"github.com/Sistal/ms-funcionario/internal/domain/funcionario"
)

type CatalogoService struct {
	cargoRepo  funcionario.CargoRepository
	generoRepo funcionario.GeneroRepository
}

func NewCatalogoService(cargoRepo funcionario.CargoRepository, generoRepo funcionario.GeneroRepository) *CatalogoService {
	return &CatalogoService{
		cargoRepo:  cargoRepo,
		generoRepo: generoRepo,
	}
}

// ListCargos obtiene todos los cargos
func (s *CatalogoService) ListCargos(ctx context.Context) ([]*funcionario.Cargo, error) {
	return s.cargoRepo.FindAll(ctx)
}

// ListGeneros obtiene todos los géneros
func (s *CatalogoService) ListGeneros(ctx context.Context) ([]*funcionario.Genero, error) {
	return s.generoRepo.FindAll(ctx)
}
