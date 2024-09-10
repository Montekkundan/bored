package repositories

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
	"gorm.io/gorm"
)

type ModerationVoteRepository struct {
	db *gorm.DB
}

func (r *ModerationVoteRepository) CastVote(ctx context.Context, vote *models.ModerationVote) error {
	return r.db.Create(vote).Error
}

func (r *ModerationVoteRepository) GetVotes(ctx context.Context, messageID uint) ([]*models.ModerationVote, error) {
	var votes []*models.ModerationVote
	err := r.db.Where("message_id = ?", messageID).Find(&votes).Error
	return votes, err
}

func NewModerationVoteRepository(db *gorm.DB) models.ModerationVoteRepository {
	return &ModerationVoteRepository{db: db}
}
