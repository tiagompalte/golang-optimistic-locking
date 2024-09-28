package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/tiagompalte/golang-clean-optimistic-locking/configs"
)

func ProviderSet() configs.Config {
	viper := NewViperConfig()

	i := 0
	path := "./configs"
	for {
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			path = fmt.Sprintf("./.%s", path)
			i++
		} else {
			break
		}

		if i > 3 {
			panic("config path is not exist")
		}
	}

	configName := configs.ViperConfigurationName
	env := os.Getenv(Env)
	if env != "" {
		configName = fmt.Sprintf("%s_%s", configName, env)
	}

	cfg, err := viper.Load(configName, configs.ViperTomlConfigurationType, path)
	if err != nil {
		panic(err)
	}

	return cfg
}
