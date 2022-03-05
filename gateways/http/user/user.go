package user

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/middleware"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	usService user.Service
	validator *util.StructValidator
}

func NewUserHandler(r *mux.Router, usService user.Service) UserHandler {
	userHandler := UserHandler{
		usService: usService,
		validator: util.NewValidate(),
	}
	r.HandleFunc("/register_user", userHandler.RegisterUser).Methods("POST")

	authRoute := r.PathPrefix("/").Subrouter()
	authRoute.Use(middleware.Authorization)

	authRoute.HandleFunc("/users/{id}", userHandler.FindById).Methods("GET")
	authRoute.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	authRoute.HandleFunc("/users/{id}/password", userHandler.UpdatePassword).Methods("PUT")
	authRoute.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")

	adminRoute := r.PathPrefix("/").Subrouter()
	adminRoute.Use(middleware.Admin)
	adminRoute.HandleFunc("/users", userHandler.List).Methods("GET")
	adminRoute.HandleFunc("/users", userHandler.Create).Methods("POST")

	return userHandler
}

func (h UserHandler) FindById(w http.ResponseWriter, r *http.Request) {

	userId := util.GetParams(r, "id")

	result, err := h.usService.FindById(r.Context(), userId)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendWithData(w, http.StatusOK, "Success Delete User", result)
	return
}

func (h UserHandler) List(w http.ResponseWriter, r *http.Request) {
	result, err := h.usService.List(r.Context())

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendWithData(w, http.StatusOK, "Success List User", result)
	return
}
