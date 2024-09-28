package auth

import "github.com/tiagompalte/golang-clean-optimistic-locking/configs"

func ProviderSet(
	config configs.Config,
) Auth {
	return NewJwtAuth(config)
}
