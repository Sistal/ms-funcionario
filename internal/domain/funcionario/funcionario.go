package funcionario

import (
	"time"

	"github.com/google/uuid"
)

type Funcionario struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Nombre    string    `json:"nombre" gorm:"type:varchar(100);not null"`
	Apellido  string    `json:"apellido" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Cargo     string    `json:"cargo" gorm:"type:varchar(100)"`
	Activo    bool      `json:"activo" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Funcionario) TableName() string {
	return "funcionarios"
}
