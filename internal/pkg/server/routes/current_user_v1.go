package routes

import (
	"github.com/tiagompalte/golang-clean-optimistic-locking/application"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/pkg/server/middleware"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/server"
)

func CreateGroupCurrentUserV1(app application.App) server.GroupRoute {
	routes := []server.Route{
		{
			Path:    "/",
			Method:  "GET",
			Handler: handler.FindUserLoggedHandler(),
		},
	}

	return server.GroupRoute{
		Path: "/current/user",
		Middlewares: []server.Middleware{
			middleware.ValidateExtractUserTokenMiddleware(constant.Authorization, app.Auth(), app.UseCase().FindUserUUIDUseCase),
		},
		Routes: routes,
	}
}
