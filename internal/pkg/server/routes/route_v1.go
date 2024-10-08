package routes

import (
	"github.com/tiagompalte/golang-clean-optimistic-locking/application"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/server"
)

func CreateRouteV1(app application.App) server.GroupRoute {
	return server.GroupRoute{
		Path: "/v1",
		GroupRoutes: []server.GroupRoute{
			CreateGroupCurrentUserV1(app),
		},
		Routes: []server.Route{
			{
				Method:  "POST",
				Path:    "/signup",
				Handler: handler.SignupHandler(app.UseCase().CreateUserUseCase, app.UseCase().GenerateUserTokenUseCase),
			},
			{
				Method:  "POST",
				Path:    "/signin",
				Handler: handler.SigninHandler(app.UseCase().ValidateUserPasswordUseCase, app.UseCase().GenerateUserTokenUseCase),
			},
		},
	}
}
