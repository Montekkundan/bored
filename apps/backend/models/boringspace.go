package models

import (
	"context"
	"time"
)

type BoringSpaceRole string

const (
	BSAdmin     BoringSpaceRole = "admin"
	BSMember    BoringSpaceRole = "member"
	BSViewer    BoringSpaceRole = "viewer"
	BSModerator BoringSpaceRole = "moderator"
)

type BoringSpace struct {
	ID          uint                `json:"id" gorm:"primarykey"`
	Name        string              `json:"name" gorm:"text;not null;unique"`
	Description string              `json:"description" gorm:"text"`
	CreatorID   uint                `json:"creator_id" gorm:"not null"`
	Creator     User                `json:"creator" gorm:"foreignkey:CreatorID"`
	Members     []BoringSpaceMember `json:"members" gorm:"foreignKey:BoringSpaceID"`
	CreatedAt   time.Time           `json:"created_at" gorm:"default:now()"`
	UpdatedAt   time.Time           `json:"updated_at" gorm:"default:now()"`
}

type BoringSpaceMember struct {
	BoringSpaceID uint            `json:"boringspace_id" gorm:"primaryKey"`
	UserID        uint            `json:"user_id" gorm:"primaryKey"`
	User          User            `json:"user" gorm:"foreignKey:UserID"`
	BoringSpace   BoringSpace     `json:"boringspace" gorm:"foreignKey:BoringSpaceID"`
	Role          BoringSpaceRole `json:"role" gorm:"type:text;not null;default:'member'"`
	JoinedAt      time.Time       `json:"joined_at" gorm:"default:now()"`
}

type BoringSpaceRepository interface {
	CreateBoringSpace(ctx context.Context, space *BoringSpace) error
	GetBoringSpaceByID(ctx context.Context, spaceID uint) (*BoringSpace, error)
	AddMember(ctx context.Context, member *BoringSpaceMember) error
	RemoveMember(ctx context.Context, spaceID uint, userID uint) error
	GetMembers(ctx context.Context, spaceID uint) ([]*BoringSpaceMember, error)
	UpdateMemberRole(ctx context.Context, spaceID uint, userID uint, role BoringSpaceRole) error
}

type BoringSpaceService interface {
	CreateBoringSpace(ctx context.Context, space *BoringSpace) error
	GetBoringSpaceByID(ctx context.Context, spaceID uint) (*BoringSpace, error)
	AddMember(ctx context.Context, member *BoringSpaceMember) error
	RemoveMember(ctx context.Context, spaceID uint, userID uint) error
	GetMembers(ctx context.Context, spaceID uint) ([]*BoringSpaceMember, error)
	UpdateMemberRole(ctx context.Context, spaceID uint, userID uint, role BoringSpaceRole) error
}
