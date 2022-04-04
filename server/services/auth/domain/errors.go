package domain

import "errors"

var (
	TransactionBeginError  = errors.New("Could not begin transaction")
	TransactionCommitError = errors.New("Could not commit transaction")
	UserNotFoundError      = errors.New("Could not find user")
	IncorrectPasswordError = errors.New("incorrect password")
)
