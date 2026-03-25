package utils

import (
	"adv_lembrete_api/internal/domain/lembretes"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func StartReminderJob(service *lembretes.Service) {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		loc = time.Local
	}

	c := cron.New(cron.WithLocation(loc))

	// envia todo dia as 8h
	_, err = c.AddFunc("0 8 * * *", func() {
		now := time.Now().In(loc)
		fmt.Println("Rodando job de lembretes em:", now.Format("2006-01-02 15:04:05"))

		items, err := service.FindDueForSend(now)
		if err != nil {
			fmt.Println("erro ao buscar lembretes para envio:", err)
			return
		}

		for _, l := range items {
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
	})

	if err != nil {
		fmt.Println("erro ao registrar cron de lembretes:", err)
		return
	}

	c.Start()
	fmt.Println("Reminder job iniciado: execução diária às 06:00")
}

func truncateDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func buildNextDailySendAt(now time.Time) time.Time {
	nextDay := now.AddDate(0, 0, 1)
	return time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 6, 0, 0, 0, nextDay.Location())
}