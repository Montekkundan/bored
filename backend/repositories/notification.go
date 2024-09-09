package repositories

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func (r *NotificationRepository) Create(ctx context.Context, notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *NotificationRepository) MarkAsRead(ctx context.Context, notificationID uint) error {
	return r.db.Model(&models.Notification{}).Where("id = ?", notificationID).Update("is_read", true).Error
}

func (r *NotificationRepository) GetNotifications(ctx context.Context, userID uint) ([]*models.Notification, error) {
	var notifications []*models.Notification
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&notifications).Error
	return notifications, err
}

func NewNotificationRepository(db *gorm.DB) models.NotificationRepository {
	return &NotificationRepository{db: db}
}
