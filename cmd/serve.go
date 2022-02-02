package cmd

import (
	user_service "github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user/service"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/handler"
	user_repo "github.com/TLSDevv/golang_catatan_keuangan_backend/repository/user"
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

	userRepo := user_repo.NewUserRepository()
	userService := user_service.NewUserService(userRepo, db)

	api := handler.NewAPI(userService)
	api.Start(conf.Host, conf.Port)
}
