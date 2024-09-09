package repositories

import (
	"context"

	"github.com/montekkundan/bored/backend/models"
	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func (r *ChatRepository) CreateChat(ctx context.Context, chat *models.Chat) error {
	return r.db.Create(chat).Error
}

func (r *ChatRepository) AddMember(ctx context.Context, chatID uint, userID uint) error {
	member := &models.ChatMember{ChatID: chatID, UserID: userID}
	return r.db.Create(member).Error
}

func (r *ChatRepository) GetMessages(ctx context.Context, chatID uint) ([]*models.Message, error) {
	var messages []*models.Message
	err := r.db.Where("chat_id = ?", chatID).Order("created_at asc").Find(&messages).Error
	return messages, err
}

func (r *ChatRepository) SendMessage(ctx context.Context, message *models.Message) error {
	return r.db.Create(message).Error
}

func NewChatRepository(db *gorm.DB) models.ChatRepository {
	return &ChatRepository{db: db}
}
