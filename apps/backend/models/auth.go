package models

import (
	"context"
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthCredentials struct {
	Email       string `json:"email,omitempty" validate:"omitempty,email"`
	Username    string `json:"username,omitempty" validate:"omitempty"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number,omitempty"`
	TwoFACode   string `json:"two_fa_code,omitempty"`
}

// user registration and retrieval
type AuthRepository interface {
	RegisterUser(ctx context.Context, registerData *AuthCredentials) (*User, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*User, error)
	VerifyPhoneNumber(ctx context.Context, userID uint, code string) error
	VerifyEmail(ctx context.Context, userID uint) error
}

// for login, registration, and verification
type AuthService interface {
	Login(ctx context.Context, loginData *AuthCredentials) (map[string]string, *User, error)
	Register(ctx context.Context, registerData *AuthCredentials) (string, *User, error)
	EnableTwoFactor(ctx context.Context, userID uint) error
	VerifyPhoneNumber(ctx context.Context, userID uint, code string) error
	VerifyEmail(ctx context.Context, userID uint) error
	IsValidTwoFACode(twoFACode string) bool
	Logout(ctx context.Context, refreshToken string) error
	RotateRefreshToken(ctx context.Context, oldRefreshToken string) (map[string]string, error)
	BlacklistAccessToken(ctx context.Context, accessToken string, expiry time.Duration) error
}

// Check if a password matches a hash
func MatchesHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Checks if an email is valid
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
