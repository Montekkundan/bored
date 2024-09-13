package handlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/montekkundan/bored/backend/models"
)

type BoringSpaceHandler struct {
	service     models.BoringSpaceService
	userService models.UserService
}

func (h *BoringSpaceHandler) CreateBoringSpace(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)

	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid input",
		})
	}

	space := &models.BoringSpace{
		Name:        input.Name,
		Description: input.Description,
		CreatorID:   userID,
	}

	if err := h.service.CreateBoringSpace(context.Background(), space); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create BoringSpace",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   space,
	})
}

func (h *BoringSpaceHandler) GetBoringSpaceByID(ctx *fiber.Ctx) error {
	spaceIDStr := ctx.Params("id")
	spaceID, err := strconv.ParseUint(spaceIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid BoringSpace ID",
		})
	}

	space, err := h.service.GetBoringSpaceByID(context.Background(), uint(spaceID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "BoringSpace not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"data":   space,
	})
}

func (h *BoringSpaceHandler) AddMember(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)

	spaceIDStr := ctx.Params("id")
	spaceID, err := strconv.ParseUint(spaceIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid BoringSpace ID",
		})
	}

	// Check if the user has permission to add members
	space, err := h.service.GetBoringSpaceByID(context.Background(), uint(spaceID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "BoringSpace not found",
		})
	}

	isAdmin := false
	for _, member := range space.Members {
		if member.UserID == userID && member.Role == models.BSAdmin {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "fail",
			"message": "Permission denied",
		})
	}

	var input struct {
		UserID uint                   `json:"user_id" validate:"required"`
		Role   models.BoringSpaceRole `json:"role" validate:"required"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid input",
		})
	}

	member := &models.BoringSpaceMember{
		BoringSpaceID: uint(spaceID),
		UserID:        input.UserID,
		Role:          input.Role,
	}

	if err := h.service.AddMember(context.Background(), member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not add member",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   member,
	})
}

func (h *BoringSpaceHandler) RemoveMember(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)

	spaceIDStr := ctx.Params("id")
	spaceID, err := strconv.ParseUint(spaceIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid BoringSpace ID",
		})
	}

	targetUserIDStr := ctx.Params("userId")
	targetUserID, err := strconv.ParseUint(targetUserIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid User ID",
		})
	}

	// Check if the user has permission to remove members
	space, err := h.service.GetBoringSpaceByID(context.Background(), uint(spaceID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "BoringSpace not found",
		})
	}

	isAdmin := false
	for _, member := range space.Members {
		if member.UserID == userID && member.Role == models.BSAdmin {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "fail",
			"message": "Permission denied",
		})
	}

	if err := h.service.RemoveMember(context.Background(), uint(spaceID), uint(targetUserID)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not remove member",
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Member removed",
	})
}

func (h *BoringSpaceHandler) GetMembers(ctx *fiber.Ctx) error {
	spaceIDStr := ctx.Params("id")
	spaceID, err := strconv.ParseUint(spaceIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid BoringSpace ID",
		})
	}

	members, err := h.service.GetMembers(context.Background(), uint(spaceID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not retrieve members",
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"data":   members,
	})
}

func (h *BoringSpaceHandler) UpdateMemberRole(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)

	spaceIDStr := ctx.Params("id")
	spaceID, err := strconv.ParseUint(spaceIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid BoringSpace ID",
		})
	}

	targetUserIDStr := ctx.Params("userId")
	targetUserID, err := strconv.ParseUint(targetUserIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid User ID",
		})
	}

	// Check if the user has permission to update roles
	space, err := h.service.GetBoringSpaceByID(context.Background(), uint(spaceID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "BoringSpace not found",
		})
	}

	isAdmin := false
	for _, member := range space.Members {
		if member.UserID == userID && member.Role == models.BSAdmin {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "fail",
			"message": "Permission denied",
		})
	}

	var input struct {
		Role models.BoringSpaceRole `json:"role" validate:"required"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid input",
		})
	}

	if err := h.service.UpdateMemberRole(context.Background(), uint(spaceID), uint(targetUserID), input.Role); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not update member role",
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Member role updated",
	})
}

func NewBoringSpaceHandler(route fiber.Router, service models.BoringSpaceService, userService models.UserService) {
	handler := &BoringSpaceHandler{
		service:     service,
		userService: userService,
	}

	route.Post("/", handler.CreateBoringSpace)
	route.Get("/:id", handler.GetBoringSpaceByID)
	route.Post("/:id/members", handler.AddMember)
	route.Delete("/:id/members/:userId", handler.RemoveMember)
	route.Get("/:id/members", handler.GetMembers)
	route.Put("/:id/members/:userId/role", handler.UpdateMemberRole)
}
