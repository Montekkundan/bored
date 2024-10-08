package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/montekkundan/bored/backend/config"
	"github.com/montekkundan/bored/backend/db"
	_ "github.com/montekkundan/bored/backend/docs"
	"github.com/montekkundan/bored/backend/handlers"
	"github.com/montekkundan/bored/backend/middlewares"
	"github.com/montekkundan/bored/backend/repositories"
	"github.com/montekkundan/bored/backend/services"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", envConfig.RedisHost, envConfig.RedisPort),
	})

	app := fiber.New(fiber.Config{
		AppName:      "Bored",
		ServerHeader: "Fiber",
	})

	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)
	userRepository := repositories.NewUserRepository(db)
	chatRepository := repositories.NewChatRepository(db)
	oauthProviderRepository := repositories.NewOAuthProviderRepository(db)
	refreshTokenRepository := repositories.NewRefreshTokenRepository(db)
	boringSpaceRepository := repositories.NewBoringSpaceRepository(db)
	publicMessageRepository := repositories.NewPublicMessageRepository(db)

	// Service
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(authRepository, userService, *envConfig, refreshTokenRepository, redisClient)
	notificationService := services.NewNotificationService(repositories.NewNotificationRepository(db))
	moderationVoteService := services.NewModerationVoteService(repositories.NewModerationVoteRepository(db))
	boringSpaceService := services.NewBoringSpaceService(boringSpaceRepository)
	publicMessageService := services.NewPublicMessageService(publicMessageRepository)

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
	handlers.NewBoringSpaceHandler(privateRoutes.Group("/boringspaces"), boringSpaceService, userService)
	handlers.NewPublicMessageHandler(privateRoutes.Group("/public-messages"), publicMessageService, userService)

	app.Listen(fmt.Sprintf(":%s", envConfig.ServerPort))
}
