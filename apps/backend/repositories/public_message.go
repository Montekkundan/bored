package repositories

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
	"gorm.io/gorm"
)

type PublicMessageRepository struct {
	db *gorm.DB
}

func NewPublicMessageRepository(db *gorm.DB) models.PublicMessageRepository {
	return &PublicMessageRepository{db: db}
}

func (r *PublicMessageRepository) CreatePublicMessage(ctx context.Context, message *models.PublicMessage) error {
	return r.db.WithContext(ctx).Create(message).Error
}

func (r *PublicMessageRepository) GetPublicMessageByID(ctx context.Context, messageID uint) (*models.PublicMessage, error) {
	var message models.PublicMessage
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Likes").
		Preload("Comments.User").
		First(&message, messageID).Error
	return &message, err
}

func (r *PublicMessageRepository) GetPublicMessages(ctx context.Context, limit, offset int) ([]*models.PublicMessage, error) {
	var messages []*models.PublicMessage
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Likes").
		Preload("Comments.User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&messages).Error
	return messages, err
}

func (r *PublicMessageRepository) DeletePublicMessage(ctx context.Context, messageID uint) error {
	return r.db.WithContext(ctx).Delete(&models.PublicMessage{}, messageID).Error
}

func (r *PublicMessageRepository) LikePublicMessage(ctx context.Context, messageID uint, userID uint) error {
	var message models.PublicMessage
	if err := r.db.WithContext(ctx).First(&message, messageID).Error; err != nil {
		return err
	}

	return r.db.WithContext(ctx).Model(&message).Association("Likes").Append(&models.User{ID: userID})
}

func (r *PublicMessageRepository) UnlikePublicMessage(ctx context.Context, messageID uint, userID uint) error {
	var message models.PublicMessage
	if err := r.db.WithContext(ctx).First(&message, messageID).Error; err != nil {
		return err
	}

	return r.db.WithContext(ctx).Model(&message).Association("Likes").Delete(&models.User{ID: userID})
}

func (r *PublicMessageRepository) CreateComment(ctx context.Context, comment *models.Comment) error {
	return r.db.WithContext(ctx).Create(comment).Error
}

func (r *PublicMessageRepository) GetCommentsByMessageID(ctx context.Context, messageID uint) ([]*models.Comment, error) {
	var comments []*models.Comment
	err := r.db.WithContext(ctx).
		Preload("User").
		Where("public_message_id = ?", messageID).
		Order("created_at ASC").
		Find(&comments).Error
	return comments, err
}
