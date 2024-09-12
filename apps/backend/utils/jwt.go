package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(claims jwt.Claims, method jwt.SigningMethod, jwtSecret string) (string, error) {
	return jwt.NewWithClaims(method, claims).SignedString([]byte(jwtSecret))
}

func GenerateAccessToken(userID uint, roles []string, tokenVersion int, jwtSecret string, expiryMinutes int) (string, error) {
	claims := jwt.MapClaims{
		"id":      userID,
		"roles":   roles,
		"version": tokenVersion,
		"exp":     time.Now().Add(time.Minute * time.Duration(expiryMinutes)).Unix(),
	}
	return GenerateJWT(claims, jwt.SigningMethodHS256, jwtSecret)
}

func GenerateRefreshToken(userID uint, tokenVersion int, jwtSecret string, expiryDays int) (string, error) {
	claims := jwt.MapClaims{
		"id":      userID,
		"version": tokenVersion,
		"exp":     time.Now().Add(time.Hour * 24 * time.Duration(expiryDays)).Unix(),
	}
	return GenerateJWT(claims, jwt.SigningMethodHS256, jwtSecret)
}

func ParseToken(tokenStr string, jwtSecret string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
}
