package repositories

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
	"gorm.io/gorm"
)

type BoringSpaceRepository struct {
	db *gorm.DB
}

func NewBoringSpaceRepository(db *gorm.DB) models.BoringSpaceRepository {
	return &BoringSpaceRepository{
		db: db,
	}
}

func (r *BoringSpaceRepository) CreateBoringSpace(ctx context.Context, space *models.BoringSpace) error {
	return r.db.WithContext(ctx).Create(space).Error
}

func (r *BoringSpaceRepository) GetBoringSpaceByID(ctx context.Context, spaceID uint) (*models.BoringSpace, error) {
	var space models.BoringSpace
	err := r.db.WithContext(ctx).
		Preload("Creator").
		Preload("Members.User").
		First(&space, "id = ?", spaceID).Error
	if err != nil {
		return nil, err
	}
	return &space, nil
}

func (r *BoringSpaceRepository) AddMember(ctx context.Context, member *models.BoringSpaceMember) error {
	return r.db.WithContext(ctx).Create(member).Error
}

func (r *BoringSpaceRepository) RemoveMember(ctx context.Context, spaceID uint, userID uint) error {
	return r.db.WithContext(ctx).
		Where("boring_space_id = ? AND user_id = ?", spaceID, userID).
		Delete(&models.BoringSpaceMember{}).Error
}

func (r *BoringSpaceRepository) GetMembers(ctx context.Context, spaceID uint) ([]*models.BoringSpaceMember, error) {
	var members []*models.BoringSpaceMember
	err := r.db.WithContext(ctx).
		Preload("User").
		Where("boring_space_id = ?", spaceID).
		Find(&members).Error
	return members, err
}

func (r *BoringSpaceRepository) UpdateMemberRole(ctx context.Context, spaceID uint, userID uint, role models.BoringSpaceRole) error {
	return r.db.WithContext(ctx).
		Model(&models.BoringSpaceMember{}).
		Where("boring_space_id = ? AND user_id = ?", spaceID, userID).
		Update("role", role).Error
}
