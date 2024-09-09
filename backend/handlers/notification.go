package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/montekkundan/bored/backend/models"
)

type NotificationHandler struct {
	service models.NotificationService
}

func (h *NotificationHandler) GetNotifications(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)
	notifications, err := h.service.GetNotifications(context.Background(), userID)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "data": notifications})
}

func NewNotificationHandler(route fiber.Router, service models.NotificationService) {
	handler := &NotificationHandler{service: service}
	route.Get("/", handler.GetNotifications)
}
