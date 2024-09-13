package services

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
)

type PublicMessageService struct {
	repo models.PublicMessageRepository
}

func NewPublicMessageService(repo models.PublicMessageRepository) models.PublicMessageService {
	return &PublicMessageService{repo: repo}
}

func (s *PublicMessageService) CreatePublicMessage(ctx context.Context, message *models.PublicMessage) error {
	return s.repo.CreatePublicMessage(ctx, message)
}

func (s *PublicMessageService) GetPublicMessageByID(ctx context.Context, messageID uint) (*models.PublicMessage, error) {
	return s.repo.GetPublicMessageByID(ctx, messageID)
}

func (s *PublicMessageService) GetPublicMessages(ctx context.Context, limit, offset int) ([]*models.PublicMessage, error) {
	return s.repo.GetPublicMessages(ctx, limit, offset)
}

func (s *PublicMessageService) DeletePublicMessage(ctx context.Context, messageID uint) error {
	return s.repo.DeletePublicMessage(ctx, messageID)
}

func (s *PublicMessageService) LikePublicMessage(ctx context.Context, messageID uint, userID uint) error {
	return s.repo.LikePublicMessage(ctx, messageID, userID)
}

func (s *PublicMessageService) UnlikePublicMessage(ctx context.Context, messageID uint, userID uint) error {
	return s.repo.UnlikePublicMessage(ctx, messageID, userID)
}

func (s *PublicMessageService) CreateComment(ctx context.Context, comment *models.Comment) error {
	return s.repo.CreateComment(ctx, comment)
}

func (s *PublicMessageService) GetCommentsByMessageID(ctx context.Context, messageID uint) ([]*models.Comment, error) {
	return s.repo.GetCommentsByMessageID(ctx, messageID)
}
