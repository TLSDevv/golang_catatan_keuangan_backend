package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/auth"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/transaction"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user"
	auth_handler "github.com/TLSDevv/golang_catatan_keuangan_backend/handler/auth"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/handler/middleware"
	transaction_handler "github.com/TLSDevv/golang_catatan_keuangan_backend/handler/transaction"
	user_handler "github.com/TLSDevv/golang_catatan_keuangan_backend/handler/user"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	userService        user.Service
	authService        auth.Service
	transactionService transaction.Service
}

func NewAPI(
	usService user.Service,
	auService auth.Service,
	tService transaction.Service,
) *API {
	return &API{
		userService:        usService,
		authService:        auService,
		transactionService: tService,
	}
}

func (a API) Start(ctx context.Context, host, port string) (err error) {
	r := mux.NewRouter()

	r.Use(middleware.CORS)

	route := r.PathPrefix("/api/v1").Subrouter()

	route.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("works"))
	}).Methods("GET")

	user_handler.NewUserHandler(route, a.userService)
	auth_handler.NewAuthHandler(route, a.authService)
	transaction_handler.NewTransactionHandler(route, a.transactionService)

	server := http.Server{
		Addr:    endPoint(host, port),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal("Listening %s to port %s", host, port)
		}
	}()

	logrus.Info("Server Started")

	<-ctx.Done()

	logrus.Info("Server Stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctxShutDown); err != nil {
		logrus.Fatal("server Shutdown Failed:%+s", err)
	}

	logrus.Info("server exited properly")

	return
}

func endPoint(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
