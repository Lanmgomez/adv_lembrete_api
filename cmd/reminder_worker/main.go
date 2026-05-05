package main

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"adv_lembrete_api/database/configuration"
	"adv_lembrete_api/internal/domain/entidades"
	"adv_lembrete_api/internal/domain/lembretes"
	"adv_lembrete_api/internal/utils"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	lembretesService *lembretes.Service
	once             sync.Once
)

/*  Essa Lambda não tem Gin, não tem API Gateway e não tem rota HTTP.
	Ela só faz isso: 
	foi chamada -> conecta no banco -> busca lembretes -> envia emails -> atualiza banco 
*/

func setupService() {
	db := configuration.ConnectDB()

	entidadesRepo := entidades.NewRepository(db)

	lembretesRepo := lembretes.NewRepository(db)
	lembretesService = lembretes.NewService(lembretesRepo, entidadesRepo)

	log.Println("Reminder worker initialized")
}

func Handler(ctx context.Context, event json.RawMessage) error {
	once.Do(setupService)

	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		loc = time.Local
	}

	now := time.Now().In(loc)

	log.Println("Rodando worker de lembretes em:", now.Format("2006-01-02 15:04:05"))

	return utils.ProcessDueReminders(ctx, lembretesService, now)
}

func main() {
	lambda.Start(Handler)
}
