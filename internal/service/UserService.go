package service

import (
	"forum/internal/models"
	"forum/internal/repository"
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
	return s.user.CreateUser(user)
} // done

func (s *UserService) Get(userId int64) (models.User, error) {
	return s.user.GetUser(userId)
} // done
