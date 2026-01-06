package funcionario

import (
	"time"
)

// Funcionario representa la entidad principal de un empleado
type Funcionario struct {
	IDFuncionario      int        `json:"id_funcionario" gorm:"column:id_funcionario;primaryKey;autoIncrement"`
	RutFuncionario     string     `json:"rut_funcionario" gorm:"column:rut_funcionario;type:varchar(20)"`
	Nombres            string     `json:"nombres" gorm:"column:nombres;type:varchar(100)"`
	ApellidoPaterno    string     `json:"apellido_paterno" gorm:"column:apellido_paterno;type:varchar(100)"`
	ApellidoMaterno    string     `json:"apellido_materno" gorm:"column:apellido_materno;type:varchar(100)"`
	Celular            string     `json:"celular" gorm:"column:celular;type:varchar(20)"`
	Telefono           string     `json:"telefono" gorm:"column:telefono;type:varchar(20)"`
	Email              string     `json:"email" gorm:"column:email;type:varchar(100)"`
	TallasRegistradas  bool       `json:"tallas_registradas" gorm:"column:tallas_registradas;default:false"`
	Direccion          string     `json:"direccion" gorm:"column:direccion;type:varchar(255)"`
	FechaCreacion      *time.Time `json:"fecha_creacion" gorm:"column:fecha_creación;type:date;autoCreateTime"`
	FechaModificacion  *time.Time `json:"fecha_modificacion" gorm:"column:fecha_modificación;type:date;autoUpdateTime"`
	IDGenero           *int       `json:"id_genero" gorm:"column:id_genero"`
	IDMedidas          *int       `json:"id_medidas" gorm:"column:id_medidas"`
	IDUsuario          *int       `json:"id_usuario" gorm:"column:id_usuario"`
	IDEstado           *int       `json:"id_estado" gorm:"column:id_estado"`
	IDSucursal         *int       `json:"id_sucursal" gorm:"column:id_sucursal"`
	IDEmpresaCliente   *int       `json:"id_empresa_cliente" gorm:"column:id_empresa_cliente"`
	IDSegmento         *int       `json:"id_segmento" gorm:"column:id_segmento"`
	IDCargo            *int       `json:"id_cargo" gorm:"column:id_cargo"`
}

// MedidasFuncionario representa las medidas corporales del funcionario
type MedidasFuncionario struct {
	IDMedidas    int        `json:"id_medidas" gorm:"column:id_medidas;primaryKey;autoIncrement"`
	EstaturaM    *float64   `json:"estatura_m" gorm:"column:estatura_m;type:numeric(5,2)"`
	PechoCm      *float64   `json:"pecho_cm" gorm:"column:pecho_cm;type:numeric(5,2)"`
	CinturaCm    *float64   `json:"cintura_cm" gorm:"column:cintura_cm;type:numeric(5,2)"`
	CaderaCm     *float64   `json:"cadera_cm" gorm:"column:cadera_cm;type:numeric(5,2)"`
	MangaCm      *float64   `json:"manga_cm" gorm:"column:manga_cm;type:numeric(5,2)"`
	FechaInicio  *time.Time `json:"fecha_inicio" gorm:"column:fecha_inicio;type:date"`
	FechaFin     *time.Time `json:"fecha_fin" gorm:"column:fecha_fin;type:date"`
}

func (Funcionario) TableName() string {
	return "Funcionario"
}

func (MedidasFuncionario) TableName() string {
	return "Medidas Funcionario"
}
