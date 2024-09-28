package errors

import (
	"net/http"

	pkg "github.com/tiagompalte/golang-clean-optimistic-locking/pkg/errors"
)

// ErrorCodeEmptyPath means that path is empty
const ErrorCodeEmptyPath = "empty-path"

func NewEmptyPathError(field string) pkg.AppError {
	return pkg.AppError{
		StatusCode: http.StatusBadRequest,
		Code:       ErrorCodeEmptyPath,
		Field:      field,
	}
}
