package users

import (
	"adv_lembrete_api/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUsers(page, limit int) ([]models.User, int, error) {
	offset := (page - 1) * limit
	return s.repo.FindAllPaginated(limit, offset)
}

func (s *Service) GetUserByID(id int64) (*models.User, error) {
	return s.repo.FindUserByID(id)
}

func (s *Service) CreateUser(input models.CreateUserInput) (*models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: input.Username,
		Email: input.Email,
		Password: string(hashedPassword),
		UserType: input.UserType,
	}

	err = s.repo.CreateNewUserInDB(user)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (s *Service) UpdateUser(id int64, input models.CreateUserInput) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: input.Username,
		Email:    input.Email,
		UserType: input.UserType,
		Password: string(hashedPassword),
	}

	return s.repo.UpdateUserInDB(id, user)
}

func (s *Service) DeleteUser(id int64) error {
	return s.repo.DeleteUserInDB(id)
}