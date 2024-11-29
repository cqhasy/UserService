package biz

import "errors"

var (
	ErrUserAlreadyExists = errors.New("the user already exists, please try a different email")
	ErrUserNotFound      = errors.New("the specified user was not found in the database")
	ErrInvalidPassword   = errors.New("the password entered is invalid, please try again")
	ErrTimeOut           = errors.New("the verification code has timed out, please request a new code")
	ErrCodeErrors        = errors.New("the provided verification code is incorrect, please try again")
	ErrConfirmPassword   = errors.New("the passwords provided do not match, please check and try again")
	ErrPassword          = errors.New("the current password entered is incorrect, please try again")
)
