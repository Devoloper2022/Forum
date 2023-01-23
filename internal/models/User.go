package models

type User struct {
	ID       int64  `json:"-"`
	Username string `json:"Username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
