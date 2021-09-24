package config

import "time"

type API struct {
	Host    string
	Port    int
	BaseUrl string
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Config struct {
	API  API
	DB   DB
	AUTH Auth
}
type Auth struct {
	JWTPrivateKey string
	JWTPublicKey  string
	JWTLifetime   map[string]time.Duration
}
