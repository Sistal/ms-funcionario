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
	repo         funcionario.Repository
	medidasRepo  funcionario.MedidasRepository
	sucursalRepo funcionario.SucursalRepository
}

func NewFuncionarioService(repo funcionario.Repository, medidasRepo funcionario.MedidasRepository, sucursalRepo funcionario.SucursalRepository) *FuncionarioService {
	return &FuncionarioService{
		repo:         repo,
		medidasRepo:  medidasRepo,
		sucursalRepo: sucursalRepo,
	}
}

// CreateFuncionario crea un nuevo funcionario
func (s *FuncionarioService) CreateFuncionario(ctx context.Context, f *funcionario.Funcionario) error {
	// Validar campos usando el m\u00e9todo Validate del dominio
	if err := f.Validate(); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidInput, err)
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

	// Estado activo por defecto (10 seg\u00fan contrato)
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

// GetFuncionarioByUserID obtiene un funcionario por ID de usuario
func (s *FuncionarioService) GetFuncionarioByUserID(ctx context.Context, userID int) (*funcionario.Funcionario, error) {
	f, err := s.repo.GetByUserID(ctx, userID)
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

	// 1. Obtener el funcionario existente
	existing, err := s.repo.GetByID(ctx, f.IDFuncionario)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || err == ErrFuncionarioNotFound {
			return ErrFuncionarioNotFound
		}
		// Mapeo de errores de DB a errores de dominio si es necesario
		return err
	}
	if existing == nil {
		return ErrFuncionarioNotFound
	}

	// 2. Merge de campos (Solo actualizamos lo que viene con valor)
	// Strings (asumimos que vacío significa no cambiar)
	if f.RutFuncionario != "" {
		existing.RutFuncionario = f.RutFuncionario
	}
	if f.Nombres != "" {
		existing.Nombres = f.Nombres
	}
	if f.ApellidoPaterno != "" {
		existing.ApellidoPaterno = f.ApellidoPaterno
	}
	if f.ApellidoMaterno != "" {
		existing.ApellidoMaterno = f.ApellidoMaterno
	}
	if f.Celular != "" {
		existing.Celular = f.Celular
	}
	if f.Telefono != "" {
		existing.Telefono = f.Telefono
	}
	if f.Email != "" {
		existing.Email = f.Email
	}
	if f.Direccion != "" {
		existing.Direccion = f.Direccion
	}

	// IDs / FKs (Punteros) - Si no son nil, actualizamos
	if f.IDGenero != nil {
		existing.IDGenero = f.IDGenero
	}
	if f.IDEstado != nil {
		existing.IDEstado = f.IDEstado
	}
	if f.IDSucursal != nil {
		existing.IDSucursal = f.IDSucursal
	}
	if f.IDCargo != nil {
		existing.IDCargo = f.IDCargo
	}
	if f.IDEmpresaCliente != nil {
		existing.IDEmpresaCliente = f.IDEmpresaCliente
	}
	if f.IDSegmento != nil {
		existing.IDSegmento = f.IDSegmento
	}
	if f.IDUsuario != nil {
		existing.IDUsuario = f.IDUsuario
	}
	if f.IDMedidas != nil {
		existing.IDMedidas = f.IDMedidas
	}

	// Actualizar fecha modificación
	now := time.Now()
	existing.FechaModificacion = &now

	// 3. Validar el objeto final antes de guardar
	if err := existing.Validate(); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	// 4. Guardar cambios
	return s.repo.Update(ctx, existing)
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
	// Validar rangos de medidas
	if err := funcionario.ValidateMedidas(medidas); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

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
	return s.repo.UpdateMedidasInfo(ctx, idFuncionario, medidas.IDMedidas)
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
	// Validar rangos de medidas
	if err := funcionario.ValidateMedidas(medidas); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

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

// ListBranches returns all branches
func (s *FuncionarioService) ListBranches(ctx context.Context) ([]*funcionario.Sucursal, error) {
	return s.sucursalRepo.FindAll(ctx)
}

// UpdateContact updates only contact info
func (s *FuncionarioService) UpdateContact(ctx context.Context, userID int, email, celular string) error {
	f, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if f == nil {
		return ErrFuncionarioNotFound
	}

	f.Email = email
	f.Celular = celular

	now := time.Now()
	f.FechaModificacion = &now
	return s.repo.Update(ctx, f)
}

// RequestTransfer updates branch directly (temporary implementation)
func (s *FuncionarioService) RequestTransfer(ctx context.Context, userID int, targetBranchID int, reason string) error {
	f, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if f == nil {
		return ErrFuncionarioNotFound
	}

	// Verify target branch exists
	_, err = s.sucursalRepo.GetByID(ctx, targetBranchID)
	if err != nil {
		return fmt.Errorf("branch error: %w", err)
	}

	// Direct update as per temporary requirement
	f.IDSucursal = &targetBranchID

	now := time.Now()
	f.FechaModificacion = &now
	return s.repo.Update(ctx, f)
}

// ManageMeasurements creates or updates measurements for user ID
func (s *FuncionarioService) ManageMeasurements(ctx context.Context, userID int, medidas *funcionario.MedidasFuncionario) error {
	f, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if f == nil {
		return ErrFuncionarioNotFound
	}

	if f.IDMedidas == nil {
		// Create new
		return s.CreateMedidas(ctx, f.IDFuncionario, medidas)
	} else {
		// Update existing
		return s.UpdateMedidas(ctx, f.IDFuncionario, medidas)
	}
}
