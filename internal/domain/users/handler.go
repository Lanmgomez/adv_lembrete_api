package users

import (
	"adv_lembrete_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetUsers(c *gin.Context) {
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

	users, total, err := h.service.GetUsers(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao buscar usuários",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":         users,
		"current_page": page,
		"total":        total,
	})
}

func (h *Handler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *Handler) CreateUser(c *gin.Context) {
	var input models.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "dados inválidos",
		})
		return
	}

	_, err := h.service.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao criar usuário",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	var input models.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "dados inválidos",
		})
		return
	}

	err := h.service.UpdateUser(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao atualizar usuário",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "usuário atualizado com sucesso",
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	err := h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao deletar usuário",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "usuário deletado com sucesso",
	})
}
