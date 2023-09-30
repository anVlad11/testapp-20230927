package config

import (
	"errors"
	"github.com/anvlad11/testapp-20230927/pkg/config"
	"github.com/spf13/viper"
)

func NewConfig(path string) (*config.App, error) {
	if path == "" {
		return nil, errors.New("config path is empty")
	}

	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg *config.App
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
