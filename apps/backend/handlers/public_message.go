package handlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/montekkundan/bored/backend/models"
)

type PublicMessageHandler struct {
	service     models.PublicMessageService
	userService models.UserService
}

func (h *PublicMessageHandler) CreatePublicMessage(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)

	var input struct {
		Content  string `json:"content" validate:"required"`
		MediaURL string `json:"media_url"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid input"})
	}

	if input.Content == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Content is required"})
	}

	message := &models.PublicMessage{
		UserID:   userID,
		Content:  input.Content,
		MediaURL: input.MediaURL,
	}

	if err := h.service.CreatePublicMessage(context.Background(), message); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to create message"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": message})
}

// GetPublicMessages handles GET /public-messages
func (h *PublicMessageHandler) GetPublicMessages(ctx *fiber.Ctx) error {
	limit, _ := strconv.Atoi(ctx.Query("limit", "20"))
	offset, _ := strconv.Atoi(ctx.Query("offset", "0"))

	messages, err := h.service.GetPublicMessages(context.Background(), limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to retrieve messages"})
	}

	return ctx.JSON(fiber.Map{"status": "success", "data": messages})
}

// GetPublicMessageByID handles GET /public-messages/:id
func (h *PublicMessageHandler) GetPublicMessageByID(ctx *fiber.Ctx) error {
	messageIDStr := ctx.Params("id")
	messageID, err := strconv.ParseUint(messageIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid message ID"})
	}

	message, err := h.service.GetPublicMessageByID(context.Background(), uint(messageID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Message not found"})
	}

	return ctx.JSON(fiber.Map{"status": "success", "data": message})
}

// DeletePublicMessage handles DELETE /public-messages/:id
func (h *PublicMessageHandler) DeletePublicMessage(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)
	messageIDStr := ctx.Params("id")
	messageID, err := strconv.ParseUint(messageIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid message ID"})
	}

	message, err := h.service.GetPublicMessageByID(context.Background(), uint(messageID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Message not found"})
	}

	// Check if the user is the author or an admin
	if message.UserID != userID {
		requestingUser, err := h.userService.GetUserByID(context.Background(), userID)
		if err != nil || !requestingUser.HasRole(models.Admin) {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "Permission denied"})
		}
	}

	if err := h.service.DeletePublicMessage(context.Background(), uint(messageID)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to delete message"})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Message deleted"})
}

// LikePublicMessage handles POST /public-messages/:id/like
func (h *PublicMessageHandler) LikePublicMessage(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)
	messageIDStr := ctx.Params("id")
	messageID, err := strconv.ParseUint(messageIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid message ID"})
	}

	if err := h.service.LikePublicMessage(context.Background(), uint(messageID), userID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to like message"})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Message liked"})
}

// UnlikePublicMessage handles POST /public-messages/:id/unlike
func (h *PublicMessageHandler) UnlikePublicMessage(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)
	messageIDStr := ctx.Params("id")
	messageID, err := strconv.ParseUint(messageIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid message ID"})
	}

	if err := h.service.UnlikePublicMessage(context.Background(), uint(messageID), userID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to unlike message"})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Message unliked"})
}

// CreateComment handles POST /public-messages/:id/comments
func (h *PublicMessageHandler) CreateComment(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)
	messageIDStr := ctx.Params("id")
	messageID, err := strconv.ParseUint(messageIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid message ID"})
	}

	var input struct {
		Content string `json:"content" validate:"required"`
	}

	if err := ctx.BodyParser(&input); err != nil || input.Content == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Content is required"})
	}

	comment := &models.Comment{
		PublicMessageID: uint(messageID),
		UserID:          userID,
		Content:         input.Content,
	}

	if err := h.service.CreateComment(context.Background(), comment); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to create comment"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": comment})
}

// GetComments handles GET /public-messages/:id/comments
func (h *PublicMessageHandler) GetComments(ctx *fiber.Ctx) error {
	messageIDStr := ctx.Params("id")
	messageID, err := strconv.ParseUint(messageIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid message ID"})
	}

	comments, err := h.service.GetCommentsByMessageID(context.Background(), uint(messageID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to retrieve comments"})
	}

	return ctx.JSON(fiber.Map{"status": "success", "data": comments})
}

func NewPublicMessageHandler(route fiber.Router, service models.PublicMessageService, userService models.UserService) {
	handler := &PublicMessageHandler{
		service:     service,
		userService: userService,
	}

	route.Post("/", handler.CreatePublicMessage)
	route.Get("/", handler.GetPublicMessages)
	route.Get("/:id", handler.GetPublicMessageByID)
	route.Delete("/:id", handler.DeletePublicMessage)
	route.Post("/:id/like", handler.LikePublicMessage)
	route.Post("/:id/unlike", handler.UnlikePublicMessage)
	route.Post("/:id/comments", handler.CreateComment)
	route.Get("/:id/comments", handler.GetComments)
}
