package users

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

func (r *Repository) FindAllPaginated(limit, offset int) ([]models.User, int, error) {
	var total int

	err := r.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, username, email, user_type, created_at
		FROM users
		ORDER BY id DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.UserType,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		users = append(users, user)
	}

	return users, total, nil
}

func (r *Repository) FindUserByID(id int64) (*models.User, error) {
	query := `
		SELECT id, username, email, user_type, created_at
		FROM users
		WHERE id = ?
	`

	var user models.User

	err := r.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.UserType,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateNewUserInDB(user *models.User) error {
	query := `
		INSERT INTO users (username, email, password, user_type)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.DB.Exec(query, user.Username, user.Email, user.Password, user.UserType)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	user.ID = id

	return nil
}

func (r *Repository) UpdateUserInDB(id int64, user *models.User) error {
	query := `
		UPDATE users
		SET username = ?, email = ?, password = ?, user_type = ?
		WHERE id = ?
	`

	_, err := r.DB.Exec(query, user.Username, user.Email, user.Password, user.UserType, id)
	return err
}

func (r *Repository) DeleteUserInDB(id int64) error {
	query := `DELETE FROM users WHERE id = ?`

	_, err := r.DB.Exec(query, id)
	return err
}