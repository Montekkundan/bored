package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/montekkundan/bored/backend/models"
	"github.com/montekkundan/bored/backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repository  models.AuthRepository
	userService models.UserService
}

func (s *AuthService) Login(ctx context.Context, loginData *models.AuthCredentials) (string, *models.User, error) {
	var user *models.User
	var err error

	// Check if username or email is provided for login
	if loginData.Username != "" {
		user, err = s.userService.GetUserByUsername(ctx, loginData.Username)
	} else if loginData.Email != "" {
		user, err = s.userService.GetUserByEmail(ctx, loginData.Email)
	} else {
		return "", nil, fmt.Errorf("please provide either email or username")
	}

	// Handle user not found
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, fmt.Errorf("invalid credentials")
		}
		return "", nil, err
	}

	// Validate password
	if !models.MatchesHash(loginData.Password, user.PasswordHash) {
		return "", nil, fmt.Errorf("invalid credentials")
	}

	// Check email verification
	if !user.EmailVerified {
		return "", nil, fmt.Errorf("email not verified")
	}

	// Check 2FA if enabled
	if user.TwoFactorEnabled && !s.IsValidTwoFACode(loginData.TwoFACode) {
		return "", nil, fmt.Errorf("invalid or missing 2FA code")
	}

	// Generate JWT
	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Roles,
		"exp":  time.Now().Add(time.Hour * 168).Unix(),
	}

	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))

	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *AuthService) Register(ctx context.Context, registerData *models.AuthCredentials) (string, *models.User, error) {
	if !models.IsValidEmail(registerData.Email) {
		return "", nil, fmt.Errorf("please, provide a valid email to register")
	}

	if _, err := s.repository.GetUser(ctx, "email = ?", registerData.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil, fmt.Errorf("the user email is already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil, err
	}

	registerData.Password = string(hashedPassword)

	user, err := s.repository.RegisterUser(ctx, registerData)
	if err != nil {
		return "", nil, err
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Roles,
		"exp":  time.Now().Add(time.Hour * 168).Unix(),
	}

	// Generate the JWT
	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *AuthService) EnableTwoFactor(ctx context.Context, userID uint) error {
	return s.repository.VerifyPhoneNumber(ctx, userID, "123456") // TODO Replace with actual logic to verify phone number
}

func (s *AuthService) VerifyEmail(ctx context.Context, userID uint) error {
	return s.repository.VerifyEmail(ctx, userID)
}

func (s *AuthService) VerifyPhoneNumber(ctx context.Context, userID uint, code string) error {
	return s.repository.VerifyPhoneNumber(ctx, userID, code)
}

// IsValidTwoFACode checks if the provided 2FA code is valid
func (s *AuthService) IsValidTwoFACode(twoFACode string) bool {
	// Here we could verify the 2FA code, for example, with a service like Google Authenticator
	// For this example, we'll assume that "654321" is a valid code
	return twoFACode == "654321"
}

func NewAuthService(repository models.AuthRepository, userService models.UserService) models.AuthService {
	return &AuthService{
		repository:  repository,
		userService: userService,
	}
}
