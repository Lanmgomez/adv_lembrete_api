package lembretes

import (
	"adv_lembrete_api/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {

	var input models.CreateLembreteInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "dados inválidos"})
		return
	}

	_, err := h.service.CreateLembrete(input)
	if err != nil {

		if err.Error() == "entidade não encontrada" {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "lembrete criado com sucesso",
	})
}

func (h *Handler) GetAll(c *gin.Context) {

	nome := c.DefaultQuery("nome", "")
	status:= c.DefaultQuery("status", "")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	data, total, err := h.service.GetAllLembretes(nome, status, page, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": "erro ao buscar lembretes"})
		return
	}

	c.JSON(200, gin.H{
		"data":         data,
		"current_page": page,
		"total":        total,
	})
}

func (h *Handler) GetByID(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	data, err := h.service.GetLembreteByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "lembrete não encontrado"})
		return
	}

	c.JSON(200, gin.H{"data": data})
}

func (h *Handler) Update(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var input models.CreateLembreteInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "dados inválidos"})
		return
	}

	err := h.service.UpdateLembrete(id, input)
	if err != nil {
		c.JSON(500, gin.H{"error": "erro ao atualizar"})
		return
	}

	c.JSON(200, gin.H{
		"message": "lembrete atualizado com sucesso",
	})
}

func (h *Handler) Delete(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := h.service.DeleteLembrete(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "erro ao deletar"})
		return
	}

	c.JSON(200, gin.H{
		"message": "lembrete deletado com sucesso",
	})
}