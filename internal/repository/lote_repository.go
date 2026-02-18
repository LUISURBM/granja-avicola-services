package repository

import (
	"granja-avicola/internal/models"
	"gorm.io/gorm"
)

type LoteRepository interface {
	Create(lote *models.Lote) error
	GetAll() ([]models.Lote, error)
	GetByID(id uint) (*models.Lote, error)
}

type loteRepo struct {
	db *gorm.DB
}

func NewLoteRepository(db *gorm.DB) LoteRepository {
	return &loteRepo{db}
}

func (r *loteRepo) Create(lote *models.Lote) error {
	return r.db.Create(lote).Error
}

func (r *loteRepo) GetAll() ([]models.Lote, error) {
	var lotes []models.Lote
	err := r.db.Preload("Mortalidades").Find(&lotes).Error
	return lotes, err
}

func (r *loteRepo) GetByID(id uint) (*models.Lote, error) {
	var lote models.Lote
	err := r.db.First(&lote, id).Error
	return &lote, err
}