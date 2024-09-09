package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/montekkundan/bored/backend/models"
)

var validate = validator.New()

type AuthHandler struct {
	authService models.AuthService
	userService models.UserService
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	creds := &models.AuthCredentials{}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := validate.Struct(creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "please provide valid credentials",
		})
	}

	token, user, err := h.authService.Login(context, creds)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Successfully logged in",
		"data": &fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

func (h *AuthHandler) Register(ctx *fiber.Ctx) error {
	creds := &models.AuthCredentials{}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := validate.Struct(creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": fmt.Errorf("please, provide a valid username, email and password").Error(),
		})
	}

	// Check if the username already exists
	_, err := h.userService.GetUserByUsername(context, creds.Username)
	if err == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Username already in use",
		})
	}

	// Check if the email already exists
	_, err = h.userService.GetUserByEmail(context, creds.Email)
	if err == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Email already in use",
		})
	}

	token, user, err := h.authService.Register(context, creds)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Successfully registered",
		"data": &fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

func (h *AuthHandler) VerifyEmail(ctx *fiber.Ctx) error {
	var payload struct {
		UserID uint `json:"user_id" validate:"required"`
	}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.authService.VerifyEmail(context, payload.UserID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Email verified successfully",
	})
}

func (h *AuthHandler) VerifyPhoneNumber(ctx *fiber.Ctx) error {
	var payload struct {
		UserID uint   `json:"user_id" validate:"required"`
		Code   string `json:"code" validate:"required"`
	}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.authService.VerifyPhoneNumber(context, payload.UserID, payload.Code); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Phone number verified successfully",
	})
}

func (h *AuthHandler) EnableTwoFactor(ctx *fiber.Ctx) error {
	var payload struct {
		UserID uint `json:"user_id" validate:"required"`
	}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.authService.EnableTwoFactor(context, payload.UserID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Two-factor authentication enabled successfully",
	})
}

func (h *AuthHandler) Logout(ctx *fiber.Ctx) error {
	// ask the client to discard the token

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Successfully logged out",
	})
}

func NewAuthHandler(route fiber.Router, authService models.AuthService, userService models.UserService) {
	handler := &AuthHandler{
		authService: authService,
		userService: userService,
	}

	route.Post("/login", handler.Login)
	route.Post("/register", handler.Register)
	route.Post("/verify-email", handler.VerifyEmail)
	route.Post("/verify-phone", handler.VerifyPhoneNumber)
	route.Post("/enable-2fa", handler.EnableTwoFactor)
	route.Post("/logout", handler.Logout)
}
