package funcionario

// Cargo representa el cargo o puesto de trabajo de un funcionario
type Cargo struct {
	IDCargo     int    `json:"id_cargo" gorm:"column:id_cargo;primaryKey;autoIncrement"`
	NombreCargo string `json:"nombre_cargo" gorm:"column:nombre_cargo;type:varchar(100)"`
}

func (Cargo) TableName() string {
	return "cargo"
}
