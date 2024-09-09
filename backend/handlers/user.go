package handlers

import (
	"context"

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

func NewUserHandler(route fiber.Router, service models.UserService) {
	handler := &UserHandler{
		service: service,
	}

	route.Get("/get-all", handler.GetAllUsers)
	route.Put("/update-user", handler.UpdateUser)
}
