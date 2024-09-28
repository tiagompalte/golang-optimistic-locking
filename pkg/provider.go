package pkg

import (
	"github.com/google/wire"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/auth"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/cache"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/config"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/crypto"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/repository"
	"github.com/tiagompalte/golang-clean-optimistic-locking/pkg/server"
)

var ProviderSet = wire.NewSet(
	config.ProviderSet,
	cache.ProviderSet,
	wire.Bind(new(repository.Connector), new(repository.DataManager)),
	repository.ProviderSet,
	server.ProviderSet,
	crypto.ProviderSet,
	auth.ProviderSet,
)
