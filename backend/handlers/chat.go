package handlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/montekkundan/bored/backend/models"
)

type ChatHandler struct {
	repository models.ChatRepository
}

func (h *ChatHandler) CreateChat(ctx *fiber.Ctx) error {
	chat := &models.Chat{}
	if err := ctx.BodyParser(chat); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := h.repository.CreateChat(context.Background(), chat); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "message": "Chat created successfully"})
}

func (h *ChatHandler) AddMember(ctx *fiber.Ctx) error {
	member := &models.ChatMember{}
	if err := ctx.BodyParser(member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := h.repository.AddMember(context.Background(), member.ChatID, member.UserID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "message": "Member added to chat"})
}

func (h *ChatHandler) SendMessage(ctx *fiber.Ctx) error {
	message := &models.Message{}
	if err := ctx.BodyParser(message); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := h.repository.SendMessage(context.Background(), message); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "message": "Message sent successfully"})
}

func (h *ChatHandler) GetMessages(ctx *fiber.Ctx) error {
	chatID, err := strconv.ParseUint(ctx.Params("chatID"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": "Invalid chat ID"})
	}

	messages, err := h.repository.GetMessages(context.Background(), uint(chatID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "data": messages})
}

func NewChatHandler(route fiber.Router, repository models.ChatRepository) {
	handler := &ChatHandler{repository: repository}
	route.Post("/", handler.CreateChat)
	route.Post("/member", handler.AddMember)
	route.Post("/message", handler.SendMessage)
	route.Get("/:chatID/messages", handler.GetMessages)
}
