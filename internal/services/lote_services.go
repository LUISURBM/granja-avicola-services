package services

import (
	"granja-avicola/internal/models"
	"granja-avicola/internal/repository"
)

type LoteService struct {
	repo repository.LoteRepository
}

func NewLoteService(repo repository.LoteRepository) *LoteService {
	return &LoteService{repo}
}

func (s *LoteService) CrearNuevoLote(lote *models.Lote) error {
	// Aquí podrías validar que el galpón tenga capacidad antes de guardar
	return s.repo.Create(lote)
}

func (s *LoteService) ObtenerTodos() ([]models.Lote, error) {
	return s.repo.GetAll()
}