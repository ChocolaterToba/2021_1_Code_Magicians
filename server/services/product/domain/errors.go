package domain

import "errors"

var (
	TransactionBeginError      = errors.New("Could not begin transaction")
	TransactionCommitError     = errors.New("Could not commit transaction")
	ShopNotFoundError          = errors.New("Could not find shop")
	ProductNotFoundError       = errors.New("Could not find product")
	CartNotFoundError          = errors.New("Could not find cart")
	ProductNotFoundInCartError = errors.New("Could not find product in cart")
	CartEmptyError             = errors.New("Cart is empty")
	OrderNotFoundError         = errors.New("Could not find order")
	ForeignOrderError          = errors.New("This order belongs to another user")
	UnsupportedExtensionError  = errors.New("File extension not supported")
)
