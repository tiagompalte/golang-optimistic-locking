package repository

import (
	"github.com/tiagompalte/golang-clean-optimistic-locking/configs"
)

func ProviderSet(
	config configs.Config,
) DataManager {
	return NewDataSqlWithConfig(config.Database)
}
