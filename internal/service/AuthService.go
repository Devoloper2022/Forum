package service

import (
	"errors"
	"forum/internal/models"
	"forum/internal/repository"
)

var (
	ErrInvalidEmail    = errors.New("Invalid email")
	ErrInvalidUsername = errors.New("Invalid username")
	ErrInvalidPassword = errors.New("Invalid password")
	ErrUserNotFound    = errors.New("User not found")
	ErrUserExist       = errors.New("User exist")
)

type Autorization interface {
	CreateAuth(user models.User) error
	// GenerateToken(username, password string) (string, time.Time, error)
	// ParseToken(token string) (models.User, error)
	// DeleteToken(token string) error
}

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateAuth(user models.User) error {
	return nil
}
