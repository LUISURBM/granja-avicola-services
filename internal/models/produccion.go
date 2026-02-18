package models

import "time"

// Tabla: mortalidad
type Mortalidad struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	LoteID   uint      `json:"lote_id"`
	Cantidad int       `gorm:"not null" json:"cantidad"`
	Causa    string    `gorm:"size:255" json:"causa"`
	Fecha    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"fecha"`
}

// Tabla: sanidad
type Sanidad struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	LoteID          uint      `json:"lote_id"`
	Tipo            string    `gorm:"type:enum('vacuna', 'medicamento', 'suplemento')" json:"tipo"`
	Descripcion     string    `gorm:"size:255" json:"descripcion"`
	FechaAplicacion time.Time `gorm:"type:date" json:"fecha_aplicacion"`
}

// Tabla: produccion_rendimiento
type Rendimiento struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	LoteID              uint      `json:"lote_id"`
	PesoPromedioGramos  float64   `gorm:"type:decimal(10,2)" json:"peso_promedio_gramos"`
	ConsumoAlimentoKg   float64   `gorm:"type:decimal(10,2)" json:"consumo_alimento_kg"`
	FechaRegistro       time.Time `gorm:"type:date" json:"fecha_registro"`
}
