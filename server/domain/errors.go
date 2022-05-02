package domain

import "errors"

var (
	ErrIncorrectPassword = errors.New("Incorrect username/password pair")
	ErrUserNotFound      = errors.New("User not found")
	ErrCookieNotFound    = errors.New("Cookie not found")
	ErrShopNotFound      = errors.New("Shop not found")
	ErrProductNotFound   = errors.New("Product not found")
)
