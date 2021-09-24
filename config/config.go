package config

import (
	"fmt"
	"time"

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
		API:  newApi(),
		DB:   newDB(),
		AUTH: newAuth(),
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

func newAuth() Auth {
	tokenLifetimeMapping := viper.GetStringMapString("auth.jwt_lifetime")

	auth := Auth{
		JWTPrivateKey: viper.GetString("auth.jwt_private_key"),
		JWTPublicKey:  viper.GetString("auth.jwt_public_key"),
		JWTLifetime:   make(map[string]time.Duration, len(tokenLifetimeMapping)),
	}

	// fmt.Println(tokenLifetimeMapping)

	for issuer, durationStr := range tokenLifetimeMapping {
		duration, err := time.ParseDuration(durationStr)
		if err != nil {
			continue
		}
		auth.JWTLifetime[issuer] = duration
	}

	return auth
}
