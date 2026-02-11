package funcionario

// Estado representa el estado de una entidad (Funcionario, Sucursal, etc.)
type Estado struct {
	IDEstado     int    `json:"id_estado" gorm:"column:id_estado;primaryKey;autoIncrement"`
	NombreEstado string `json:"nombre_estado" gorm:"column:nombre_estado;type:varchar(100)"`
	TablaEstado  string `json:"tabla_estado" gorm:"column:tabla_estado;type:varchar(50)"`
}

func (Estado) TableName() string {
	return "Estado"
}
