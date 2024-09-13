package services

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
)

type UserService struct {
	repository models.UserRepository
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return s.repository.GetAllUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, userID uint) (*models.User, error) {
	return s.repository.GetUserByID(ctx, userID)
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.repository.GetUserByUsername(ctx, username)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repository.GetUserByEmail(ctx, email)
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	return s.repository.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUserByID(ctx context.Context, userID uint) error {
	return s.repository.DeleteUserByID(ctx, userID)
}

func (s *UserService) DeactivateUser(ctx context.Context, userID uint) error {
	return s.repository.DeactivateUser(ctx, userID)
}

func (s *UserService) GetUserBoringSpaces(ctx context.Context, userID uint) ([]*models.BoringSpaceMember, error) {
	return s.repository.GetUserBoringSpaces(ctx, userID)
}

func (s *UserService) GetAllPublicMessages(ctx context.Context, limit, offset int) ([]*models.PublicMessage, error) {
	return s.repository.GetAllPublicMessages(ctx, limit, offset)
}

func NewUserService(repository models.UserRepository) models.UserService {
	return &UserService{
		repository: repository,
	}
}
