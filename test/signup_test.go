//go:build integration
// +build integration

package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/pkg/server/handler"
)

func TestSignupHandler(t *testing.T) {
	t.Parallel()

	t.Run("it should create new user and return 200 and token", func(t *testing.T) {
		t.Parallel()

		signupRequest := handler.SignupRequest{
			Name:     RandomName(),
			Email:    Email(),
			Password: "Pass!1234",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(signupRequest)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/signup", httpTestUrl), &buf)
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("it should return 401 if informed invalid payload", func(t *testing.T) {
		t.Parallel()

		signupRequest := handler.SignupRequest{
			Name:     "",
			Email:    "invalid_email",
			Password: "Pass!1234",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(signupRequest)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/signup", httpTestUrl), &buf)
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
	})
}
