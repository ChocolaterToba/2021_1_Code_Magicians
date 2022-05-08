package domain

import "errors"

var (
	ErrIncorrectPassword     = errors.New("Incorrect username/password pair")
	ErrUserNotFound          = errors.New("User not found")
	ErrCookieNotFound        = errors.New("Cookie not found")
	ErrShopNotFound          = errors.New("Shop not found")
	ErrProductNotFound       = errors.New("Product not found")
	ErrCartNotFound          = errors.New("Cart not found")
	ErrProductNotFoundInCart = errors.New("Product not found in cart")
	ErrCartEmpty             = errors.New("Cart is empty")
)
