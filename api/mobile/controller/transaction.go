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

type TransactionController struct {
	TransactionService service.TransactionServiceInterface
}

func NewTransactionController(transactionSerivice service.TransactionServiceInterface) TransactionControllerInterface {
	return &TransactionController{
		TransactionService: transactionSerivice,
	}
}

func (t *TransactionController) ListTransaction(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		err := errors.New("Params Is Empty")
		helper.PanicIfError(err)
	}

	userId, _ := strconv.Atoi(id)
	limit, _ := strconv.Atoi(request.URL.Query().Get("limit"))
	page, _ := strconv.Atoi(request.URL.Query().Get("page"))

	transactions := t.TransactionService.ListTransaction(request.Context(), limit, page, userId)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS GET DATA TRANSACTION",
		Data:   transactions,
	}

	helper.WriterToResponseBody(writer, webResponse)
}
func (t *TransactionController) GetTransaction(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		err := errors.New("Params Is Empty")
		helper.PanicIfError(err)
	}

	userId, _ := strconv.Atoi(id)

	transaction := t.TransactionService.GetTransaction(request.Context(), userId)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS GET DATA TRANSACTION",
		Data:   transaction,
	}

	helper.WriterToResponseBody(writer, webResponse)
}
func (t *TransactionController) CreateTransaction(writer http.ResponseWriter, request *http.Request) {
	transactionRequest := web.TransactionCreateRequest{}

	helper.ReadFromRequestBody(request, &transactionRequest)

	t.TransactionService.CreateTransaction(request.Context(), transactionRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "SUCCESS CREATE TRANSACTION",
	}

	helper.WriterToResponseBody(writer, webResponse)
}
func (t *TransactionController) UpdateTransaction(writer http.ResponseWriter, request *http.Request) {
	transactionRequest := web.TransactionUpdateRequest{}

	helper.ReadFromRequestBody(request, &transactionRequest)

	id := chi.URLParam(request, "id")

	if id == "" {
		err := errors.New("Params Is Empty")
		helper.PanicIfError(err)
	}

	transactionId, _ := strconv.Atoi(id)
	transactionRequest.Id = uint8(transactionId)

	t.TransactionService.UpdateTransaction(request.Context(), transactionRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS UPDATE TRANSACTION",
	}

	helper.WriterToResponseBody(writer, webResponse)
}
func (t *TransactionController) DeleteTransaction(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		err := errors.New("Params Is Empty")
		helper.PanicIfError(err)
	}

	transactionId, _ := strconv.Atoi(id)

	t.TransactionService.DeleteTransaction(request.Context(), transactionId)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS DELETE TRANSACTION",
	}

	helper.WriterToResponseBody(writer, webResponse)
}
