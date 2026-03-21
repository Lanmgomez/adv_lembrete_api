package lembretes

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

func (r *Repository) CreateLembreteInDB(lembrete *models.Lembrete) error {

	query := `
		INSERT INTO lembretes (
			entidade_id,
			nome_lembrete,
			descricao,
			status,
			data_vencimento,
			dias_antecedencia,
			email_notificacao
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.DB.Exec(
		query,
		lembrete.EntidadeID,
		lembrete.NomeLembrete,
		lembrete.Descricao,
		lembrete.Status,
		lembrete.DataVencimento,
		lembrete.DiasAntecedencia,
		lembrete.EmailNotificacao,
	)

	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	lembrete.ID = id

	return nil
}

func (r *Repository) GetAllLembretesInDB(nome string, limit, offset int) ([]models.Lembrete, int, error) {

	var total int

	// COUNT (com filtro)
	countQuery := `
		SELECT COUNT(*)
		FROM lembretes
		WHERE nome_lembrete LIKE ?
	`

	err := r.DB.QueryRow(countQuery, "%"+nome+"%").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// SELECT com paginação
	query := `
		SELECT 
			id,
			entidade_id,
			nome_lembrete,
			descricao,
			status,
			data_vencimento,
			dias_antecedencia,
			email_notificacao,
			created_at,
			updated_at
		FROM lembretes
		WHERE nome_lembrete LIKE ?
		ORDER BY id DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.DB.Query(query, "%"+nome+"%", limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	lista := make([]models.Lembrete, 0)

	for rows.Next() {
		var l models.Lembrete

		err := rows.Scan(
			&l.ID,
			&l.EntidadeID,
			&l.NomeLembrete,
			&l.Descricao,
			&l.Status,
			&l.DataVencimento,
			&l.DiasAntecedencia,
			&l.EmailNotificacao,
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

func (r *Repository) GetLembreteByIDInDB(id int64) (*models.Lembrete, error) {

	query := `
		SELECT 
			id, entidade_id, nome_lembrete, descricao, status,
			data_vencimento, dias_antecedencia, email_notificacao,
			created_at, updated_at
		FROM lembretes
		WHERE id = ?
	`

	var l models.Lembrete

	err := r.DB.QueryRow(query, id).Scan(
		&l.ID,
		&l.EntidadeID,
		&l.NomeLembrete,
		&l.Descricao,
		&l.Status,
		&l.DataVencimento,
		&l.DiasAntecedencia,
		&l.EmailNotificacao,
		&l.CreatedAt,
		&l.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (r *Repository) UpdateLembreteInDB(id int64, l *models.Lembrete) error {

	query := `
		UPDATE lembretes
		SET 
			nome_lembrete = ?,
			descricao = ?,
			status = ?,
			data_vencimento = ?,
			dias_antecedencia = ?,
			email_notificacao = ?
		WHERE id = ?
	`

	_, err := r.DB.Exec(
		query,
		l.NomeLembrete,
		l.Descricao,
		l.Status,
		l.DataVencimento,
		l.DiasAntecedencia,
		l.EmailNotificacao,
		id,
	)

	return err
}

func (r *Repository) DeleteLembreteInDB(id int64) error {
	query := `DELETE FROM lembretes WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

// UPDATE STATUS (para concluir depois)
func (r *Repository) UpdateLembreteStatusInDB(id int64, status string) error {
	query := `UPDATE lembretes SET status = ? WHERE id = ?`
	_, err := r.DB.Exec(query, status, id)
	return err
}