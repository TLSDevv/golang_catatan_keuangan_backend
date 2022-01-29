package handler

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/dto"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/service"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	usService service.IUserService
}

// r.URL.Query()["name"]

func NewUserHandler(r *mux.Router, usService service.IUserService) UserHandler {
	userHandler := UserHandler{
		usService: usService,
	}

	r.HandleFunc("/users", userHandler.List).Methods("GET")
	r.HandleFunc("/users", userHandler.Create).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.FindById).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")

	return userHandler
}

func (ush UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var reqBody dto.UserRequest

	err := Encode(w, &reqBody)

	if err != nil {
		SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = reqBody.Validate()

	if err != nil {
		SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	err = ush.usService.Create(r.Context(), reqBody)

	if err != nil {
		SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendNoData(w, http.StatusOK, "Success Create User")
	return
}

func (ush UserHandler) Update(w http.ResponseWriter, r *http.Request) {

	var reqBody dto.UserRequest

	err := Encode(w, &reqBody)

	userId := GetParams(r, "id")

	if err != nil {
		SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = reqBody.Validate()

	if err != nil {
		SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	err = ush.usService.Update(r.Context(), userId, reqBody)

	if err != nil {
		SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendNoData(w, http.StatusOK, "Success Update User")
	return
}

func (ush UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

	userId := GetParams(r, "id")

	err := ush.usService.Delete(r.Context(), userId)

	if err != nil {
		SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendNoData(w, http.StatusOK, "Success Delete User")
	return
}

func (ush UserHandler) FindById(w http.ResponseWriter, r *http.Request) {

	userId := GetParams(r, "id")

	result, err := ush.usService.FindById(r.Context(), userId)

	if err != nil {
		SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendWithData(w, http.StatusOK, "Success Delete User", result)
	return
}

func (ush UserHandler) List(w http.ResponseWriter, r *http.Request) {
	result, err := ush.usService.List(r.Context())

	if err != nil {
		SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendWithData(w, http.StatusOK, "Success List User", result)
	return
}
