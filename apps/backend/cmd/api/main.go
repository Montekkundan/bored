package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/montekkundan/bored/backend/config"
	"github.com/montekkundan/bored/backend/db"
	"github.com/montekkundan/bored/backend/handlers"
	"github.com/montekkundan/bored/backend/middlewares"
	"github.com/montekkundan/bored/backend/repositories"
	"github.com/montekkundan/bored/backend/services"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Bored",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)
	userRepository := repositories.NewUserRepository(db)
	chatRepository := repositories.NewChatRepository(db)
	oauthProviderRepository := repositories.NewOAuthProviderRepository(db)

	// Service
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(authRepository, userService)
	notificationService := services.NewNotificationService(repositories.NewNotificationRepository(db))
	moderationVoteService := services.NewModerationVoteService(repositories.NewModerationVoteRepository(db))

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService, userService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	// Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)
	handlers.NewChatHandler(privateRoutes.Group("/chat"), chatRepository)
	handlers.NewOAuthProviderHandler(privateRoutes.Group("/oauth"), oauthProviderRepository)
	handlers.NewNotificationHandler(privateRoutes.Group("/notifications"), notificationService)
	handlers.NewModerationVoteHandler(privateRoutes.Group("/moderation"), moderationVoteService)
	handlers.NewUserHandler(privateRoutes.Group("/users"), userService)

	app.Listen(fmt.Sprintf(":%s", envConfig.ServerPort))
}
