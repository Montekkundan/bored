package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"github.com/montekkundan/bored/backend/config"
	"github.com/montekkundan/bored/backend/models"
	"github.com/montekkundan/bored/backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repository       models.AuthRepository
	userService      models.UserService
	refreshTokenRepo models.RefreshTokenRepository
	redisClient      *redis.Client
	config           *config.EnvConfig
}

func (s *AuthService) Login(ctx context.Context, loginData *models.AuthCredentials) (map[string]string, *models.User, error) {
	var user *models.User
	var err error

	if loginData.Username != "" {
		user, err = s.userService.GetUserByUsername(ctx, loginData.Username)
	} else if loginData.Email != "" {
		user, err = s.userService.GetUserByEmail(ctx, loginData.Email)
	} else {
		return nil, nil, errors.New("please provide either email or username")
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, errors.New("invalid credentials")
		}
		return nil, nil, err
	}

	if !models.MatchesHash(loginData.Password, user.PasswordHash) {
		return nil, nil, errors.New("invalid credentials")
	}

	if !user.EmailVerified {
		return nil, nil, errors.New("email not verified")
	}

	tokenVersion := user.TokenVersion

	accessToken, err := utils.GenerateAccessToken(user.ID, user.Roles, tokenVersion, s.config.AccessTokenSecret, s.config.AccessTokenExpiry)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, tokenVersion, s.config.RefreshTokenSecret, s.config.RefreshTokenExpiry)
	if err != nil {
		return nil, nil, err
	}

	// Store Refresh Token in the DB
	expiresAt := time.Now().Add(time.Hour * 24 * time.Duration(s.config.RefreshTokenExpiry))
	err = s.refreshTokenRepo.Create(ctx, &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return nil, nil, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, user, nil
}

func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {
	// Remove refresh token from database
	return s.refreshTokenRepo.Delete(ctx, refreshToken)
}

func (s *AuthService) RotateRefreshToken(ctx context.Context, oldRefreshToken string) (map[string]string, error) {
	// Verify the refresh token
	storedToken, err := s.refreshTokenRepo.FindByToken(ctx, oldRefreshToken)
	if err != nil || storedToken.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("invalid or expired refresh token")
	}

	user, err := s.userService.GetUserByID(ctx, storedToken.UserID)
	if err != nil {
		return nil, err
	}

	// Invalidate the old refresh token
	err = s.refreshTokenRepo.Delete(ctx, oldRefreshToken)
	if err != nil {
		return nil, err
	}

	// Generate new tokens
	tokenVersion := user.TokenVersion
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Roles, tokenVersion, s.config.AccessTokenSecret, s.config.AccessTokenExpiry)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, tokenVersion, s.config.RefreshTokenSecret, s.config.RefreshTokenExpiry)
	if err != nil {
		return nil, err
	}

	expiresAt := time.Now().Add(time.Hour * 24 * time.Duration(s.config.RefreshTokenExpiry))
	err = s.refreshTokenRepo.Create(ctx, &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

// Blacklist access tokens in Redis
func (s *AuthService) BlacklistAccessToken(ctx context.Context, accessToken string, expiry time.Duration) error {
	return s.redisClient.Set(ctx, accessToken, "blacklisted", expiry).Err()
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

func NewAuthService(
	repository models.AuthRepository,
	userService models.UserService,
	config config.EnvConfig,
	refreshTokenRepo models.RefreshTokenRepository,
	redisClient *redis.Client,
) models.AuthService {
	return &AuthService{
		repository:       repository,
		userService:      userService,
		config:           &config,
		refreshTokenRepo: refreshTokenRepo,
		redisClient:      redisClient,
	}
}
