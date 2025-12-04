package apperror

import (
	"github.com/cockroachdb/errors"
)

func New(format string, args ...interface{}) error {
	return errors.Newf(format, args...)
}

func WithMessage(err error, format string, args ...interface{}) error {
	return errors.WithMessagef(err, format, args...)
}

func Wrap(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return errors.Wrapf(err, format, args...)
}

type AppError interface {
	error
	Code() int
	Unwrap() error
}

type appError struct {
	cause   error
	code    int
	message string
}
