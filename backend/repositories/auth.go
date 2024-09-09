package repositories

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func (r *AuthRepository) RegisterUser(ctx context.Context, registerData *models.AuthCredentials) (*models.User, error) {
	user := &models.User{
		Email:        registerData.Email,
		PasswordHash: registerData.Password,
		PhoneNumber:  registerData.PhoneNumber,
		Username:     registerData.Username,
	}

	res := r.db.Model(&models.User{}).Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (r *AuthRepository) GetUser(ctx context.Context, query interface{}, args ...interface{}) (*models.User, error) {
	user := &models.User{}

	if res := r.db.Model(user).Where(query, args...).First(user); res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (r *AuthRepository) VerifyEmail(ctx context.Context, userID uint) error {
	user := &models.User{}
	if err := r.db.Model(user).Where("id = ?", userID).First(user).Error; err != nil {
		return err
	}

	user.EmailVerified = true
	return r.db.Save(user).Error
}

func (r *AuthRepository) VerifyPhoneNumber(ctx context.Context, userID uint, code string) error {
	user := &models.User{}
	if err := r.db.Model(user).Where("id = ?", userID).First(user).Error; err != nil {
		return err
	}

	// Logic to verify the phone number with the given code
	// Here, you could compare the 'code' with a code sent to the user's phone
	// For demonstration purposes, we'll assume the code is "123456"
	if code != "123456" {
		return gorm.ErrRecordNotFound // Or a custom error indicating the code is invalid
	}

	user.PhoneVerified = true
	return r.db.Save(user).Error
}

func NewAuthRepository(db *gorm.DB) models.AuthRepository {
	return &AuthRepository{
		db: db,
	}
}
