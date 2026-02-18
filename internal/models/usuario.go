package models

import "time"

type Usuario struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Nombre        string    `gorm:"size:100;not null" json:"nombre"`
	Email         string    `gorm:"size:100;unique;not null" json:"email"`
	Password      string    `gorm:"size:255;not null" json:"-"` // Oculto en JSON
	Rol           string    `gorm:"type:enum('admin', 'operario');default:'operario'" json:"rol"`
	FechaRegistro time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"fecha_registro"`
}
