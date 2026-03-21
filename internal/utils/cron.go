package utils

import (
	"adv_lembrete_api/internal/domain/lembretes"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)
func StartReminderJob(service *lembretes.Service) {

	c := cron.New()

	// roda todo dia às 06:00
	c.AddFunc("0 6 * * *", func() {

		lembretes, _, _ := service.GetAllLembretes("", 1, 1000)

		for _, l := range lembretes {

			if l.Status == "concluido" {
				continue
			}

			// calcula quando começa envio
			start := l.DataVencimento.AddDate(0, 0, -l.DiasAntecedencia)
			hoje := time.Now()

			if hoje.After(start) {

				// aqui entra envio de email
				fmt.Println("Enviar email para:", l.EmailNotificacao)

				// depois você integra com SMTP / serviço real
			}
		}
	})

	c.Start()
}