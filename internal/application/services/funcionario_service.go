package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Sistal/ms-funcionario/internal/domain/funcionario"
	"gorm.io/gorm"
)

var (
	ErrFuncionarioNotFound      = errors.New("funcionario not found")
	ErrFuncionarioAlreadyExists = errors.New("funcionario already exists")
	ErrInvalidInput             = errors.New("invalid input")
	ErrMedidasNotFound          = errors.New("medidas not found")
	ErrMedidasAlreadyExists     = errors.New("medidas already exist for this funcionario")
)

type FuncionarioService struct {
	repo        funcionario.Repository
	medidasRepo funcionario.MedidasRepository
}

func NewFuncionarioService(repo funcionario.Repository, medidasRepo funcionario.MedidasRepository) *FuncionarioService {
	return &FuncionarioService{
		repo:        repo,
		medidasRepo: medidasRepo,
	}
}

// CreateFuncionario crea un nuevo funcionario
func (s *FuncionarioService) CreateFuncionario(ctx context.Context, f *funcionario.Funcionario) error {
	if f.Nombres == "" || f.ApellidoPaterno == "" || f.Email == "" || f.RutFuncionario == "" {
		return ErrInvalidInput
	}

	// Verificar que no exista por RUT
	existing, _ := s.repo.GetByRut(ctx, f.RutFuncionario)
	if existing != nil {
		return fmt.Errorf("%w: rut %s already exists", ErrFuncionarioAlreadyExists, f.RutFuncionario)
	}

	// Verificar que no exista por email
	existingEmail, _ := s.repo.GetByEmail(ctx, f.Email)
	if existingEmail != nil {
		return fmt.Errorf("%w: email %s already exists", ErrFuncionarioAlreadyExists, f.Email)
	}

	// Establecer valores por defecto
	now := time.Now()
	f.FechaCreacion = &now
	f.TallasRegistradas = false
	
	// Establecer estado activo por defecto (1 = activo)
	if f.IDEstado == nil {
		estado := 1
		f.IDEstado = &estado
	}

	return s.repo.Create(ctx, f)
}

// GetFuncionario obtiene un funcionario por ID
func (s *FuncionarioService) GetFuncionario(ctx context.Context, id int) (*funcionario.Funcionario, error) {
	f, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFuncionarioNotFound
		}
		return nil, err
	}
	return f, nil
}

// GetAllFuncionarios obtiene todos los funcionarios
func (s *FuncionarioService) GetAllFuncionarios(ctx context.Context) ([]*funcionario.Funcionario, error) {
	return s.repo.GetAll(ctx)
}

// GetFuncionariosByFilter obtiene funcionarios según filtros
func (s *FuncionarioService) GetFuncionariosByFilter(ctx context.Context, filter funcionario.FuncionarioFilter) ([]*funcionario.Funcionario, error) {
	return s.repo.GetByFilter(ctx, filter)
}

// CountFuncionarios cuenta funcionarios según filtros
func (s *FuncionarioService) CountFuncionarios(ctx context.Context, filter funcionario.FuncionarioFilter) (int64, error) {
	return s.repo.Count(ctx, filter)
}

// UpdateFuncionario actualiza un funcionario existente
func (s *FuncionarioService) UpdateFuncionario(ctx context.Context, f *funcionario.Funcionario) error {
	if f.IDFuncionario == 0 {
		return ErrInvalidInput
	}

	_, err := s.repo.GetByID(ctx, f.IDFuncionario)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrFuncionarioNotFound
		}
		return err
	}

	now := time.Now()
	f.FechaModificacion = &now

	return s.repo.Update(ctx, f)
}

// DeleteFuncionario elimina un funcionario
func (s *FuncionarioService) DeleteFuncionario(ctx context.Context, id int) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrFuncionarioNotFound
		}
		return err
	}

	return s.repo.Delete(ctx, id)
}

// GetFuncionarioByRut obtiene un funcionario por RUT
func (s *FuncionarioService) GetFuncionarioByRut(ctx context.Context, rut string) (*funcionario.Funcionario, error) {
	f, err := s.repo.GetByRut(ctx, rut)
	if err != nil {
		return nil, err
	}
	if f == nil {
		return nil, ErrFuncionarioNotFound
	}
	return f, nil
}

