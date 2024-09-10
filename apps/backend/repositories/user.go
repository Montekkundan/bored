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

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &UserRepository{
		db: db,
	}
}
