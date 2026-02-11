package repository

import (
	"context"

	"github.com/Sistal/ms-funcionario/internal/domain/funcionario"
	"gorm.io/gorm"
)

type GeneroRepository struct {
	db *gorm.DB
}

func NewGeneroRepository(db *gorm.DB) *GeneroRepository {
	return &GeneroRepository{
		db: db,
	}
}

// FindAll obtiene todos los géneros
func (r *GeneroRepository) FindAll(ctx context.Context) ([]*funcionario.Genero, error) {
	var generos []*funcionario.Genero
	err := r.db.WithContext(ctx).Order("id_genero ASC").Find(&generos).Error
	return generos, err
}
