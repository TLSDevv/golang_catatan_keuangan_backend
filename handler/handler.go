package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user"
	user_handler "github.com/TLSDevv/golang_catatan_keuangan_backend/handler/user"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	userService user.Service
}

func NewAPI(
	usService user.Service,
) *API {
	return &API{
		userService: usService,
	}
}

func (a API) Start(ctx context.Context, host, port string) (err error) {
	r := mux.NewRouter()

	fmt.Println(host, port)

	server := http.Server{
		Addr:    endPoint(host, port),
		Handler: r,
	}

	apiRoute := r.PathPrefix("/api/v1").Subrouter()

	user_handler.NewUserHandler(apiRoute, a.userService)

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
