package handler

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/service"
	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	trcService service.ITransactionService
}

func NewTransactionHandler(r *mux.Router, trcService service.ITransactionService) TransactionHandler {
	th := TransactionHandler{
		trcService: trcService,
	}

	r.HandleFunc("/transactions", th.FindAll).Methods("GET")

	return th
}

func (th TransactionHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	result, err := th.trcService.FindAll(r.Context())

	if err != nil {
		SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendWithData(w, http.StatusOK, "Success", result)
	return
}
