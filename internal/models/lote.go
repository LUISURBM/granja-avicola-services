package models

type Galpon struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	Nombre       string  `gorm:"size:50;not null" json:"nombre"`
	CapacidadMax int     `gorm:"not null" json:"capacidad_max"`
	Ubicacion    string  `gorm:"size:100" json:"ubicacion"`
	Lotes        []Lote  `gorm:"foreignKey:GalponID" json:"lotes,omitempty"`
}
