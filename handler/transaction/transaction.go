package transaction

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/transaction"
	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(r *mux.Router, s transaction.Service) *TransactionHandler {
	h := &TransactionHandler{
		service: s,
	}

	r.HandleFunc("/transactions", h.List).Methods("GET")
	r.HandleFunc("/transactions", h.Create).Methods("POST")

	return h
}
