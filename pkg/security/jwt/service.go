package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("your-secret-key") // 🔒 Use env vars in production

type Service interface {
	GenerateToken(userID string, role string) (string, error)
	ValidateToken(token string) (userID string, role string, err error)
}

type DefaultService struct {
	secretKey []byte
}

func (s DefaultService) GenerateToken(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s DefaultService) ValidateToken(tokenString string) (userID string, role string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("cannot parse claims")
	}

	userID, ok1 := claims["sub"].(string)
	role, ok2 := claims["role"].(string)

	if !ok1 || !ok2 {
		return "", "", errors.New("invalid claims data")
	}

	return userID, role, nil
}

func NewJwtService(secret string) Service {
	return &DefaultService{
		secretKey: []byte(secret),
	}
}
