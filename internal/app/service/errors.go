package service

import "github.com/pkg/errors"

type ValidationError struct {
	Err error
}

func NewValidationError(err error) error {
	return ValidationError{Err: err}
}

func (e ValidationError) Error() string {
	return "validation: " + e.Err.Error()
}

type UnauthenticatedError struct {
	Err error
}

func IsUnauthenticatedError(err error) (e UnauthenticatedError, ok bool) {
	ok = errors.As(err, &e)
	return e, ok
}
