package models

import "time"

type Lembrete struct {
	ID                 int64      `json:"id"`
	EntidadeID         int64      `json:"entidade_id"`
	NomeLembrete       string     `json:"nome_lembrete"`
	Descricao          string     `json:"descricao"`
	Status             string     `json:"status"`
	DataVencimento     time.Time  `json:"data_vencimento"`
	DiasAntecedencia   int        `json:"dias_antecedencia"`
	EmailNotificacao   string     `json:"email_notificacao"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	LastSentAt         *time.Time `json:"last_sent_at"`
	NextSendAt         *time.Time `json:"next_send_at"`
	DiasRestantes      string     `json:"dias_restantes"`
}

type CreateLembreteInput struct {
	EntidadeID       int64  `json:"entidade_id" binding:"required"`
	NomeLembrete     string `json:"nome_lembrete" binding:"required"`
	Descricao        string `json:"descricao"`
	Status           string    `json:"status"`
	DataVencimento   string `json:"data_vencimento" binding:"required"` // string → parse depois
	DiasAntecedencia int    `json:"dias_antecedencia"`
	EmailNotificacao string `json:"email_notificacao"`
}