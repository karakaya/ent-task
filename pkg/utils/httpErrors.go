package utils

import "errors"

var (
	ErrTransactionExists = errors.New("transaction already exists")
	ErrMissingSourceType = errors.New("missing Source-Type header")
	ErrInvalidState      = errors.New("invalid state")

	ErrInvalidUserId     = errors.New("invalid userId")
	ErrInvalidJsonBody   = errors.New("invalid json body")
	ErrInternalServerErr = errors.New("internal server error")

	ErrInvalidAmount = errors.New("invalid amount")
)
