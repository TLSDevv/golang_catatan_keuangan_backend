package controller

import "net/http"

type CategoryControllerInterface interface {
	CreateCategory(writer http.ResponseWriter, request *http.Request)
	GetCategory(writer http.ResponseWriter, request *http.Request)
	ListCategory(writer http.ResponseWriter, request *http.Request)
	UpdateCategory(writer http.ResponseWriter, request *http.Request)
	DeleteCategory(writer http.ResponseWriter, request *http.Request)
}
