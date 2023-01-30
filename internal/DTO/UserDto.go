package dto

import (
	"forum/internal/models"
	"time"
)

type UserDto struct {
	ID       int64  `json:"-"`
	Username string `json:"Username"`
}

func (dto *UserDto) GetUserModel() models.User {
	return models.User{
		ID:       dto.ID,
		Username: dto.Username,
	}
}

func GetUserDto(u models.User) UserDto {
	return UserDto{
		ID:       u.ID,
		Username: u.Username,
	}
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Cook struct {
	Token  string    `json:"Token"`
	Expiry time.Time `json:"Expiry"`
}

// type ProfileDto struct {
// 	ID       int64  `json:"-"`
// 	Username string `json:"Username"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// func GetProfileDto() {
// }

// func (p *ProfileDto) GetProfileModel() {
// }
