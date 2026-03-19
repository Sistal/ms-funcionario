package funcionario

import (
	"context"
)

// FuncionarioFilter representa los criterios de filtrado para búsqueda de funcionarios
type FuncionarioFilter struct {
	RutFuncionario    string
	Email             string
	IDEmpresaCliente  *int
	IDSucursal        *int
	IDSegmento        *int
	IDEstado          *int
	IDCargo           *int
	IDGenero          *int
	TallasRegistradas *bool
	Search            string // Búsqueda por texto en nombres, apellidos, RUT, email
	SortBy            string // Campo de ordenamiento
	Order             string // Orden: asc o desc
	Limit             int
	Offset            int
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
	GetByUserID(ctx context.Context, userID int) (*Funcionario, error)
	GetByFilter(ctx context.Context, filter FuncionarioFilter) ([]*Funcionario, error)
	Count(ctx context.Context, filter FuncionarioFilter) (int64, error)

	// Operaciones relacionadas con empresa y sucursal
	GetByEmpresa(ctx context.Context, idEmpresa int) ([]*Funcionario, error)
	GetBySucursal(ctx context.Context, idSucursal int) ([]*Funcionario, error)
	GetBySegmento(ctx context.Context, idSegmento int) ([]*Funcionario, error)

	// Activación/Desactivación
	ActivateByID(ctx context.Context, id int) error
	DeactivateByID(ctx context.Context, id int) error

	// Actualización de medidas
	UpdateMedidasInfo(ctx context.Context, idFuncionario int, idMedidas int) error
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

// SucursalRepository define las operaciones de persistencia para Sucursal
type SucursalRepository interface {
	FindAll(ctx context.Context) ([]*Sucursal, error)
	GetByID(ctx context.Context, id int) (*Sucursal, error)
}

// CargoRepository define las operaciones de persistencia para Cargo
type CargoRepository interface {
	FindAll(ctx context.Context) ([]*Cargo, error)
}

// GeneroRepository define las operaciones de persistencia para Genero
type GeneroRepository interface {
	FindAll(ctx context.Context) ([]*Genero, error)
}
