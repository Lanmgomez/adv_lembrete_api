package main

import (
	"context"
	"log"
	"time"

	"adv_lembrete_api/database/configuration"
	"adv_lembrete_api/internal/domain/entidades"
	"adv_lembrete_api/internal/domain/lembretes"
	"adv_lembrete_api/internal/utils"
)

/*
	Na Render, este binário deve ser executado como Cron Job.
	Cada execução faz:

	1. conecta no banco
	2. busca lembretes vencidos
	3. envia os e-mails
	4. atualiza o banco
	5. encerra o processo
*/

func setupService() *lembretes.Service {
	db := configuration.ConnectDB()

	entidadesRepo := entidades.NewRepository(db)

	lembretesRepo := lembretes.NewRepository(db)
	lembretesService := lembretes.NewService(lembretesRepo, entidadesRepo)

	log.Println("Reminder worker initialized")

	return lembretesService
}

func getCurrentTime() time.Time {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Printf("Erro ao carregar timezone America/Sao_Paulo, usando timezone local: %v", err)
		return time.Now()
	}

	return time.Now().In(loc)
}

func main() {
	lembretesService := setupService()

	now := getCurrentTime()

	log.Println("Rodando worker de lembretes em:", now.Format("2006-01-02 15:04:05"))

	ctx := context.Background()

	if err := utils.ProcessDueReminders(ctx, lembretesService, now); err != nil {
		log.Fatalf("Erro ao processar lembretes: %v", err)
	}

	log.Println("Worker de lembretes finalizado com sucesso")
}