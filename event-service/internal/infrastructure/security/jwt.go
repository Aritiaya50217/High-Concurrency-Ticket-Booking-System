package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	secretKey string
	expire    time.Duration
}

func NewJWTService(secretKey string, expire time.Duration) *JWTService {
	return &JWTService{secretKey: secretKey, expire: expire}
}

func (j *JWTService) ValidateToken(tokenStr string) (uint, error) {

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(j.secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	// check if the token has expired.
	exp, ok := claims["exp"].(float64)
	if !ok {
		return 0, errors.New("token expiration mission.")
	}

	if time.Unix(int64(exp), 0).Before(time.Now()) {
		return 0, errors.New("token expired")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id not found")
	}

	return uint(userID), nil
}