// GetFuncionariosByEmpresa obtiene funcionarios por empresa
func (s *FuncionarioService) GetFuncionariosByEmpresa(ctx context.Context, idEmpresa int) ([]*funcionario.Funcionario, error) {
	return s.repo.GetByEmpresa(ctx, idEmpresa)
}

// GetFuncionariosBySucursal obtiene funcionarios por sucursal
func (s *FuncionarioService) GetFuncionariosBySucursal(ctx context.Context, idSucursal int) ([]*funcionario.Funcionario, error) {
	return s.repo.GetBySucursal(ctx, idSucursal)
}

// GetFuncionariosBySegmento obtiene funcionarios por segmento
func (s *FuncionarioService) GetFuncionariosBySegmento(ctx context.Context, idSegmento int) ([]*funcionario.Funcionario, error) {
	return s.repo.GetBySegmento(ctx, idSegmento)
}

// ActivateFuncionario activa un funcionario
func (s *FuncionarioService) ActivateFuncionario(ctx context.Context, id int) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrFuncionarioNotFound
		}
		return err
	}

	return s.repo.ActivateByID(ctx, id)
}

// DeactivateFuncionario desactiva un funcionario
func (s *FuncionarioService) DeactivateFuncionario(ctx context.Context, id int) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrFuncionarioNotFound
		}
		return err
	}

	return s.repo.DeactivateByID(ctx, id)
}

// CreateMedidas crea nuevas medidas para un funcionario
func (s *FuncionarioService) CreateMedidas(ctx context.Context, idFuncionario int, medidas *funcionario.MedidasFuncionario) error {
	// Verificar que el funcionario existe
	f, err := s.repo.GetByID(ctx, idFuncionario)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrFuncionarioNotFound
		}
		return err
	}

	// Si el funcionario ya tiene medidas activas, cerrarlas
	if f.IDMedidas != nil {
		existingMedidas, err := s.medidasRepo.GetByID(ctx, *f.IDMedidas)
		if err == nil && existingMedidas.FechaFin == nil {
			now := time.Now()
			existingMedidas.FechaFin = &now
			if err := s.medidasRepo.Update(ctx, existingMedidas); err != nil {
				return err
			}
		}
	}

	// Establecer fecha de inicio si no existe
	if medidas.FechaInicio == nil {
		now := time.Now()
		medidas.FechaInicio = &now
	}

	// Crear las nuevas medidas
	if err := s.medidasRepo.Create(ctx, medidas); err != nil {
		return err
	}

	// Actualizar el funcionario con el ID de las nuevas medidas
	f.IDMedidas = &medidas.IDMedidas
	f.TallasRegistradas = true
	now := time.Now()
	f.FechaModificacion = &now

	return s.repo.Update(ctx, f)
}

// GetMedidasActivas obtiene las medidas activas de un funcionario
func (s *FuncionarioService) GetMedidasActivas(ctx context.Context, idFuncionario int) (*funcionario.MedidasFuncionario, error) {
	f, err := s.repo.GetByID(ctx, idFuncionario)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFuncionarioNotFound
		}
		return nil, err
	}

	if f.IDMedidas == nil {
		return nil, ErrMedidasNotFound
	}

	medidas, err := s.medidasRepo.GetActivasByFuncionario(ctx, *f.IDMedidas)
	if err != nil {
		return nil, err
	}

	return medidas, nil
}

// GetHistorialMedidas obtiene el historial de medidas de un funcionario
func (s *FuncionarioService) GetHistorialMedidas(ctx context.Context, idFuncionario int) ([]*funcionario.MedidasFuncionario, error) {
	f, err := s.repo.GetByID(ctx, idFuncionario)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFuncionarioNotFound
		}
		return nil, err
	}

	if f.IDMedidas == nil {
		return nil, ErrMedidasNotFound
	}

	return s.medidasRepo.GetHistorialByFuncionario(ctx, *f.IDMedidas)
}

// UpdateMedidas actualiza las medidas activas de un funcionario
func (s *FuncionarioService) UpdateMedidas(ctx context.Context, idFuncionario int, medidas *funcionario.MedidasFuncionario) error {
	f, err := s.repo.GetByID(ctx, idFuncionario)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrFuncionarioNotFound
		}
		return err
	}

	if f.IDMedidas == nil {
		return ErrMedidasNotFound
	}

	medidas.IDMedidas = *f.IDMedidas
	return s.medidasRepo.Update(ctx, medidas)
}
