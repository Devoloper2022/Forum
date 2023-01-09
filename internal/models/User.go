package models

type User struct {
	ID       int64  `json:"-"`
	Name     string `json:"name"`
	Username string `json:"Username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
