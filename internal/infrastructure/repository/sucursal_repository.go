package repository

import (
	"context"

	"github.com/Sistal/ms-funcionario/internal/domain/funcionario"
	"gorm.io/gorm"
)

type SucursalRepository struct {
	db *gorm.DB
}

func NewSucursalRepository(db *gorm.DB) *SucursalRepository {
	return &SucursalRepository{
		db: db,
	}
}

// FindAll devuelve todas las sucursales
func (r *SucursalRepository) FindAll(ctx context.Context) ([]*funcionario.Sucursal, error) {
	var sucursales []*funcionario.Sucursal
	err := r.db.WithContext(ctx).Find(&sucursales).Error
	return sucursales, err
}

// GetByID busca una sucursal por ID
func (r *SucursalRepository) GetByID(ctx context.Context, id int) (*funcionario.Sucursal, error) {
	var s funcionario.Sucursal
	err := r.db.WithContext(ctx).Where("id_sucursal = ?", id).First(&s).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}
