package middlewares

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/montekkundan/bored/backend/models"
	"gorm.io/gorm"
)

func AuthProtected(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Println("Middleware: AuthProtected invoked")
		authHeader := ctx.Get("Authorization")

		if authHeader == "" {
			log.Println("empty authorization header")

			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Unauthorized",
			})
		}

		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			log.Println("invalid token parts")

			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Unauthorized",
			})
		}

		tokenStr := tokenParts[1]
		secret := []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.GetSigningMethod("HS256").Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			log.Println("invalid token")

			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Unauthorized",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("Failed to parse token claims")
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Unauthorized",
			})
		}

		userId, ok := claims["id"].(float64)
		if !ok {
			log.Println("Invalid token ID format")
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Unauthorized",
			})
		}

		var user models.User
		if err := db.First(&user, uint(userId)).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Println("User not found in the database")
				return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
					"status":  "fail",
					"message": "Unauthorized",
				})
			}
			// Handle unexpected database errors
			return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Internal server error",
			})
		}

		ctx.Locals("userId", uint(userId))
		log.Printf("Middleware: UserID %v authenticated successfully", userId)

		return ctx.Next()
	}
}
