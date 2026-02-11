package funcionario

// Genero representa el género de un funcionario
type Genero struct {
	IDGenero     int    `json:"id_genero" gorm:"column:id_genero;primaryKey;autoIncrement"`
	NombreGenero string `json:"nombre_genero" gorm:"column:nombre_genero;type:varchar(50)"`
}

func (Genero) TableName() string {
	return "Genero"
}
