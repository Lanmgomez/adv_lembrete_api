package lembretes

import (
	"adv_lembrete_api/internal/domain/entidades"
	"adv_lembrete_api/internal/models"
	"errors"
	"fmt"
	"time"
)

type Service struct {
	repo *Repository
	entidadesRepo  entidades.RepositoryInterface
}

func NewService(repo *Repository, entidadesRepo *entidades.Repository) *Service {
	return &Service{
		repo:          repo,
		entidadesRepo: entidadesRepo,
	}
}

func calcularStatusEDias(l *models.Lembrete) {

	hoje := time.Now().Truncate(24 * time.Hour)
	vencimento := l.DataVencimento

	diff := int(vencimento.Sub(hoje).Hours() / 24)

	if l.Status != "concluido" {

		if diff < 0 {
			l.Status = "atrasado"
			l.DiasRestantes = fmt.Sprintf("%d dias em atraso", -diff)
		} else {
			l.DiasRestantes = fmt.Sprintf("%d dias restantes", diff)
		}
	}
}

func (s *Service) CreateLembrete(input models.CreateLembreteInput) (*models.Lembrete, error) {

	exists, err := s.entidadesRepo.ExistsByID(input.EntidadeID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("entidade não encontrada")
	}

	// converter data
	dataVencimento, err := time.Parse("2006-01-02", input.DataVencimento)
	if err != nil {
		return nil, errors.New("formato de data inválido (use YYYY-MM-DD)")
	}

	lembrete := &models.Lembrete{
		EntidadeID:       input.EntidadeID,
		NomeLembrete:     input.NomeLembrete,
		Descricao:        input.Descricao,
		Status:           "pendente", // 🔥 sempre controlado pelo sistema
		// DataVencimento:   input.DataVencimento,
		DataVencimento:   dataVencimento,
		DiasAntecedencia: input.DiasAntecedencia,
		EmailNotificacao: input.EmailNotificacao,
	}

	err = s.repo.CreateLembreteInDB(lembrete)
	if err != nil {
		return nil, err
	}

	return lembrete, nil
}

func (s *Service) GetAllLembretes(nome string, page int, limit int) ([]models.Lembrete, int, error) {
	offset := (page - 1) * limit

	list, total, err := s.repo.GetAllLembretesInDB(nome, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	for i := range list {
		calcularStatusEDias(&list[i])
	}

	return list, total, nil
}

func (s *Service) GetLembreteByID(id int64) (*models.Lembrete, error) {

	l, err := s.repo.GetLembreteByIDInDB(id)
	if err != nil {
		return nil, err
	}

	calcularStatusEDias(l)

	return l, nil
}

func (s *Service) UpdateLembrete(id int64, input models.CreateLembreteInput) error {

	dataVencimento, err := time.Parse("2006-01-02", input.DataVencimento)
	if err != nil {
		return errors.New("data inválida")
	}

	l := &models.Lembrete{
		NomeLembrete:     input.NomeLembrete,
		Descricao:        input.Descricao,
		Status:           input.Status,
		DataVencimento:   dataVencimento,
		DiasAntecedencia: input.DiasAntecedencia,
		EmailNotificacao: input.EmailNotificacao,
	}

	return s.repo.UpdateLembreteInDB(id, l)
}

func (s *Service) DeleteLembrete(id int64) error {
	return s.repo.DeleteLembreteInDB(id)
}