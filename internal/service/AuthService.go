package service

import (
	dto "forum/internal/DTO"
	"forum/internal/models"
	"forum/internal/repository"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Autorization interface {
	GenerateToken(login dto.Credentials) (dto.Cook, error)
	ParseToken(token string) (models.User, error)
	DeleteToken(token string) error
	DeleteTokenWhenExpireTime() error
}

type AuthService struct {
	auth repository.Autorization
	user repository.User
}

func NewAuthService(auth repository.Autorization, user repository.User) *AuthService {
	return &AuthService{
		auth: auth,
		user: user,
	}
}

func (s *AuthService) GenerateToken(login dto.Credentials) (dto.Cook, error) {
	var user models.User
	validEmail, err := regexp.MatchString(mailValidation, login.Username)

	if err != nil {
		return dto.Cook{}, err
	} else if validEmail {
		user, err = s.user.GetUserByEmail(login.Username)

		if err != nil {
			return dto.Cook{}, dto.ErrUserNotFound
		}
	} else {
		user, err = s.user.GetUserByUsername(login.Username)

		if err != nil {
			return dto.Cook{}, dto.ErrUserNotFound
		}
	}

	if err := checkHash(user.Password, login.Password); err != nil {
		return dto.Cook{}, dto.ErrPasswdNotMatch
	}

	err = s.auth.GetTokens(user.ID)

	if err == nil {
		s.auth.DeleteTokenByUserID(user.ID)
	}

	cook := dto.Cook{
		Token:  uuid.NewString(),
		Expiry: time.Now().Add(15 * time.Minute),
	}

	if err := s.auth.SaveToken(user.ID, cook.Token, cook.Expiry); err != nil {
		return dto.Cook{}, err
	}
	return cook, nil
}

func (s *AuthService) ParseToken(token string) (models.User, error) {
	user, err := s.user.GetUserByToken(token)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *AuthService) DeleteToken(token string) error {
	return s.auth.DeleteToken(token)
}

func (s *AuthService) DeleteTokenWhenExpireTime() error {
	return s.auth.DeleteTokenWhenExpireTime()
}

func checkHash(hpass, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hpass), []byte(password))
}
