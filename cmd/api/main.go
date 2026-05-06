package main

import (
	"context"
	"log"
	"sync"

	"adv_lembrete_api/database/configuration"
	"adv_lembrete_api/internal/domain/auth"
	"adv_lembrete_api/internal/domain/entidades"
	"adv_lembrete_api/internal/domain/lembretes"
	"adv_lembrete_api/internal/domain/users"
	"adv_lembrete_api/internal/router"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var (
    ginLambda *ginadapter.GinLambdaV2
    once      sync.Once
)

func setupRouter() *gin.Engine {
    db := configuration.ConnectDB()

    // auth
    authRepo := auth.NewRepository(db)
    authService := auth.NewService(authRepo)
    authHandler := auth.NewHandler(authService)

    // users
    usersRepo := users.NewRepository(db)
    usersService := users.NewService(usersRepo)
    usersHandler := users.NewHandler(usersService)

    // entidades
    entidadesRepo := entidades.NewRepository(db)
    entidadesService := entidades.NewService(entidadesRepo)
    entidadesHandler := entidades.NewHandler(entidadesService)

    // lembretes
    lembretesRepo := lembretes.NewRepository(db)
    lembretesService := lembretes.NewService(lembretesRepo, entidadesRepo)
    lembretesHandler := lembretes.NewHandler(lembretesService)

    return router.SetupRouter(authHandler, usersHandler, entidadesHandler, lembretesHandler)

}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
    once.Do(func() {
        r := setupRouter()

        ginLambda = ginadapter.NewV2(r)

        log.Println("Lambda initialized")
    })

    return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
    lambda.Start(Handler)
}