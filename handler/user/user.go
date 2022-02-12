package user

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/handler/middleware"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/handler/util"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	usService user.Service
}

func NewUserHandler(r *mux.Router, usService user.Service) UserHandler {
	userHandler := UserHandler{
		usService: usService,
	}
	r.HandleFunc("/users", userHandler.Create).Methods("POST")

	authRoute := r.PathPrefix("/").Subrouter()
	authRoute.Use(middleware.Authorization)
	authRoute.HandleFunc("/users", userHandler.List).Methods("GET")
	authRoute.HandleFunc("/users/{id}", userHandler.FindById).Methods("GET")
	authRoute.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
	authRoute.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")

	return userHandler
}

func (ush UserHandler) Create(w http.ResponseWriter, r *http.Request) {

	var reqBody entities.UserInput

	err := util.Decode(r, &reqBody)

	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	err = reqBody.Validate()

	if err != nil {
		util.SendNoData(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = ush.usService.Create(r.Context(), reqBody)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Create User")
	return
}

func (ush UserHandler) Update(w http.ResponseWriter, r *http.Request) {

	var reqBody entities.UserInput

	err := util.Decode(r, &reqBody)

	userId := util.GetParams(r, "id")

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = reqBody.Validate()

	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	err = ush.usService.Update(r.Context(), userId, reqBody)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Update User")
	return
}

func (ush UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

	userId := util.GetParams(r, "id")

	err := ush.usService.Delete(r.Context(), userId)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Delete User")
	return
}

func (ush UserHandler) FindById(w http.ResponseWriter, r *http.Request) {

	userId := util.GetParams(r, "id")

	result, err := ush.usService.FindById(r.Context(), userId)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendWithData(w, http.StatusOK, "Success Delete User", result)
	return
}

func (ush UserHandler) List(w http.ResponseWriter, r *http.Request) {
	result, err := ush.usService.List(r.Context())

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendWithData(w, http.StatusOK, "Success List User", result)
	return
}
