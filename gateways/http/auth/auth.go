package user

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/auth"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/middleware"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
	"github.com/gorilla/mux"
)

type AuthHandler struct {
	AuthService auth.Service
}

func NewAuthHandler(r *mux.Router, authService auth.Service) AuthHandler {
	authHandler := AuthHandler{
		AuthService: authService,
	}

	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	authRoute := r.PathPrefix("/").Subrouter()
	authRoute.Use(middleware.Authorization)
	authRoute.HandleFunc("/refresh", authHandler.Refresh).Methods("UPDATE")
	authRoute.HandleFunc("/logout", authHandler.Logout).Methods("DELETE")

	return authHandler
}

func (handler AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var reqBody entities.Login

	err := util.Decode(r, &reqBody)

	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	err = reqBody.Validate()
	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := handler.AuthService.Login(r.Context(), reqBody)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendWithData(w, http.StatusOK, "Success Login", token)
	return
}

func (handler AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	accessDetails, err := pkg.ExtractTokenMetadata(r, pkg.TypeTokenAccess)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := handler.AuthService.Refresh(r.Context(), accessDetails.Id)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendWithData(w, http.StatusOK, "Success Refresh Token", token)
	return
}

func (handler AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	err := handler.AuthService.Logout(r.Context())

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Logout")
	return
}
