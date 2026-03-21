package entidades

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

func (r *Repository) CreateNewEntidadeInDB(entidade *models.Entidade) error {
	query := `
		INSERT INTO entidades (nome_entidade)
		VALUES (?)
	`

	result, err := r.DB.Exec(query, entidade.NomeEntidade)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	entidade.ID = id

	return nil
}

func (r *Repository) ExistsByID(id int64) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM entidades WHERE id = ?)`

	err := r.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *Repository) ExistsByNome(nome string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM entidades WHERE nome_entidade = ?)`
	err := r.DB.QueryRow(query, nome).Scan(&exists)

	return exists, err
}

func (r *Repository) FindAllEntidadesPagineted(nome string, limit, offset int) ([]models.Entidade, int, error) {

	var total int

	// COUNT com filtro
	countQuery := `SELECT COUNT(*) FROM entidades WHERE nome_entidade LIKE ?`
	err := r.DB.QueryRow(countQuery, "%"+nome+"%").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// SELECT com filtro + paginação
	query := `
		SELECT id, nome_entidade, created_at, updated_at
		FROM entidades
		WHERE nome_entidade LIKE ?
		ORDER BY id DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.DB.Query(query, "%"+nome+"%", limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var lista []models.Entidade

	for rows.Next() {
		var l models.Entidade

		err := rows.Scan(
			&l.ID,
			&l.NomeEntidade,
			&l.CreatedAt,
			&l.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		lista = append(lista, l)
	}

	return lista, total, nil
}

func (r *Repository) FindByID(id int64) (*models.Entidade, error) {
	query := `
		SELECT id, nome_entidade, created_at, updated_at
		FROM entidades
		WHERE id = ?
	`

	var l models.Entidade

	err := r.DB.QueryRow(query, id).Scan(
		&l.ID,
		&l.NomeEntidade,
		&l.CreatedAt,
		&l.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (r *Repository) UpdateEntidadeInDB(id int64, lembrete *models.Entidade) error {
	query := `
		UPDATE entidades
		SET nome_entidade = ?
		WHERE id = ?
	`

	_, err := r.DB.Exec(query, lembrete.NomeEntidade, id)
	return err
}

func (r *Repository) DeleteEntidadeInDB(id int64) error {
	query := `DELETE FROM entidades WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}
