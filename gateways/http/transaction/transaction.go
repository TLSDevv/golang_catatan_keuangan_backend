package transaction

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/transaction"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/middleware"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	service   transaction.Service
	validator *util.StructValidator
}

func NewTransactionHandler(r *mux.Router, s transaction.Service) *TransactionHandler {
	th := &TransactionHandler{
		service:   s,
		validator: util.NewValidate(),
	}

	pr := r.PathPrefix("/").Subrouter()
	pr.Use(middleware.Authorization)
	pr.HandleFunc("/transactions", th.List).Methods("GET")
	pr.HandleFunc("/transactions", th.Create).Methods("POST")
	pr.HandleFunc("/transactions/{id}", th.GetByID).Methods("GET")
	pr.HandleFunc("/transactions/{id}/update", th.Update).Methods("PUT")
	pr.HandleFunc("/transactions/{id}/delete", th.Delete).Methods("PUT")
	pr.HandleFunc("/transactions/{id}/restore", th.Restore).Methods("PUT")
	pr.HandleFunc("/transactions/{id}", th.Purge).Methods("DELETE")

	return th
}
