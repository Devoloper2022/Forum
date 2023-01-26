package dto

import "forum/internal/models"

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
