package service

import (
	"e-commerce/config"
	"e-commerce/internal/app/storage"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type Service struct {
	Storage storage.Storage
	Config  config.Config
}

func validate(errs v.Errors) error {
	err := errs.Filter()
	if err != nil {
		return e(NewValidationError(err))
	}
	return nil
}

func validateReadRequest(r *storage.ReadRequest, searchFields []string) error {
	if r == nil {
		return NewValidationError(errors.New("нет нужных полей в запросе"))
	}
	err := r.Validate(searchFields)
	if err != nil {
		return e(NewValidationError(err))
	}
	return nil
}

var (
	e  = errors.WithStack
	ef = errors.Wrapf
)

func IsValidationError(err error) (e ValidationError, ok bool) {
	ok = errors.As(err, &e)
	return
}
