package funcionario

// Sucursal representa una oficina o sucursal de la empresa
type Sucursal struct {
	IDSucursal     int    `json:"id_sucursal" gorm:"column:id_sucursal;primaryKey;autoIncrement"`
	NombreSucursal string `json:"nombre_sucursal" gorm:"column:nombre_sucursal;type:varchar(100)"`
	Direccion      string `json:"direccion" gorm:"column:direccion;type:varchar(255)"`
	EstadoSucursal int    `json:"estado_sucursal" gorm:"column:estado_sucursal"`
}

func (Sucursal) TableName() string {
	return "Sucursal"
}
