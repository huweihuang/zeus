package errors

import (
	go_errors "errors"
)

// Provide the canonical errors.New, so we don't need to import canonical errors pkg again
func New(text string) error {
	return go_errors.New(text)
}

var (
	// NotFound errors
	ErrJobNotFound      = New("job not found")
	ErrInstanceNotFound = New("instance not found")

	// Validation errors
	ErrInvalidName = New("invalid name value")

	// MySQL errors
	MySQLDuplicateEntry = New("duplicate entry")
)
