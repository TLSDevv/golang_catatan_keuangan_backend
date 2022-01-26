package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port string
}

func loadConfig() {
	viper.SetConfigName(".env")  // name of config file (without extension)
	viper.SetConfigType("yml")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file

	if err != nil {
		logrus.Error("error load config ", err)
		return
	}

	return
}

func initConfig() *Config {

	host := viper.GetString("HOST")
	if len(host) == 0 {
		host = "localhost"
	}

	port := viper.GetString("PORT")
	if len(port) == 0 {
		port = "5000"
	}

	return &Config{
		Host: host,
		Port: port,
	}
}
