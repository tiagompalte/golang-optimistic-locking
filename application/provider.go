package application

import (
	"github.com/google/wire"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/pkg/infra"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg"
)

var ProviderSet = wire.NewSet(
	pkg.ProviderSet,
	infra.ProviderSet,
	usecase.ProviderSet,
	wire.Struct(new(usecase.UseCase), "*"),
	ProvideApplication,
)
