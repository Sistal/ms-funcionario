package repository

import (
	"context"

	"github.com/Sistal/ms-funcionario/internal/domain/funcionario"
	"gorm.io/gorm"
)

type CargoRepository struct {
	db *gorm.DB
}

func NewCargoRepository(db *gorm.DB) *CargoRepository {
	return &CargoRepository{
		db: db,
	}
}

// FindAll obtiene todos los cargos
func (r *CargoRepository) FindAll(ctx context.Context) ([]*funcionario.Cargo, error) {
	var cargos []*funcionario.Cargo
	err := r.db.WithContext(ctx).Order("nombre_cargo ASC").Find(&cargos).Error
	return cargos, err
}
