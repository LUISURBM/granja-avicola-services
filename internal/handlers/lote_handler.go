package handlers

import (
	"granja-avicola/internal/models"
	"granja-avicola/internal/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type LoteHandler struct {
	service *services.LoteService
}

func NewLoteHandler(s *services.LoteService) *LoteHandler {
	return &LoteHandler{service: s}
}

func (h *LoteHandler) GetAll(c *gin.Context) {
	lotes, err := h.service.ObtenerTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener lotes"})
		return
	}
	c.JSON(http.StatusOK, lotes)
}

func (h *LoteHandler) Create(c *gin.Context) {
	var lote models.Lote
	if err := c.ShouldBindJSON(&lote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CrearNuevoLote(&lote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear lote"})
		return
	}
	c.JSON(http.StatusCreated, lote)
}