package funcionario

import (
	"context"
)

// FuncionarioFilter representa los criterios de filtrado para búsqueda de funcionarios
type FuncionarioFilter struct {
	RutFuncionario   string
	Email            string
	IDEmpresaCliente *int
	IDSucursal       *int
	IDSegmento       *int
	IDEstado         *int
	IDCargo          *int
	TallasRegistradas *bool
	Limit            int
	Offset           int
}

// Repository define las operaciones de persistencia para Funcionario
type Repository interface {
	// Operaciones CRUD básicas
	Create(ctx context.Context, funcionario *Funcionario) error
	GetByID(ctx context.Context, id int) (*Funcionario, error)
	GetAll(ctx context.Context) ([]*Funcionario, error)
	Update(ctx context.Context, funcionario *Funcionario) error
	Delete(ctx context.Context, id int) error
	
	// Búsquedas específicas
	GetByRut(ctx context.Context, rut string) (*Funcionario, error)
	GetByEmail(ctx context.Context, email string) (*Funcionario, error)
	GetByFilter(ctx context.Context, filter FuncionarioFilter) ([]*Funcionario, error)
	Count(ctx context.Context, filter FuncionarioFilter) (int64, error)
	
	// Operaciones relacionadas con empresa y sucursal
	GetByEmpresa(ctx context.Context, idEmpresa int) ([]*Funcionario, error)
	GetBySucursal(ctx context.Context, idSucursal int) ([]*Funcionario, error)
	GetBySegmento(ctx context.Context, idSegmento int) ([]*Funcionario, error)
	
	// Activación/Desactivación
	ActivateByID(ctx context.Context, id int) error
	DeactivateByID(ctx context.Context, id int) error
}

// MedidasRepository define las operaciones de persistencia para MedidasFuncionario
type MedidasRepository interface {
	Create(ctx context.Context, medidas *MedidasFuncionario) error
	GetByID(ctx context.Context, id int) (*MedidasFuncionario, error)
	Update(ctx context.Context, medidas *MedidasFuncionario) error
	Delete(ctx context.Context, id int) error
	GetActivasByFuncionario(ctx context.Context, idMedidas int) (*MedidasFuncionario, error)
	GetHistorialByFuncionario(ctx context.Context, idMedidas int) ([]*MedidasFuncionario, error)
}
