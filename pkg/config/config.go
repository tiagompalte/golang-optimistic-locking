package config

import "github.com/tiagompalte/golang-clean-optimistic-locking/configs"

type Config interface {
	Load(filename string, configType string, path string) (configs.Config, error)
}
