package service

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/transaction"
)

var _ transaction.Service = Transaction{}

type Transaction struct {
	repo transaction.Repository
}

func NewTransactionService(repo transaction.Repository) *Transaction {
	return &Transaction{
		repo: repo,
	}
}
