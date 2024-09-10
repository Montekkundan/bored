package models

import (
	"context"
	"time"
)

type Notification struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
}

type NotificationRepository interface {
	Create(ctx context.Context, notification *Notification) error
	MarkAsRead(ctx context.Context, notificationID uint) error
	GetNotifications(ctx context.Context, userID uint) ([]*Notification, error)
}
type NotificationService interface {
	Create(ctx context.Context, notification *Notification) error
	MarkAsRead(ctx context.Context, notificationID uint) error
	GetNotifications(ctx context.Context, userID uint) ([]*Notification, error)
}
