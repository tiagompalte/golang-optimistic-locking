package crypto

import "github.com/tiagompalte/golang-clean-optimistic-locking/configs"

func ProviderSet(
	config configs.Config,
) Crypto {
	return NewBcrypt(config.Bcrypt)
}
