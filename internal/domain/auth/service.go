package auth

import (
	"adv_lembrete_api/internal/models"
	"adv_lembrete_api/internal/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Login(input models.LoginInput) (*models.LoginResponse, error) {

    user, err := s.repo.FindUserByEmail(input.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("usuário não encontrado no banco")
        }
        return nil, errors.New("erro interno ao buscar usuário: " + err.Error())
    }

    if user.Password != input.Password {
        return nil, errors.New("usuário ou senha inválidos")
    }

    // gerar token
    expiresAt := time.Now().Add(utils.GetJWTExpiresIn())

    claims := utils.CustomClaims{
        UserID:   user.ID,
        Username: user.Username,
        Email:    user.Email,
        UserType: user.UserType,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expiresAt),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Subject:   user.Username,
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    signedToken, err := token.SignedString([]byte(utils.GetJWTSecret()))
    if err != nil {
        return nil, err
    }

    // Limpa a senha antes de retornar o objeto para o cliente
    user.Password = ""

    return &models.LoginResponse{
        AccessToken: signedToken,
        TokenType:   "Bearer",
        User:        *user,
    }, nil
}