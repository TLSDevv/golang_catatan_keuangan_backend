package auth

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/auth"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/middleware"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/gorilla/mux"
)

type AuthHandler struct {
	AuthService auth.Service
	validator   *util.StructValidator
}

func NewAuthHandler(r *mux.Router, authService auth.Service) AuthHandler {
	authHandler := AuthHandler{
		AuthService: authService,
		validator:   util.NewValidate(),
	}

	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/refresh", authHandler.Refresh).Methods("PUT")

	authRoute := r.PathPrefix("/").Subrouter()
	authRoute.Use(middleware.Authorization)
	authRoute.HandleFunc("/logout", authHandler.Logout).Methods("DELETE")

	return authHandler
}
