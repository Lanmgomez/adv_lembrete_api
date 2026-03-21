package entidades

import "adv_lembrete_api/internal/models"

type RepositoryInterface interface {
	CreateNewEntidadeInDB(lembrete *models.Entidade) error
	FindAllEntidadesPagineted(nome string, limit, offset int) ([]models.Entidade, int, error)
	FindByID(id int64) (*models.Entidade, error)
	UpdateEntidadeInDB(id int64, lembrete *models.Entidade) error
	DeleteEntidadeInDB(id int64) error
	ExistsByID(id int64) (bool, error)
	ExistsByNome(nome string) (bool, error)
}

type ServiceInterface interface {
	CreateNewEntidade(input models.CreateEntidadeInput) (*models.Entidade, error)
	GetAllEntidades(nome string, page, limit int) ([]models.Entidade, int, error)
	GetEntidadeByID(id int64) (*models.Entidade, error)
	UpdateEntidade(id int64, input models.CreateEntidadeInput) error
	DeleteEntidade(id int64) error
}