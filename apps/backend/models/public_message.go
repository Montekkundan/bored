package models

import (
	"context"
	"time"
)

type PublicMessage struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	UserID    uint      `json:"user_id" gorm:"not null"` // The user who posted the message
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Content   string    `json:"content" gorm:"text;not null"`
	MediaURL  string    `json:"media_url" gorm:"text"` // Optional media (image/video)
	Likes     []*User   `json:"likes" gorm:"many2many:public_message_likes"`
	Comments  []Comment `json:"comments" gorm:"foreignKey:PublicMessageID"`
	Shares    int       `json:"shares" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Comment struct {
	ID              uint          `json:"id" gorm:"primarykey"`
	PublicMessageID uint          `json:"public_message_id" gorm:"not null"`
	PublicMessage   PublicMessage `json:"public_message" gorm:"foreignKey:PublicMessageID"`
	UserID          uint          `json:"user_id" gorm:"not null"`
	User            User          `json:"user" gorm:"foreignKey:UserID"`
	Content         string        `json:"content" gorm:"text;not null"`
	CreatedAt       time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
}

type PublicMessageRepository interface {
	CreatePublicMessage(ctx context.Context, message *PublicMessage) error
	GetPublicMessageByID(ctx context.Context, messageID uint) (*PublicMessage, error)
	GetPublicMessages(ctx context.Context, limit, offset int) ([]*PublicMessage, error)
	DeletePublicMessage(ctx context.Context, messageID uint) error
	LikePublicMessage(ctx context.Context, messageID uint, userID uint) error
	UnlikePublicMessage(ctx context.Context, messageID uint, userID uint) error
	CreateComment(ctx context.Context, comment *Comment) error
	GetCommentsByMessageID(ctx context.Context, messageID uint) ([]*Comment, error)
}

type PublicMessageService interface {
	CreatePublicMessage(ctx context.Context, message *PublicMessage) error
	GetPublicMessageByID(ctx context.Context, messageID uint) (*PublicMessage, error)
	GetPublicMessages(ctx context.Context, limit, offset int) ([]*PublicMessage, error)
	DeletePublicMessage(ctx context.Context, messageID uint) error
	LikePublicMessage(ctx context.Context, messageID uint, userID uint) error
	UnlikePublicMessage(ctx context.Context, messageID uint, userID uint) error
	CreateComment(ctx context.Context, comment *Comment) error
	GetCommentsByMessageID(ctx context.Context, messageID uint) ([]*Comment, error)
}
