package models

import (
	"context"
	"time"
)

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"not null;unique"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RefreshTokenRepository interface {
	Create(ctx context.Context, refreshToken *RefreshToken) error
	FindByToken(ctx context.Context, token string) (*RefreshToken, error)
	Delete(ctx context.Context, token string) error
	InvalidateOldTokens(ctx context.Context, userID uint) error
}
