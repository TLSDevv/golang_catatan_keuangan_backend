package service

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/transaction"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user"
)

var _ transaction.Service = Transaction{}

type Transaction struct {
	tr transaction.Repository
	ur user.Repository
}

func NewTransactionService(tr transaction.Repository, ur user.Repository) *Transaction {
	return &Transaction{
		tr: tr,
		ur: ur,
	}
}
