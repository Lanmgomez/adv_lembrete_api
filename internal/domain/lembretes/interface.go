package lembretes

import "adv_lembrete_api/internal/models"

type RepositoryInterface interface {
	Create(lembrete *models.Lembrete) error
}

type ServiceInterface interface {
	Create(input models.CreateLembreteInput) (*models.Lembrete, error)
}