package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/montekkundan/bored/backend/models"
)

type ModerationVoteHandler struct {
	service models.ModerationVoteService
}

func (h *ModerationVoteHandler) CastVote(ctx *fiber.Ctx) error {
	vote := &models.ModerationVote{}
	if err := ctx.BodyParser(vote); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := h.service.CastVote(context.Background(), vote); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "message": "Vote cast successfully"})
}

func NewModerationVoteHandler(route fiber.Router, service models.ModerationVoteService) {
	handler := &ModerationVoteHandler{service: service}
	route.Post("/moderation/vote", handler.CastVote)
}
