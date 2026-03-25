package lembretes

import (
	"adv_lembrete_api/internal/models"
	"database/sql"
	"time"
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
			email_notificacao,
			next_send_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
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
		lembrete.NextSendAt,
	)

	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	lembrete.ID = id

	return nil
}

func (r *Repository) GetAllLembretesInDB(nome string, status string, limit, offset int) ([]models.Lembrete, int, error) {
	var total int

	countQuery := `
		SELECT COUNT(*)
		FROM lembretes
		WHERE nome_lembrete LIKE ?
		  AND (? = '' OR status = ?)
	`

	err := r.DB.QueryRow(countQuery, "%"+nome+"%", status, status).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

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
		  AND (? = '' OR status = ?)
		ORDER BY id DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.DB.Query(query, "%"+nome+"%", status, status, limit, offset)
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

func (r *Repository) UpdateLembreteStatusInDB(id int64, status string) error {
	query := `UPDATE lembretes SET status = ? WHERE id = ?`
	_, err := r.DB.Exec(query, status, id)
	return err
}

func (r *Repository) FindDueForSend(now time.Time) ([]models.Lembrete, error) {
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
			updated_at,
			last_sent_at,
			next_send_at
		FROM lembretes
		WHERE status IN ('pendente', 'atrasado')
		  AND next_send_at IS NOT NULL
		  AND next_send_at <= ?
		ORDER BY id ASC
	`

	rows, err := r.DB.Query(query, now)
	if err != nil {
		return nil, err
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
			&l.LastSentAt,
			&l.NextSendAt,
		)
		if err != nil {
			return nil, err
		}
		lista = append(lista, l)
	}

	return lista, nil
}

func (r *Repository) UpdateSendControl(id int64, status string, lastSentAt, nextSendAt time.Time) error {
	query := `
		UPDATE lembretes
		SET status = ?, last_sent_at = ?, next_send_at = ?
		WHERE id = ?
	`
	_, err := r.DB.Exec(query, status, lastSentAt, nextSendAt, id)
	return err
}

func (r *Repository) MarkAsConcluido(id int64) error {
	query := `
		UPDATE lembretes
		SET status = 'concluido',
		    next_send_at = NULL
		WHERE id = ?
	`

	_, err := r.DB.Exec(query, id)
	return err
}