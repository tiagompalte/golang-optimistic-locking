package server

import (
	"github.com/tiagompalte/golang-clean-optimistic-locking/configs"
)

func ProviderSet(
	config configs.Config,
) Server {
	return NewGoChiServer(config)
}
