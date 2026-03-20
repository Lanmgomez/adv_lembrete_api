package models

import "time"

type Entidade struct {
	ID            int64     `json:"id"`
	NomeEntidade  string    `json:"nome_entidade"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateEntidadeInput struct {
	NomeEntidade string `json:"nome_entidade" binding:"required"`
}