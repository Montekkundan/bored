package models

import (
	"context"
	"time"
)

type OAuthProvider struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	UserID     uint      `json:"user_id"`
	Provider   string    `json:"provider"`    // e.g., GitHub, Google
	ProviderID string    `json:"provider_id"` // The provider-specific user ID
	Token      string    `json:"token"`       // OAuth access token
	CreatedAt  time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:now()"`
}

type OAuthProviderRepository interface {
	AddProvider(ctx context.Context, provider *OAuthProvider) error
	GetProvidersByUser(ctx context.Context, userID uint) ([]*OAuthProvider, error)
}
