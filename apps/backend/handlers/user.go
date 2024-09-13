package handlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/montekkundan/bored/backend/models"
)

type UserHandler struct {
	service models.UserService
}

func (h *UserHandler) GetAllUsers(ctx *fiber.Ctx) error {
	userIdValue := ctx.Locals("userId")
	if userIdValue == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "User ID not found in the context",
		})
	}
	userID, ok := userIdValue.(uint)

	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Invalid user ID",
		})
	}

	user, err := h.service.GetUserByID(context.Background(), userID)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Unauthorized",
		})
	}

	isAdmin := false
	for _, role := range user.Roles {
		if models.UserRole(role) == models.Admin {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Unauthorized",
		})
	}

	users, err := h.service.GetAllUsers(context.Background())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   users,
	})
}

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)
	user, err := h.service.GetUserByID(context.Background(), userID)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": "User not found",
		})
	}

	var updateData struct {
		Bio            string   `json:"bio"`
		Interests      []string `json:"interests"`
		ProfilePicture string   `json:"profile_picture"`
		CoverPhoto     string   `json:"cover_photo"`
		SocialLinks    string   `json:"social_links"`
		Latitude       float64  `json:"latitude"`
		Longitude      float64  `json:"longitude"`
		AudioEnabled   bool     `json:"audio_enabled"`
		VideoEnabled   bool     `json:"video_enabled"`
	}

	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if updateData.Bio != "" {
		user.Bio = updateData.Bio
	}
	if updateData.Interests != nil {
		user.Interests = updateData.Interests
	}
	if updateData.ProfilePicture != "" {
		user.ProfilePicture = updateData.ProfilePicture
	}
	if updateData.CoverPhoto != "" {
		user.CoverPhoto = updateData.CoverPhoto
	}
	if updateData.SocialLinks != "" {
		user.SocialLinks = updateData.SocialLinks
	}
	if updateData.Latitude != 0 {
		user.Latitude = updateData.Latitude
	}
	if updateData.Longitude != 0 {
		user.Longitude = updateData.Longitude
	}
	user.AudioEnabled = updateData.AudioEnabled
	user.VideoEnabled = updateData.VideoEnabled

	if err := h.service.UpdateUser(context.Background(), user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User updated successfully",
	})
}

func (h *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)

	err := h.service.DeleteUserByID(context.Background(), userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Failed to delete user",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User deleted successfully",
	})
}

func (h *UserHandler) DeactivateAccount(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)

	err := h.service.DeactivateUser(context.Background(), userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Failed to deactivate account",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Account deactivated successfully",
	})
}

func (h *UserHandler) AdminDeleteUser(ctx *fiber.Ctx) error {
	adminUser := ctx.Locals("userId").(uint)

	// Check if the user has Admin privileges
	user, err := h.service.GetUserByID(context.Background(), adminUser)
	if err != nil || !user.HasRole(models.Admin) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Admin privileges required",
		})
	}

	targetUserIDStr := ctx.Params("id")
	targetUserID, err := strconv.ParseUint(targetUserIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Invalid user ID format",
		})
	}

	err = h.service.DeleteUserByID(context.Background(), uint(targetUserID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Failed to delete user",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User deleted successfully",
	})
}

func (h *UserHandler) GetUserBoringSpaces(ctx *fiber.Ctx) error {
	userIDValue := ctx.Locals("userId")
	if userIDValue == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "User ID not found in the context",
		})
	}
	userID, ok := userIDValue.(uint)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Invalid user ID",
		})
	}

	boringSpaces, err := h.service.GetUserBoringSpaces(context.Background(), userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "error",
			"message": "Could not retrieve BoringSpaces",
		})
	}

	return ctx.JSON(&fiber.Map{
		"status": "success",
		"data":   boringSpaces,
	})
}

func (h *UserHandler) GetAllPublicMessages(ctx *fiber.Ctx) error {
	limit, _ := strconv.Atoi(ctx.Query("limit", "20"))
	offset, _ := strconv.Atoi(ctx.Query("offset", "0"))

	messages, err := h.service.GetAllPublicMessages(context.Background(), limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to retrieve messages"})
	}

	return ctx.JSON(fiber.Map{"status": "success", "data": messages})
}

func NewUserHandler(route fiber.Router, service models.UserService) {
	handler := &UserHandler{
		service: service,
	}

	route.Get("/get-all", handler.GetAllUsers)
	route.Put("/update-user", handler.UpdateUser)
	route.Delete("/delete", handler.DeleteUser)
	route.Put("/deactivate-account", handler.DeactivateAccount)
	route.Delete("/admin-delete/:id", handler.AdminDeleteUser)
	route.Get("/boringspaces", handler.GetUserBoringSpaces)
	route.Get("/public-messages", handler.GetAllPublicMessages)
}
