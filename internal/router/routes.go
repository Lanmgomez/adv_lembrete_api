package router

import (
	"adv_lembrete_api/internal/domain/auth"
	"adv_lembrete_api/internal/domain/users"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *auth.Handler, usersHandler *users.Handler,) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // libera tudo (dev)
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")

	// auth
	authGroup := api.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
	}

	// rotas protegidas
	protected := api.Group("/")
	protected.Use(AuthMiddleware())

	{
		protected.POST("/logout", authHandler.Logout)

		// users
		protected.GET("/users", usersHandler.GetUsers)
		protected.GET("/users/:id", usersHandler.GetUserByID)
		protected.POST("/users", usersHandler.CreateUser)
		protected.PUT("/users/:id", usersHandler.UpdateUser)
		protected.DELETE("/users/:id", usersHandler.DeleteUser)
	}

	return r
}