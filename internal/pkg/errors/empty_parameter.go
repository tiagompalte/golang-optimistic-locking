package errors

import (
	"net/http"

	pkg "github.com/tiagompalte/golang-clean-optimistic-locking/pkg/errors"
)

// ErrorCodeEmptyParameter means that parameter is empty
const ErrorCodeEmptyParameter = "empty-parameter"

func NewEmptyParameterError(field string) pkg.AppError {
	return pkg.AppError{
		StatusCode: http.StatusBadRequest,
		Code:       ErrorCodeEmptyParameter,
		Field:      field,
	}
}
