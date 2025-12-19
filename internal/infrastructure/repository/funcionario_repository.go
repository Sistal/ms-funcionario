package repository

import (
	"context"

	"github.com/Sistal/ms-funcionario/internal/domain/funcionario"
	"github.com/google/uuid"
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

func (r *FuncionarioRepository) Create(ctx context.Context, f *funcionario.Funcionario) error {
	return r.db.WithContext(ctx).Create(f).Error
}

func (r *FuncionarioRepository) GetByID(ctx context.Context, id uuid.UUID) (*funcionario.Funcionario, error) {
	var f funcionario.Funcionario
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&f).Error
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *FuncionarioRepository) GetAll(ctx context.Context) ([]*funcionario.Funcionario, error) {
	var funcionarios []*funcionario.Funcionario
	err := r.db.WithContext(ctx).Find(&funcionarios).Error
	return funcionarios, err
}

func (r *FuncionarioRepository) Update(ctx context.Context, f *funcionario.Funcionario) error {
	return r.db.WithContext(ctx).Save(f).Error
}

func (r *FuncionarioRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&funcionario.Funcionario{}, id).Error
}

func (r *FuncionarioRepository) GetByEmail(ctx context.Context, email string) (*funcionario.Funcionario, error) {
	var f funcionario.Funcionario
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&f).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &f, nil
}
