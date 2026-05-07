package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"adv_lembrete_api/internal/domain/lembretes"
)

const (
	reminderSendHour = 8

	statusPendente = "pendente"
	statusAtrasado = "atrasado"
)

func ProcessDueReminders(ctx context.Context, service *lembretes.Service, now time.Time) error {
	items, err := service.FindDueForSend(now)
	if err != nil {
		return fmt.Errorf("erro ao buscar lembretes para envio: %w", err)
	}

	today := truncateDate(now)

	for _, l := range items {
		if err := ctx.Err(); err != nil {
			return err
		}

		if l.Status != statusPendente && l.Status != statusAtrasado {
			continue
		}

		status := l.Status
		due := truncateDate(l.DataVencimento.In(now.Location()))

		if today.After(due) {
			status = statusAtrasado
		}

		subject := "Lembrete: " + l.NomeLembrete

		body := fmt.Sprintf(
			"Olá!\n\nLembrete: %s\nDescrição: %s\nVencimento: %s\nStatus atual: %s\n",
			l.NomeLembrete,
			l.Descricao,
			l.DataVencimento.In(now.Location()).Format("02/01/2006"),
			status,
		)

		if err := SendEmailWithContext(ctx, l.EmailNotificacao, subject, body); err != nil {
			log.Printf("erro ao enviar email para %s: %v", l.EmailNotificacao, err)
			continue
		}

		nextSendAt := buildNextDailySendAt(now)

		if err := service.UpdateSendControl(l.ID, status, now, nextSendAt); err != nil {
			log.Printf("erro ao atualizar controle de envio do lembrete %v: %v", l.ID, err)
			continue
		}

		log.Printf("email enviado com sucesso para: %s", l.EmailNotificacao)
	}

	return nil
}

func truncateDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func buildNextDailySendAt(now time.Time) time.Time {
	nextDay := now.AddDate(0, 0, 1)

	return time.Date(
		nextDay.Year(),
		nextDay.Month(),
		nextDay.Day(),
		reminderSendHour,
		0,
		0,
		0,
		nextDay.Location(),
	)
}