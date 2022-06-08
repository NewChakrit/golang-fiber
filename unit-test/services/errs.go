package services

import "errors"

var (
	ErrZeroAmount = errors.New("Purchase amount cound not be zero")
	ErrRepository = errors.New("Repository error")
)
