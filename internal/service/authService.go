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
	CreateUser(user models.User) error
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

func (s *AuthService) CreateUser(user models.User) error {
	// userCheck, err := s.repo.GetUser(user.Username)
	// if err != nil {
	// 	return err
	// }
	// if userCheck.Username == user.Username {
	// 	return ErrUserExist
	// }
	// if err := checkUser(user); err != nil {
	// 	return err
	// }
	// if user.Password, err = generatePasswordHash(user.Password); err != nil {
	// 	return err
	// }
	// return s.repo.CreateUser(user)
	return nil
}
