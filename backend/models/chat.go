package models

import (
	"context"
	"time"
)

type Chat struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Name      string    `json:"name"`
	IsGroup   bool      `json:"is_group" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
}

type ChatMember struct {
	ChatID   uint      `json:"chat_id" gorm:"primarykey"`
	UserID   uint      `json:"user_id" gorm:"primarykey"`
	JoinedAt time.Time `json:"joined_at" gorm:"default:now()"`
}

type Message struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	ChatID    uint      `json:"chat_id" gorm:"not null"`
	SenderID  uint      `json:"sender_id" gorm:"not null"`
	Content   string    `json:"content" gorm:"text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
}

type ChatRepository interface {
	CreateChat(ctx context.Context, chat *Chat) error
	AddMember(ctx context.Context, chatID uint, userID uint) error
	GetMessages(ctx context.Context, chatID uint) ([]*Message, error)
	SendMessage(ctx context.Context, message *Message) error
}
