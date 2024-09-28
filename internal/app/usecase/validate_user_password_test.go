package usecase

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	internalErr "github.com/tiagompalte/golang-clean-optimistic-locking/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/crypto"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/errors"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/repository"
)

func TestValidateUserPasswordExecute(t *testing.T) {
	t.Parallel()
	crypto := crypto.NewCryptoMock()

	t.Run("Should be return user", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		userRepository := data.NewUserRepository(datasql)

		us := ValidateUserPasswordUseCaseImpl{
			userRepository: userRepository,
			crypto:         crypto,
		}

		mock.ExpectQuery(
			`SELECT pass_encrypted FROM tb_user WHERE NOT deleted_at AND email = \?`,
		).WithArgs("user@email.com").WillReturnRows(sqlmock.NewRows([]string{
			"pass_encrypted",
		}).AddRow("pass"))

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_user u WHERE NOT u.deleted_at AND u.email = \?`,
		).WithArgs("user@email.com").WillReturnRows(sqlmock.NewRows([]string{
			"u.id", "u.created_at", "u.updated_at", "u.uuid", "u.name", "u.email",
		}).AddRow(1, time.Time{}, time.Time{}, "uuid", "User", "user@email.com"))

		user, err := us.Execute(ctx, ValidateUserPasswordInput{
			Email:    "user@email.com",
			Password: "pass",
		})

		if err != nil {
			t.Error(err)
		}

		if user.ID != 1 {
			t.Errorf("task id should be 1 but is %d", user.ID)
		}

		if user.UUID != "uuid" {
			t.Errorf("task uuid should be uuid but is %s", user.UUID)
		}
	})

	t.Run("Should be return invalid login error if password is wrong", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		userRepository := data.NewUserRepository(datasql)

		us := ValidateUserPasswordUseCaseImpl{
			userRepository: userRepository,
			crypto:         crypto,
		}

		mock.ExpectQuery(
			`SELECT pass_encrypted FROM tb_user WHERE NOT deleted_at AND email = \?`,
		).WithArgs("user@email.com").WillReturnRows(sqlmock.NewRows([]string{
			"pass_encrypted",
		}).AddRow("pass"))

		_, err = us.Execute(ctx, ValidateUserPasswordInput{
			Email:    "user@email.com",
			Password: "wrong",
		})

		if !errors.IsAppError(err, internalErr.ErrorCodeInvalidLogin) {
			t.Error(err)
		}
	})
}
