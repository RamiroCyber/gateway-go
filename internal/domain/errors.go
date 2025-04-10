package domain

import "errors"

var (
	ErrAccountNotFound     = errors.New("account not found")
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrUnauthorized        = errors.New("unauthorized access")
	ErrAPIKeyExists        = errors.New("api key already exists")
	ErrInvoiceNotFound     = errors.New("invoice not found")
)
