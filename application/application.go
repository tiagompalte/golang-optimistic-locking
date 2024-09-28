package application

import (
	"github.com/tiagompalte/golang-clean-optimistic-locking/configs"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/auth"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/server"
)

type App struct {
	config  configs.Config
	server  server.Server
	useCase usecase.UseCase
	auth    auth.Auth
}

func ProvideApplication(
	config configs.Config,
	server server.Server,
	useCase usecase.UseCase,
	auth auth.Auth,
) App {
	return App{
		config,
		server,
		useCase,
		auth,
	}
}

func (app App) Config() configs.Config {
	return app.config
}

func (app App) Server() server.Server {
	return app.server
}

func (app App) UseCase() usecase.UseCase {
	return app.useCase
}

func (app App) Auth() auth.Auth {
	return app.auth
}
