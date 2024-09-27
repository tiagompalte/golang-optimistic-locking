package usecase

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

func TestFindUserUUIDExecute(t *testing.T) {
	t.Parallel()

	t.Run("should be find user by UUID", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		userRepository := data.NewUserRepository(datasql)

		us := FindUserUUIDUseCaseImpl{
			userRepository: userRepository,
		}

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_user u WHERE NOT u.deleted_at AND u.uuid = \?`,
		).WithArgs("uuid").WillReturnRows(sqlmock.NewRows([]string{
			"u.id", "u.created_at", "u.updated_at", "u.uuid", "u.name", "u.email",
		}).AddRow(1, time.Time{}, time.Time{}, "uuid", "User", "user@email.com"))

		task, err := us.Execute(ctx, "uuid")

		if err != nil {
			t.Error(err)
		}

		if task.ID != 1 {
			t.Errorf("task id should be 1 but is %d", task.ID)
		}

		if task.UUID != "uuid" {
			t.Errorf("task uuid should be uuid but is %s", task.UUID)
		}
	})
}
