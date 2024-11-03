package users

import (
	"errors"
	"fmt"
)

// AlreadyInUseError - this error is returned when a value of a unique field already in use.
//
// For example a user with given phone number already exists.
type AlreadyInUseError struct {
	Field string
}

var _ error = &AlreadyInUseError{}

func (err *AlreadyInUseError) Error() string {
	return fmt.Sprintf("field %s is already in use", err.Field)
}

var (
	ErrNotFound = errors.New("object is not found")
	ErrInternal = errors.New("internal error has occurred")
)
