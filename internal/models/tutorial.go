package models

import "time"

type Tutorial struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Titulo           string    `gorm:"size:150" json:"titulo"`
	Categoria        string    `gorm:"type:enum('crianza', 'mejora_produccion', 'bioseguridad')" json:"categoria"`
	Contenido        string    `gorm:"type:text" json:"contenido"`
	UrlVideo         string    `gorm:"size:255" json:"url_video"`
	FechaPublicacion time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"fecha_publicacion"`
}
