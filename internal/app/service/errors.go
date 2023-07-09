package service

type ValidationError struct {
	Err error
}

func NewValidationError(err error) error {
	return ValidationError{Err: err}
}

func (e ValidationError) Error() string {
	return "validation: " + e.Err.Error()
}
