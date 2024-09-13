package repositories

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID uint) (*models.User, error) {
	user := &models.User{}
	if err := r.db.Where("id = ?", userID).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	if err := r.db.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUserByID(ctx context.Context, userID uint) error {
	return r.db.Delete(&models.User{}, userID).Error
}

func (r *UserRepository) DeactivateUser(ctx context.Context, userID uint) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("deactivated", true).Error
}

func (r *UserRepository) GetUserBoringSpaces(ctx context.Context, userID uint) ([]*models.BoringSpaceMember, error) {
	var memberships []*models.BoringSpaceMember
	err := r.db.WithContext(ctx).
		Preload("BoringSpace").
		Where("user_id = ?", userID).
		Find(&memberships).Error
	if err != nil {
		return nil, err
	}
	return memberships, nil
}

func (r *UserRepository) GetAllPublicMessages(ctx context.Context, limit, offset int) ([]*models.PublicMessage, error) {
	var messages []*models.PublicMessage
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Likes").
		Preload("Comments").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&messages).Error
	return messages, err
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &UserRepository{
		db: db,
	}
}
