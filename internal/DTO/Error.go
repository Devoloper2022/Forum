package dto

import "errors"

type SystemErr struct {
	Status int
	Msg    string
}

var (
	ErrInvalidEmail    = errors.New("Invalid email")
	ErrInvalidUsername = errors.New("Invalid username")
	ErrInvalidPassword = errors.New("Invalid password")
	ErrUserNotFound    = errors.New("User not found")
	ErrUserExist       = errors.New("User exist")
	ErrUsernameExist   = errors.New("Username exist")
	ErrEmailExist      = errors.New("Email exist")
	ErrPasswdNotMatch  = errors.New("Password doesn't match")
)
