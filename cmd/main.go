package main

import (
	"adv_lembrete_api/database/configuration"
	"adv_lembrete_api/internal/domain/auth"
	"adv_lembrete_api/internal/domain/entidades"
	"adv_lembrete_api/internal/domain/users"
	"adv_lembrete_api/internal/router"
)

func main() {
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

	r := router.SetupRouter(authHandler, usersHandler, entidadesHandler)

	r.Run(":8080")
}