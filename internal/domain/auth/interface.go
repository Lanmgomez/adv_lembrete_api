package auth

import "adv_lembrete_api/internal/models"

type ServiceInterface interface {
	Login(input models.LoginInput) (*models.LoginResponse, error)
}