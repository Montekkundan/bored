package models

import (
	"context"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type UserRole string

const (
	Manager   UserRole = "manager"
	Attendee  UserRole = "attendee"
	Admin     UserRole = "admin"
	Viewer    UserRole = "viewer"
	Editor    UserRole = "editor"
	BoredUser UserRole = "bored_user"
	Moderator UserRole = "moderator"
)

type User struct {
	ID               uint                `json:"id" gorm:"primarykey"`
	Username         string              `json:"username" gorm:"text;not null;unique"`
	Email            string              `json:"email" gorm:"text;not null;unique"` // email will be like montekkundan@bored.rocks
	PasswordHash     string              `json:"-" gorm:"text;not null"`            // Do not expose password hash in JSON
	TokenVersion     int                 `json:"token_version" gorm:"default:1"`
	Bio              string              `json:"bio" gorm:"text"`
	Interests        []string            `json:"interests" gorm:"type:text[]"`
	Latitude         float64             `json:"latitude" gorm:"numeric(9,6)"`
	Longitude        float64             `json:"longitude" gorm:"numeric(9,6)"`
	ProfilePicture   string              `json:"profile_picture" gorm:"text"`
	CoverPhoto       string              `json:"cover_photo" gorm:"text"`
	SocialLinks      string              `json:"social_links" gorm:"jsonb"`                // JSONB to store GitHub, Google, etc.
	OAuthProviders   []OAuthProvider     `json:"oauth_providers" gorm:"foreignkey:UserID"` // Connected OAuth accounts
	AudioEnabled     bool                `json:"audio_enabled" gorm:"default:false"`
	VideoEnabled     bool                `json:"video_enabled" gorm:"default:false"`
	Roles            pq.StringArray      `json:"roles" gorm:"type:text[];default:ARRAY['bored_user']::text[]" swaggertype:"array,string"`
	EmailVerified    bool                `json:"email_verified" gorm:"default:false"`
	PhoneNumber      string              `json:"phone_number" gorm:"text;unique"` // Phone number for 2FA
	PhoneVerified    bool                `json:"phone_verified" gorm:"default:false"`
	TwoFactorEnabled bool                `json:"two_factor_enabled" gorm:"default:false"` // 2FA enabled
	RewardPoints     int                 `json:"reward_points" gorm:"default:0"`
	Followers        []User              `gorm:"many2many:user_followers"`               // Followers relationship
	Following        []User              `gorm:"many2many:user_following"`               // Following relationship
	Notifications    []Notification      `json:"notifications" gorm:"foreignkey:UserID"` // User notifications
	CreatedAt        time.Time           `json:"created_at" gorm:"default:now()"`
	UpdatedAt        time.Time           `json:"updated_at" gorm:"default:now()"`
	Chats            []Chat              `json:"chats" gorm:"many2many:user_chats"`         // Direct and group chats
	ModerationVotes  []ModerationVote    `json:"moderation_votes" gorm:"foreignkey:UserID"` // Votes for content moderation
	Deactivated      bool                `json:"deactivated" gorm:"default:false"`
	PublicMessages   []PublicMessage     `json:"public_messages" gorm:"foreignkey:UserID"`
	BoringSpaces     []BoringSpaceMember `json:"boring_spaces" gorm:"foreignKey:UserID"`
}

type UserService interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserByID(ctx context.Context, userID uint) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUserByID(ctx context.Context, userID uint) error
	DeactivateUser(ctx context.Context, userID uint) error
	GetUserBoringSpaces(ctx context.Context, userID uint) ([]*BoringSpaceMember, error)
	GetAllPublicMessages(ctx context.Context, limit, offset int) ([]*PublicMessage, error)
}

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserByID(ctx context.Context, userID uint) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUserByID(ctx context.Context, userID uint) error
	DeactivateUser(ctx context.Context, userID uint) error
	GetUserBoringSpaces(ctx context.Context, userID uint) ([]*BoringSpaceMember, error)
	GetAllPublicMessages(ctx context.Context, limit, offset int) ([]*PublicMessage, error)
}

// AfterCreate hook to assign the admin role to the first user in the database
func (u *User) AfterCreate(db *gorm.DB) (err error) {
	if u.ID == 1 {
		db.Model(u).Update("roles", pq.StringArray{string(Admin)})
	}
	return
}

func (u *User) HasRole(role UserRole) bool {
	for _, r := range u.Roles {
		if UserRole(r) == role {
			return true
		}
	}
	return false
}

func (u *User) HasBoringSpaceRole(spaceID uint, role BoringSpaceRole) bool {
	for _, member := range u.BoringSpaces {
		if member.BoringSpaceID == spaceID && member.Role == role {
			return true
		}
	}
	return false
}
