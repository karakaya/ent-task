package utils

import "errors"

var (
	ErrTransactionExists = errors.New("transaction already exists")
	ErrMissingSourceType = errors.New("missing Source-Type header")
	ErrInvalidState      = errors.New("invalid state")

	ErrInvalidUserId     = errors.New("invalid userId")
	ErrInternalServerErr = errors.New("internal server error")

	ErrInvalidJsonBody    = errors.New("invalid json body")
	ErrInvalidContentType = errors.New("invalid content-type")

	ErrInvalidAmount                  = errors.New("invalid amount")
	ErrAccountBalanceCannotBeNegative = errors.New("transaction would result in negative balance")
)
