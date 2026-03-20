package entidades

import (
	"adv_lembrete_api/internal/models"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateNewEntidade(input models.CreateEntidadeInput) (*models.Entidade, error) {
	exists, err := s.repo.ExistsByNome(input.NomeEntidade)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("entidade já cadastrada")
	}
	
	entidade := &models.Entidade{
		NomeEntidade: input.NomeEntidade,
	}

	err = s.repo.CreateNewEntidadeInDB(entidade)
	if err != nil {
		return nil, err
	}

	return entidade, nil
}

func (s *Service) GetAllEntidades(nome string, page int, limit int) ([]models.Entidade, int, error) {
	offset := (page - 1) * limit
	return s.repo.FindAllEntidadesPagineted(nome, limit, offset)
}

func (s *Service) GetEntidadeByID(id int64) (*models.Entidade, error) {
	return s.repo.FindByID(id)
}

func (s *Service) UpdateEntidade(id int64, input models.CreateEntidadeInput) error {
	entidade := &models.Entidade{
		NomeEntidade: input.NomeEntidade,
	}

	return s.repo.UpdateEntidadeInDB(id, entidade)
}

func (s *Service) DeleteEntidade(id int64) error {
	return s.repo.DeleteEntidadeInDB(id)
}