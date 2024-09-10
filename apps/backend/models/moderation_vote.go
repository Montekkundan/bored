package models

import (
	"context"
	"time"
)

type ModerationVote struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	UserID    uint      `json:"user_id"`
	ContentID uint      `json:"content_id"`
	VoteType  string    `json:"vote_type"` // 'flag', 'approve', 'reject'
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
}

type ModerationVoteRepository interface {
	CastVote(ctx context.Context, vote *ModerationVote) error
	GetVotes(ctx context.Context, messageID uint) ([]*ModerationVote, error)
}

type ModerationVoteService interface {
	CastVote(ctx context.Context, vote *ModerationVote) error
	GetVotes(ctx context.Context, messageID uint) ([]*ModerationVote, error)
}
