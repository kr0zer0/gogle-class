package auth

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenManager interface {
	GenerateToken(userId string, ttl time.Duration) (string, error)
	ParseToken(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) *Manager {
	return &Manager{signingKey: signingKey}
}

func (m *Manager) GenerateToken(userId string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userId,
		"ExpiresAt": time.Now().Add(ttl).Unix(),
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("error when getting user claims from token")
	}

	return claims["userID"].(string), nil
}

func (m *Manager) NewRefreshToken() (string, error) {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", token), nil
}
