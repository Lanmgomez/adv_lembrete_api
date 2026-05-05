package utils

import (
	"context"
	"fmt"
	"time"

	"adv_lembrete_api/internal/domain/lembretes"
)

const reminderSendHour = 8

func ProcessDueReminders(ctx context.Context, service *lembretes.Service, now time.Time) error {
	items, err := service.FindDueForSend(now)
	if err != nil {
		return fmt.Errorf("erro ao buscar lembretes para envio: %w", err)
	}

	for _, l := range items {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if l.Status != "pendente" && l.Status != "atrasado" {
			continue
		}

		status := l.Status
		today := truncateDate(now)
		due := truncateDate(l.DataVencimento)

		if today.After(due) {
			status = "atrasado"
		}

		subject := "Lembrete: " + l.NomeLembrete

		body := fmt.Sprintf(
			"Olá!\n\nLembrete: %s\nDescrição: %s\nVencimento: %s\nStatus atual: %s\n",
			l.NomeLembrete,
			l.Descricao,
			l.DataVencimento.Format("2006-01-02"),
			status,
		)

		err := SendEmail(l.EmailNotificacao, subject, body)
		if err != nil {
			fmt.Println("erro ao enviar email para", l.EmailNotificacao, ":", err)
			continue
		}

		nextSendAt := buildNextDailySendAt(now)

		err = service.UpdateSendControl(l.ID, status, now, nextSendAt)
		if err != nil {
			fmt.Println("erro ao atualizar controle de envio do lembrete", l.ID, ":", err)
			continue
		}

		fmt.Println("email enviado com sucesso para:", l.EmailNotificacao)
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