//go:build integration
// +build integration

package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
)

func TestUpdateTaskDoneHandler(t *testing.T) {
	t.Parallel()

	user, token := GenerateUserAndToken()

	t.Run("it should return 204 and update task to done", func(t *testing.T) {
		ctx := context.Background()

		task, err := app.UseCase().CreateTaskUseCase.Execute(ctx, usecase.CreateTaskInput{
			Name:         "Task to update done",
			Description:  "Description",
			CategoryName: "Category update done",
			UserID:       user.ID,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/v1/tasks/%s/done", httpTestUrl, task.UUID), nil)
		assert.NoError(t, err)
		req.Header.Add(constant.Authorization, fmt.Sprintf("bearer %s", token))

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)

		taskUpdated, err := app.UseCase().FindOneTaskUseCase.Execute(ctx, task.UUID)
		assert.NoError(t, err)

		assert.True(t, taskUpdated.Done)
	})

}
