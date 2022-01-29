package cmd

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/handler"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/repository"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/service"
)

var (
	conf *Config
)

func init() {
	loadConfig()
	conf = initConfig()
	dbconf := loadConfigDB()
	db = initDatabase(dbconf)
}

func Execute() {
	defer db.Close()
	//register all server needs, db,repo, etc

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo, db)

	api := handler.NewAPI(userService)
	api.Start(conf.Host, conf.Port)
}
