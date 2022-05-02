package domain

import "errors"

var (
	TransactionBeginError  = errors.New("Could not begin transaction")
	TransactionCommitError = errors.New("Could not commit transaction")
	ShopNotFoundError      = errors.New("Could not find shop")
	ProductNotFoundError   = errors.New("Could not find product")
)
