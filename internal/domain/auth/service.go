package auth

import (
	"adv_lembrete_api/internal/models"
	"adv_lembrete_api/internal/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
			return nil, errors.New("usuário ou senha inválidos")
		}
		return nil, err
	}

	// valida senha
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("usuário ou senha inválidos")
	}

	// gerar token
	expiresAt := time.Now().Add(utils.GetJWTExpiresIn())

	claims := utils.CustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email: user.Email,
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

	user.Password = ""

	return &models.LoginResponse{
		AccessToken: signedToken,
		TokenType:   "Bearer",
		User:        *user,
	}, nil
}