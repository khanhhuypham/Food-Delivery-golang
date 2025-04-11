package user_model

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmailNotFound      = errors.New("email not found")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrInvalidEmail       = errors.New("invalid email")
	ErrInvalidPhone       = errors.New("invalid phone")
	ErrInvalidAvatar      = errors.New("invalid avatar")
	ErrInvalidFirstName   = errors.New("invalid first name")
	ErrInvalidLastName    = errors.New("invalid last name")

	ErrUserBannedOrDeleted     = errors.New("user banned or deleted")
	ErrInvalidEmailAndPassword = errors.New("invalid email or password")
)
