package infra

import (
	"github.com/google/wire"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/pkg/infra/uow"
)

var ProviderSet = wire.NewSet(
	data.ProviderSet,
	uow.ProviderSet,
)
