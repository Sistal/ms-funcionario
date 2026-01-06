package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Sistal/ms-funcionario/internal/domain/funcionario"
	"gorm.io/gorm"
)

type FuncionarioRepository struct {
	db *gorm.DB
}

func NewFuncionarioRepository(db *gorm.DB) *FuncionarioRepository {
	return &FuncionarioRepository{
		db: db,
	}
}

// Create crea un nuevo funcionario
func (r *FuncionarioRepository) Create(ctx context.Context, f *funcionario.Funcionario) error {
	return r.db.WithContext(ctx).Create(f).Error
}

// GetByID obtiene un funcionario por ID
func (r *FuncionarioRepository) GetByID(ctx context.Context, id int) (*funcionario.Funcionario, error) {
	var f funcionario.Funcionario
	err := r.db.WithContext(ctx).Where("id_funcionario = ?", id).First(&f).Error
	if err != nil {
		return nil, err
	}
	return &f, nil
}

// GetAll obtiene todos los funcionarios
func (r *FuncionarioRepository) GetAll(ctx context.Context) ([]*funcionario.Funcionario, error) {
	var funcionarios []*funcionario.Funcionario
	err := r.db.WithContext(ctx).Find(&funcionarios).Error
	return funcionarios, err
}

// Update actualiza un funcionario existente
func (r *FuncionarioRepository) Update(ctx context.Context, f *funcionario.Funcionario) error {
	return r.db.WithContext(ctx).Save(f).Error
}

// Delete elimina un funcionario (soft delete)
func (r *FuncionarioRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&funcionario.Funcionario{}, id).Error
}

// GetByRut obtiene un funcionario por RUT
func (r *FuncionarioRepository) GetByRut(ctx context.Context, rut string) (*funcionario.Funcionario, error) {
	var f funcionario.Funcionario
	err := r.db.WithContext(ctx).Where("rut_funcionario = ?", rut).First(&f).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &f, nil
}

// GetByEmail obtiene un funcionario por email
func (r *FuncionarioRepository) GetByEmail(ctx context.Context, email string) (*funcionario.Funcionario, error) {
	var f funcionario.Funcionario
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&f).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &f, nil
}

// GetByFilter obtiene funcionarios según criterios de filtrado
func (r *FuncionarioRepository) GetByFilter(ctx context.Context, filter funcionario.FuncionarioFilter) ([]*funcionario.Funcionario, error) {
	var funcionarios []*funcionario.Funcionario
	query := r.db.WithContext(ctx)

	query = r.applyFilters(query, filter)

	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	err := query.Find(&funcionarios).Error
	return funcionarios, err
}

// Count cuenta funcionarios según criterios de filtrado
func (r *FuncionarioRepository) Count(ctx context.Context, filter funcionario.FuncionarioFilter) (int64, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&funcionario.Funcionario{})
	query = r.applyFilters(query, filter)
	err := query.Count(&count).Error
	return count, err
}

// applyFilters aplica los filtros a la query
func (r *FuncionarioRepository) applyFilters(query *gorm.DB, filter funcionario.FuncionarioFilter) *gorm.DB {
	if filter.RutFuncionario != "" {
		query = query.Where("rut_funcionario = ?", filter.RutFuncionario)
	}
	if filter.Email != "" {
		query = query.Where("email = ?", filter.Email)
	}
	if filter.IDEmpresaCliente != nil {
		query = query.Where("id_empresa_cliente = ?", *filter.IDEmpresaCliente)
	}
	if filter.IDSucursal != nil {
		query = query.Where("id_sucursal = ?", *filter.IDSucursal)
	}
	if filter.IDSegmento != nil {
		query = query.Where("id_segmento = ?", *filter.IDSegmento)
	}
	if filter.IDEstado != nil {
		query = query.Where("id_estado = ?", *filter.IDEstado)
	}
	if filter.IDCargo != nil {
		query = query.Where("id_cargo = ?", *filter.IDCargo)
	}
	if filter.TallasRegistradas != nil {
		query = query.Where("tallas_registradas = ?", *filter.TallasRegistradas)
	}
	return query
}

