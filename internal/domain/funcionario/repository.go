package funcionario

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, funcionario *Funcionario) error
	GetByID(ctx context.Context, id uuid.UUID) (*Funcionario, error)
	GetAll(ctx context.Context) ([]*Funcionario, error)
	Update(ctx context.Context, funcionario *Funcionario) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByEmail(ctx context.Context, email string) (*Funcionario, error)
}
