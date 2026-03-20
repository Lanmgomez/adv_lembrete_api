package entidades

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

func (h *Handler) CreateEntidade(c *gin.Context) {
	var input models.CreateEntidadeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "dados inválidos"})
		return
	}

	_, err := h.service.CreateNewEntidade(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "entidade criada com sucesso",
	})
}

func (h *Handler) GetAllEntidades(c *gin.Context) {

	nome_entidade := c.DefaultQuery("nome_entidade", "")
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

	data, total, err := h.service.GetAllEntidades(nome_entidade, page, limit)
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

func (h *Handler) GetEntidadeByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	data, err := h.service.GetEntidadeByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "lembrete não encontrado"})
		return
	}

	c.JSON(200, gin.H{"data": data})
}

func (h *Handler) UpdateEntidade(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var input models.CreateEntidadeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "dados inválidos"})
		return
	}

	err := h.service.UpdateEntidade(id, input)
	if err != nil {
		c.JSON(500, gin.H{"error": "erro ao atualizar"})
		return
	}

	c.JSON(200, gin.H{"message": "atualizado com sucesso"})
}

func (h *Handler) DeleteEntidade(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := h.service.DeleteEntidade(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "erro ao deletar"})
		return
	}

	c.JSON(200, gin.H{"message": "deletado com sucesso"})
}