package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile/service"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
	"github.com/go-chi/chi"
)

type UserController struct {
	UserService service.UserServiceInterface
}

func NewUserController(userService service.UserServiceInterface) UserControllerInterface {
	return &UserController{
		UserService: userService,
	}
}

func (u *UserController) CreateUser(writer http.ResponseWriter, request *http.Request) {
	userRequest := web.UserCreateRequest{}

	helper.ReadFromRequestBody(request, userRequest)

	u.UserService.CreateUser(request.Context(), userRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "SUCCESS CREATE USER",
	}

	helper.WriterToResponseBody(writer, webResponse)
}
func (u *UserController) GetUser(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		err := errors.New("Params Is Empty")
		helper.PanicIfError(err)
	}

	userId, _ := strconv.Atoi(id)

	user := u.UserService.GetUser(request.Context(), userId)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS GET DATA USER",
		Data:   user,
	}

	helper.WriterToResponseBody(writer, webResponse)
}
func (u *UserController) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		err := errors.New("Params Is Empty")
		helper.PanicIfError(err)
	}

	userId, _ := strconv.Atoi(id)

	userRequest := web.UserUpdateRequest{}

	helper.ReadFromRequestBody(request, userRequest)
	userRequest.Id = uint8(userId)

	u.UserService.UpdateUser(request.Context(), userRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS UPDATE USER DATA",
	}

	helper.WriterToResponseBody(writer, webResponse)
}
