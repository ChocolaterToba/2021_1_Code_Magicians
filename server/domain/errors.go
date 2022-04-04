package domain

import "errors"

var (
	ErrIncorrectPassword = errors.New("Incorrect username/password pair")
	ErrUserNotFound      = errors.New("User not found")
	ErrCookieNotFound    = errors.New("Cookie not found")
)
