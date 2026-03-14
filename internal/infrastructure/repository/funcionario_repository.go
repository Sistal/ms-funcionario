package repository

import (
	"context"
	"database/sql"
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

	// ESTRATEGIA FINAL: Bypass total de GORM para evitar prepared statements cache
	sqlDB, err := r.db.DB()
	if err != nil {
		return nil, err
	}

	// 1. Cargar Funcionario Base (bypass GORM)
	// Nota: Las columnas fecha_creación y fecha_modificación llevan tilde en la BD
	query := fmt.Sprintf(`
		SELECT id_funcionario, rut_funcionario, nombres, apellido_paterno, apellido_materno, 
		       celular, telefono, email, tallas_registradas, direccion, fecha_creación, fecha_modificación,
		       id_genero, id_medidas, id_usuario, id_estado, id_sucursal, id_empresa_cliente, id_segmento, id_cargo
		FROM "Funcionario" WHERE id_funcionario = %d`, id)

	err = sqlDB.QueryRowContext(ctx, query).Scan(
		&f.IDFuncionario, &f.RutFuncionario, &f.Nombres, &f.ApellidoPaterno, &f.ApellidoMaterno,
		&f.Celular, &f.Telefono, &f.Email, &f.TallasRegistradas, &f.Direccion, &f.FechaCreacion, &f.FechaModificacion,
		&f.IDGenero, &f.IDMedidas, &f.IDUsuario, &f.IDEstado, &f.IDSucursal, &f.IDEmpresaCliente, &f.IDSegmento, &f.IDCargo,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	// 2. Cargar Relaciones (bypass GORM)

	// Cargo
	if f.IDCargo != nil {
		var cargo funcionario.Cargo
		q := fmt.Sprintf(`SELECT id_cargo, nombre_cargo FROM "cargo" WHERE id_cargo = %d`, *f.IDCargo)
		err = sqlDB.QueryRowContext(ctx, q).Scan(&cargo.IDCargo, &cargo.NombreCargo)
		if err == nil {
			f.Cargo = &cargo
		}
	}

	// Sucursal
	if f.IDSucursal != nil {
		var sucursal funcionario.Sucursal
		q := fmt.Sprintf(`SELECT id_sucursal, nombre_sucursal, direccion, estado_sucursal FROM "Sucursal" WHERE id_sucursal = %d`, *f.IDSucursal)
		err = sqlDB.QueryRowContext(ctx, q).Scan(&sucursal.IDSucursal, &sucursal.NombreSucursal, &sucursal.Direccion, &sucursal.EstadoSucursal)
		if err == nil {
			f.Sucursal = &sucursal
		}
	}

	// Medidas
	if f.IDMedidas != nil {
		var medidas funcionario.MedidasFuncionario
		q := fmt.Sprintf(`SELECT id_medidas, estatura_m, pecho_cm, cintura_cm, cadera_cm, manga_cm, fecha_inicio, fecha_fin FROM "Medidas Funcionario" WHERE id_medidas = %d`, *f.IDMedidas)
		err = sqlDB.QueryRowContext(ctx, q).Scan(&medidas.IDMedidas, &medidas.EstaturaM, &medidas.PechoCm, &medidas.CinturaCm, &medidas.CaderaCm, &medidas.MangaCm, &medidas.FechaInicio, &medidas.FechaFin)
		if err == nil {
			f.Medidas = &medidas
		}
	}

	// Genero
	if f.IDGenero != nil {
		var genero funcionario.Genero
		q := fmt.Sprintf(`SELECT id_genero, nombre_genero FROM "Genero" WHERE id_genero = %d`, *f.IDGenero)
		err = sqlDB.QueryRowContext(ctx, q).Scan(&genero.IDGenero, &genero.NombreGenero)
		if err == nil {
			f.Genero = &genero
		}
	}

	// Estado
	if f.IDEstado != nil {
		var estado funcionario.Estado
		q := fmt.Sprintf(`SELECT id_estado, nombre_estado, tabla_estado FROM "Estado" WHERE id_estado = %d`, *f.IDEstado)
		err = sqlDB.QueryRowContext(ctx, q).Scan(&estado.IDEstado, &estado.NombreEstado, &estado.TablaEstado)
		if err == nil {
			f.Estado = &estado
		}
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

// GetByUserID obtiene un funcionario por ID de usuario
func (r *FuncionarioRepository) GetByUserID(ctx context.Context, userID int) (*funcionario.Funcionario, error) {
	var f funcionario.Funcionario

	// ESTRATEGIA FINAL DE EMERGENCIA: Bypass total de GORM
	// Usamos directamente sql.DB para ejecutar la query sin ningún middleware de GORM
	sqlDB, err := r.db.DB()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`SELECT id_funcionario FROM "Funcionario" WHERE id_usuario = %d`, userID)

	// Ejecutamos QueryRow directa
	// Solo necesitamos el ID para este endpoint, escaneamos solo eso para minimizar errores de mapeo
	var idFuncionario int
	err = sqlDB.QueryRowContext(ctx, query).Scan(&idFuncionario)

	if err != nil {
		// Si no hay filas, sql.ErrNoRows
		if err.Error() == "sql: no rows in result set" {
			return nil, nil // No encontrado
		}
		return nil, err
	}

	f.IDFuncionario = idFuncionario
	return &f, nil
}

// GetByFilter obtiene funcionarios según criterios de filtrado
func (r *FuncionarioRepository) GetByFilter(ctx context.Context, filter funcionario.FuncionarioFilter) ([]*funcionario.Funcionario, error) {
	var funcionarios []*funcionario.Funcionario
	query := r.db.WithContext(ctx).
		Preload("Cargo").
		Preload("Sucursal").
		Preload("Genero").
		Preload("Estado")

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
	if filter.IDGenero != nil {
		query = query.Where("id_genero = ?", *filter.IDGenero)
	}
	if filter.TallasRegistradas != nil {
		query = query.Where("tallas_registradas = ?", *filter.TallasRegistradas)
	}
	// Búsqueda por texto en múltiples campos
	if filter.Search != "" {
		searchPattern := "%" + filter.Search + "%"
		query = query.Where(
			"nombres ILIKE ? OR apellido_paterno ILIKE ? OR apellido_materno ILIKE ? OR rut_funcionario ILIKE ? OR email ILIKE ?",
			searchPattern, searchPattern, searchPattern, searchPattern, searchPattern,
		)
	}
	// Ordenamiento
	if filter.SortBy != "" {
		order := "DESC" // default
		if filter.Order != "" {
			if filter.Order == "asc" || filter.Order == "ASC" {
				order = "ASC"
			}
		}
		query = query.Order(filter.SortBy + " " + order)
	} else {
		// Ordenamiento por defecto
		query = query.Order("fecha_creación DESC")
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
	// Estado 10 = Activo según contrato
	return r.db.WithContext(ctx).Model(&funcionario.Funcionario{}).
		Where("id_funcionario = ?", id).
		Update("id_estado", 10).Error
}

// DeactivateByID desactiva un funcionario
func (r *FuncionarioRepository) DeactivateByID(ctx context.Context, id int) error {
	// Estado 11 = Inactivo (asumido, ajustar según tabla Estado real)
	return r.db.WithContext(ctx).Model(&funcionario.Funcionario{}).
		Where("id_funcionario = ?", id).
		Update("id_estado", 11).Error
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
