package middleware

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/auth"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/errors"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/server"
)

func AuthMiddleware(header string, auth auth.Auth) server.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, ok := server.ExtractHeaderBearerToken(r, header)
			if !ok {
				server.RespondError(w, errors.NewAppUnauthorizedError())
				return
			}

			isValid, err := auth.ValidateToken(r.Context(), token)
			if err != nil || !isValid {
				server.RespondError(w, errors.NewAppUnauthorizedError())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
