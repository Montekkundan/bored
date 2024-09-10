package services

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
)

type ModerationVoteService struct {
	repository models.ModerationVoteRepository
}

func (s *ModerationVoteService) CastVote(ctx context.Context, vote *models.ModerationVote) error {
	return s.repository.CastVote(ctx, vote)
}

func (s *ModerationVoteService) GetVotes(ctx context.Context, messageID uint) ([]*models.ModerationVote, error) {
	return s.repository.GetVotes(ctx, messageID)
}

func NewModerationVoteService(repository models.ModerationVoteRepository) *ModerationVoteService {
	return &ModerationVoteService{repository: repository}
}
