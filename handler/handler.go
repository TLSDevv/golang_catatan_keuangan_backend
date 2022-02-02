package handler

import (
	"fmt"
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/handler/middleware"
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

func (a API) Start(host, port string) {
	r := mux.NewRouter()

	r.Use(middleware.CORS)

	apiRoute := r.PathPrefix("/api/v1").Subrouter()

	user_handler.NewUserHandler(apiRoute, a.userService)

	server := http.Server{
		Addr:    endPoint(host, port),
		Handler: r,
	}

	fmt.Printf("Listening %s to port %s", host, port)
	err := server.ListenAndServe()

	if err != nil {
		logrus.Error("error Listen serve ", err)
		return
	}
}

func endPoint(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
