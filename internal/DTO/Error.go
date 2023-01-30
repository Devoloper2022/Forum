package dto

import (
	"errors"
)

type SystemErr struct {
	Status int
	Name   string
	Msg    string
}

var (
	// user error
	ErrUserNotFound = errors.New("User not found")
	ErrUserExist    = errors.New("User exist")

	// email error
	ErrEmailInvalid = errors.New("Invalid email")
	ErrEmailExist   = errors.New("Email exist")

	// username error
	ErrUsernameExist   = errors.New("Username exist")
	ErrUsernameInvalid = errors.New("Invalid username")

	// password error
	ErrPasswordInvalid = errors.New("Invalid password")
	ErrPasswdNotMatch  = errors.New("Password doesn't match")
)
