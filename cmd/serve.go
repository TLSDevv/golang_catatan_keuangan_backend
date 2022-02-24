package cmd

import (
	"context"
	"os"
	"os/signal"

	auth_service "github.com/TLSDevv/golang_catatan_keuangan_backend/domain/auth/service"
	transaction_service "github.com/TLSDevv/golang_catatan_keuangan_backend/domain/transaction/service"
	user_service "github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user/service"
	auth_repo "github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/db/mysql/auth"
	transaction_repo "github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/db/mysql/transaction"
	user_repo "github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/db/mysql/user"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http"
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
	userRepo := user_repo.NewUserRepository(db)
	userService := user_service.NewUserService(userRepo)
	authRepo := auth_repo.NewAuthRepository(db)
	authService := auth_service.NewAuthService(authRepo, userRepo)
	tr := transaction_repo.NewTransactionRepository(db)
	ts := transaction_service.NewTransactionService(tr, userRepo)

	api := http.NewAPI(userService, authService, ts)
	api.Start(ctx, conf.Host, conf.Port)
}
