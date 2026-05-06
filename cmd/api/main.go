package main

import (
	"log"

	"adv_lembrete_api/database/configuration"
	"adv_lembrete_api/internal/domain/auth"
	"adv_lembrete_api/internal/domain/entidades"
	"adv_lembrete_api/internal/domain/lembretes"
	"adv_lembrete_api/internal/domain/users"
	"adv_lembrete_api/internal/router"

	"github.com/gin-gonic/gin"
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

func main() {
	r := setupRouter()
	port := ":8080"

	log.Printf("Servidor iniciando na porta %s", port)

	if err := r.Run(port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}