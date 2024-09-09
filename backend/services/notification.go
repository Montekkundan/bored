package services

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
)

type NotificationService struct {
	repository models.NotificationRepository
}

func (s *NotificationService) Create(ctx context.Context, notification *models.Notification) error {
	return s.repository.Create(ctx, notification)
}

func (s *NotificationService) MarkAsRead(ctx context.Context, notificationID uint) error {
	return s.repository.MarkAsRead(ctx, notificationID)
}

func (s *NotificationService) GetNotifications(ctx context.Context, userID uint) ([]*models.Notification, error) {
	return s.repository.GetNotifications(ctx, userID)
}

func NewNotificationService(repository models.NotificationRepository) *NotificationService {
	return &NotificationService{repository: repository}
}
