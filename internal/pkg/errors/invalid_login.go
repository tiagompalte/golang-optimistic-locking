package errors

import (
	"net/http"

	pkg "github.com/tiagompalte/golang-clean-optimistic-locking/pkg/errors"
)

// ErrorCodeInvalidLogin means that path is empty
const ErrorCodeInvalidLogin = "invalid-login"

func NewInvalidLoginError() pkg.AppError {
	return pkg.AppError{
		StatusCode: http.StatusUnauthorized,
		Code:       ErrorCodeInvalidLogin,
	}
}
