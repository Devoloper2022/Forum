package dto

type UserDto struct {
	ID       int64  `json:"-"`
	Username string `json:"Username"`
}
