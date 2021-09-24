package controller

import (
	"encoding/json"
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile/parameter"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile/service"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
)

type Auth struct {
	authService *service.AuthService
}

func NewAuth(authService *service.AuthService) *Auth {
	return &Auth{
		authService: authService,
	}
}

func (auth *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var req parameter.LoginRequest

	json.NewDecoder(r.Body).Decode(&req)

	// TODO validation

	res, err := auth.authService.Login(r.Context(), req)

	helper.PanicIfError(err)

	helper.WriterToResponseBody(w, res)
}

func (auth *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	var req parameter.LogoutRequest

	json.NewDecoder(r.Body).Decode(&req)

	// TODO validation

	res, err := auth.authService.Logout(r.Context(), req)

	helper.PanicIfError(err)

	helper.WriterToResponseBody(w, res)
}
