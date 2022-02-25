package transaction

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/transaction"
	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(r *mux.Router, s transaction.Service) *TransactionHandler {
	th := &TransactionHandler{
		service: s,
	}

	r.HandleFunc("/transactions", th.List).Methods("GET")
	r.HandleFunc("/transactions", th.Create).Methods("POST")
	r.HandleFunc("/transactions/{id}", th.GetByID).Methods("GET")
	r.HandleFunc("/transactions/{id}", th.Update).Methods("PUT")
	r.HandleFunc("/transactions/{id}", th.Delete).Methods("PUT")
	r.HandleFunc("/transactions/{id}", th.Restore).Methods("PUT")
	r.HandleFunc("/transactions/{id}", th.Purge).Methods("DELETE")

	return th
}
