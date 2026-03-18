package auth

import (
	"adv_lembrete_api/internal/models"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) FindUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, username, email, password, user_type, created_at
		FROM users
		WHERE email = ?
	`

	var user models.User

	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.UserType,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}