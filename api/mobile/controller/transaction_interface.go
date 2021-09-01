package controller

import (
	"net/http"
)

type TransactionControllerInterface interface {
	ListTransaction(writer http.ResponseWriter, request *http.Request)
	GetTransaction(writer http.ResponseWriter, request *http.Request)
	CreateTransaction(writer http.ResponseWriter, request *http.Request)
	UpdateTransaction(writer http.ResponseWriter, request *http.Request)
	DeleteTransaction(writer http.ResponseWriter, request *http.Request)
}
