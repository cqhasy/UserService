package biz

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrTimeOut           = errors.New("verification code timed out")
	ErrCodeErrors        = errors.New("validation code errors")
	ErrConfirmPassword   = errors.New("passwords are inconsistent")
	ErrPassword          = errors.New("wrong password")
)
