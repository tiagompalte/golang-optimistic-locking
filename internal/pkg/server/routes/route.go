package routes

import (
	"github.com/tiagompalte/golang-clean-optimistic-locking/application"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/server"
)

func CreateRoute(app application.App) []server.GroupRoute {
	return []server.GroupRoute{
		{
			Path: "/api",
			GroupRoutes: []server.GroupRoute{
				CreateRouteV1(app),
			},
			Routes: []server.Route{
				{
					Method:  "GET",
					Path:    "/health-check",
					Handler: handler.HealthCheckHandler(app.UseCase().HealthCheckUseCase),
				},
			},
		},
	}
}
