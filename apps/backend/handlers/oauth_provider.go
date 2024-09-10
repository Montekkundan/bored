package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/montekkundan/bored/backend/models"
)

type OAuthProviderHandler struct {
	repository models.OAuthProviderRepository
}

func (h *OAuthProviderHandler) AddProvider(ctx *fiber.Ctx) error {
	provider := &models.OAuthProvider{}
	if err := ctx.BodyParser(provider); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := h.repository.AddProvider(context.Background(), provider); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "message": "Provider added successfully"})
}

func NewOAuthProviderHandler(route fiber.Router, repository models.OAuthProviderRepository) {
	handler := &OAuthProviderHandler{repository: repository}
	route.Post("/oauth/provider", handler.AddProvider)
}
