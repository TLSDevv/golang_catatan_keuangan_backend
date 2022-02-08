package cmd

import (
	"context"
	"os"
	"os/signal"

	user_service "github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user/service"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/handler"
	user_repo "github.com/TLSDevv/golang_catatan_keuangan_backend/repository/user"
	"github.com/sirupsen/logrus"
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
	ctx := context.Background()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	logrus.SetFormatter(&logrus.JSONFormatter{})

	go func() {
		oscall := <-c
		logrus.Info("system call: %v", oscall)
		cancel()
	}()

	defer db.Close()
	//register all server needs, db,repo, etc

	userRepo := user_repo.NewUserRepository()
	userService := user_service.NewUserService(userRepo, db)

	api := handler.NewAPI(userService)
	api.Start(ctx, conf.Host, conf.Port)
}
