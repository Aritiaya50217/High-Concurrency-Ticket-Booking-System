package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey string
	expire    time.Duration
}

func NewJWTService(secretKey string, expire time.Duration) *JWTService {
	return &JWTService{secretKey: secretKey, expire: expire}
}

func (j *JWTService) GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(j.expire).Unix(),
	})

	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) ValidateToken(tokenStr string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	cliams, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userID, ok := cliams["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id not found in token")
	}

	return uint(userID), nil
}
