package models

import "time"

type Lote struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	GalponID        uint      `json:"galpon_id"`
	Raza            string    `gorm:"size:50" json:"raza"`
	CantidadInicial int       `gorm:"not null" json:"cantidad_inicial"`
	FechaLlegada    time.Time `gorm:"type:date;not null" json:"fecha_llegada"`
	Estado          string    `gorm:"type:enum('activo', 'finalizado');default:'activo'" json:"estado"`
	// Relaciones para reportes completos
	Mortalidades    []Mortalidad `gorm:"foreignKey:LoteID" json:"mortalidades,omitempty"`
	Sanidad         []Sanidad    `gorm:"foreignKey:LoteID" json:"sanidad,omitempty"`
}