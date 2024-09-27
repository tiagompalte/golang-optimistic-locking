package usecase

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	pkgErrors "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

func TestDeleteTaskExecute(t *testing.T) {
	t.Parallel()

	t.Run("should delete task", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		taskRepository := data.NewTaskRepository(datasql)

		us := DeleteTaskUseCaseImpl{
			taskRepository: taskRepository,
		}

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_task t JOIN tb_category c ON NOT c.deleted_at AND t.category_id = c.id WHERE NOT t.deleted_at AND t.uuid = \?`,
		).WithArgs("uuid").WillReturnRows(sqlmock.NewRows([]string{
			"t.id", "t.created_at", "t.updated_at", "t.uuid", "t.name", "t.description", "t.done", "c.id", "c.created_at", "c.updated_at", "c.slug", "c.name", "t.user_id",
		}).AddRow(1, time.Time{}, time.Time{}, "uuid", "task", "new task", false, 1, time.Time{}, time.Time{}, "category", "category", 1))

		mock.ExpectExec(
			`UPDATE tb_task	SET deleted_at = NOW\(\) WHERE NOT deleted_at AND id = \?`,
		).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

		input := DeleteTaskUseCaseInput{
			UUID:   "uuid",
			UserID: 1,
		}
		err = us.Execute(ctx, input)

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("should return error if user is invalid", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		taskRepository := data.NewTaskRepository(datasql)

		us := DeleteTaskUseCaseImpl{
			taskRepository: taskRepository,
		}

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_task t JOIN tb_category c ON NOT c.deleted_at AND t.category_id = c.id WHERE NOT t.deleted_at AND t.uuid = \?`,
		).WithArgs("uuid").WillReturnRows(sqlmock.NewRows([]string{
			"t.id", "t.created_at", "t.updated_at", "t.uuid", "t.name", "t.description", "t.done", "c.id", "c.created_at", "c.updated_at", "c.slug", "c.name", "t.user_id",
		}).AddRow(1, time.Time{}, time.Time{}, "uuid", "task", "new task", false, 1, time.Time{}, time.Time{}, "category", "category", 2))

		input := DeleteTaskUseCaseInput{
			UUID:   "uuid",
			UserID: 1,
		}
		err = us.Execute(ctx, input)

		if err == nil || !errors.IsAppError(err, pkgErrors.ErrorCodeInvalidUser) {
			t.Error(err)
		}
	})

	t.Run("should return error if does not find task", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		taskRepository := data.NewTaskRepository(datasql)

		us := DeleteTaskUseCaseImpl{
			taskRepository: taskRepository,
		}

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_task t JOIN tb_category c ON NOT c.deleted_at AND t.category_id = c.id WHERE NOT t.deleted_at AND t.uuid = \?`,
		).WithArgs("uuid").WillReturnError(errors.NewAppNotFoundError("task"))

		input := DeleteTaskUseCaseInput{
			UUID:   "uuid",
			UserID: 1,
		}
		err = us.Execute(ctx, input)

		if err == nil || !errors.IsAppError(err, errors.ErrorCodeNotFound) {
			t.Error(err)
		}
	})
}
