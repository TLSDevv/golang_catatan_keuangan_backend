package controller

import (
	"net/http"
)

type UserControllerInterface interface {
	CreateUser(writer http.ResponseWriter, request *http.Request)
	GetUser(writer http.ResponseWriter, request *http.Request)
	UpdateUser(writer http.ResponseWriter, request *http.Request)
}
