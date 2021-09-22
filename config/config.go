package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func New() (*Config, error) {
	viper.SetConfigFile("config.yml")

	err := viper.ReadInConfig()

	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)

		if ok {
			return nil, fmt.Errorf("config file not found: %w", err)
		} else {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}

	}
	viper.AutomaticEnv()

	config := &Config{
		API: newApi(),
		DB:  newDB(),
	}

	return config, nil
}

func newApi() API {
	return API{
		Host:    viper.GetString("api.host"),
		Port:    viper.GetInt("api.port"),
		BaseUrl: viper.GetString("api.base_url"),
	}
}

func newDB() DB {
	return DB{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Name:     viper.GetString("db.name"),
	}
}
