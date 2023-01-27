package service

import (
	"fmt"
	dto "forum/internal/DTO"
	"forum/internal/models"
	"forum/internal/repository"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User interface {
	CreateUser(user models.User) error
	Get(userId int64) (models.User, error)
	// UpdateUser(user models.User) error
	// DeleteUser(user models.User) error
}

type UserService struct {
	user repository.User
}

func NewUserService(user repository.User) *UserService {
	return &UserService{user: user}
}

func (s *UserService) CreateUser(user models.User) error {
	validEmail, err := regexp.MatchString(`[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`, user.Email)
	if err != nil {
		return err
	} else if !validEmail {
		return dto.ErrInvalidEmail
	}

	hash, err := generatePasswordHash(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash

	err = s.user.CreateUser(user)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.Email" {
			return dto.ErrEmailExist
		} else if err.Error() == "UNIQUE constraint failed: users.Username" {
			return dto.ErrUsernameExist
		}
	}
	return err
} // done

func (s *UserService) Get(userId int64) (models.User, error) {
	return s.user.GetUser(userId)
} // done

func generatePasswordHash(password string) (string, error) {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		return "", fmt.Errorf("service : generatePassword : %v", err)
	}
	return string(hash), nil
}

func checkHash(hpass, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hpass), []byte(password))
}
