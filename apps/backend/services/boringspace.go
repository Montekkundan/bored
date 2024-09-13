package services

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
)

type BoringSpaceService struct {
	repository models.BoringSpaceRepository
}

func NewBoringSpaceService(repository models.BoringSpaceRepository) models.BoringSpaceService {
	return &BoringSpaceService{
		repository: repository,
	}
}

func (s *BoringSpaceService) CreateBoringSpace(ctx context.Context, space *models.BoringSpace) error {
	// Start a transaction to ensure atomicity
	txErr := s.repository.CreateBoringSpace(ctx, space)
	if txErr != nil {
		return txErr
	}

	// Add the creator as an admin member
	member := &models.BoringSpaceMember{
		BoringSpaceID: space.ID,
		UserID:        space.CreatorID,
		Role:          models.BSAdmin,
	}
	return s.repository.AddMember(ctx, member)
}

func (s *BoringSpaceService) GetBoringSpaceByID(ctx context.Context, spaceID uint) (*models.BoringSpace, error) {
	return s.repository.GetBoringSpaceByID(ctx, spaceID)
}

func (s *BoringSpaceService) AddMember(ctx context.Context, member *models.BoringSpaceMember) error {
	return s.repository.AddMember(ctx, member)
}

func (s *BoringSpaceService) RemoveMember(ctx context.Context, spaceID uint, userID uint) error {
	return s.repository.RemoveMember(ctx, spaceID, userID)
}

func (s *BoringSpaceService) GetMembers(ctx context.Context, spaceID uint) ([]*models.BoringSpaceMember, error) {
	return s.repository.GetMembers(ctx, spaceID)
}

func (s *BoringSpaceService) UpdateMemberRole(ctx context.Context, spaceID uint, userID uint, role models.BoringSpaceRole) error {
	return s.repository.UpdateMemberRole(ctx, spaceID, userID, role)
}
