package db

import (
	"github.com/montekkundan/bored/backend/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Event{},
		&models.Ticket{},
		&models.User{},
		&models.Notification{},
		&models.Chat{},
		&models.ChatMember{},
		&models.Message{},
		&models.ModerationVote{},
		&models.RefreshToken{},
		&models.PublicMessage{},
		&models.Comment{},
		&models.BoringSpace{},
		&models.BoringSpaceMember{},
	)
}
