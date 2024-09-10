package repositories

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
	"gorm.io/gorm"
)

type OAuthProviderRepository struct {
	db *gorm.DB
}

func (r *OAuthProviderRepository) AddProvider(ctx context.Context, provider *models.OAuthProvider) error {
	return r.db.Create(provider).Error
}

func (r *OAuthProviderRepository) GetProvidersByUser(ctx context.Context, userID uint) ([]*models.OAuthProvider, error) {
	var providers []*models.OAuthProvider
	err := r.db.Where("user_id = ?", userID).Find(&providers).Error
	return providers, err
}

func NewOAuthProviderRepository(db *gorm.DB) models.OAuthProviderRepository {
	return &OAuthProviderRepository{db: db}
}