// GetByEmpresa obtiene funcionarios por empresa
func (r *FuncionarioRepository) GetByEmpresa(ctx context.Context, idEmpresa int) ([]*funcionario.Funcionario, error) {
	var funcionarios []*funcionario.Funcionario
	err := r.db.WithContext(ctx).Where("id_empresa_cliente = ?", idEmpresa).Find(&funcionarios).Error
	return funcionarios, err
}

// GetBySucursal obtiene funcionarios por sucursal
func (r *FuncionarioRepository) GetBySucursal(ctx context.Context, idSucursal int) ([]*funcionario.Funcionario, error) {
	var funcionarios []*funcionario.Funcionario
	err := r.db.WithContext(ctx).Where("id_sucursal = ?", idSucursal).Find(&funcionarios).Error
	return funcionarios, err
}

// GetBySegmento obtiene funcionarios por segmento
func (r *FuncionarioRepository) GetBySegmento(ctx context.Context, idSegmento int) ([]*funcionario.Funcionario, error) {
	var funcionarios []*funcionario.Funcionario
	err := r.db.WithContext(ctx).Where("id_segmento = ?", idSegmento).Find(&funcionarios).Error
	return funcionarios, err
}

// ActivateByID activa un funcionario
func (r *FuncionarioRepository) ActivateByID(ctx context.Context, id int) error {
	// Asumiendo que existe un id_estado donde 1 = activo
	return r.db.WithContext(ctx).Model(&funcionario.Funcionario{}).
		Where("id_funcionario = ?", id).
		Update("id_estado", 1).Error
}

// DeactivateByID desactiva un funcionario
func (r *FuncionarioRepository) DeactivateByID(ctx context.Context, id int) error {
	// Asumiendo que existe un id_estado donde 2 = inactivo
	return r.db.WithContext(ctx).Model(&funcionario.Funcionario{}).
		Where("id_funcionario = ?", id).
		Update("id_estado", 2).Error
}

// MedidasRepository implementa las operaciones de medidas
type MedidasRepository struct {
	db *gorm.DB
}

func NewMedidasRepository(db *gorm.DB) *MedidasRepository {
	return &MedidasRepository{
		db: db,
	}
}

// Create crea nuevas medidas
func (r *MedidasRepository) Create(ctx context.Context, medidas *funcionario.MedidasFuncionario) error {
	return r.db.WithContext(ctx).Create(medidas).Error
}

// GetByID obtiene medidas por ID
func (r *MedidasRepository) GetByID(ctx context.Context, id int) (*funcionario.MedidasFuncionario, error) {
	var medidas funcionario.MedidasFuncionario
	err := r.db.WithContext(ctx).Where("id_medidas = ?", id).First(&medidas).Error
	if err != nil {
		return nil, err
	}
	return &medidas, nil
}

// Update actualiza medidas existentes
func (r *MedidasRepository) Update(ctx context.Context, medidas *funcionario.MedidasFuncionario) error {
	return r.db.WithContext(ctx).Save(medidas).Error
}

// Delete elimina medidas
func (r *MedidasRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&funcionario.MedidasFuncionario{}, id).Error
}

// GetActivasByFuncionario obtiene las medidas activas (sin fecha fin) de un funcionario
func (r *MedidasRepository) GetActivasByFuncionario(ctx context.Context, idMedidas int) (*funcionario.MedidasFuncionario, error) {
	var medidas funcionario.MedidasFuncionario
	err := r.db.WithContext(ctx).
		Where("id_medidas = ? AND fecha_fin IS NULL", idMedidas).
		First(&medidas).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no se encontraron medidas activas")
		}
		return nil, err
	}
	return &medidas, nil
}

// GetHistorialByFuncionario obtiene el historial de medidas de un funcionario
func (r *MedidasRepository) GetHistorialByFuncionario(ctx context.Context, idMedidas int) ([]*funcionario.MedidasFuncionario, error) {
	var medidas []*funcionario.MedidasFuncionario
	err := r.db.WithContext(ctx).
		Where("id_medidas = ?", idMedidas).
		Order("fecha_inicio DESC").
		Find(&medidas).Error
	return medidas, err
}
