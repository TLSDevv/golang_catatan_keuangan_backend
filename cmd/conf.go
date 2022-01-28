package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port string
}

type Database struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func loadConfig() {
	viper.SetConfigName(".env") // name of config file (without extension)
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

func loadConfigDB() *Database {
	driver := viper.GetString("DB_DRIVER")
	if len(driver) == 0 {
		driver = "mysql"
	}

	host := viper.GetString("DB_HOST")
	if len(host) == 0 {
		host = "localhost"
	}

	port := viper.GetString("DB_PORT")
	if len(port) == 0 {
		port = "5000"
	}

	username := viper.GetString("DB_USERNAME")
	if len(username) == 0 {
		username = "root"
	}

	password := viper.GetString("DB_PASSWORD")
	if len(password) == 0 {
		password = ""
	}

	name := viper.GetString("DB_NAME")
	if len(name) == 0 {
		name = "keuanganku"
	}

	return &Database{
		Driver:   driver,
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Name:     name,
	}
}
